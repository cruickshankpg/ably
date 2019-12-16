package stateful

import (
	"ably/internal/store"
	proto "ably/protos/stateful"
	"crypto/sha256"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"strconv"
	"time"
)

func (s *StatefulServer) GenerateSequence(req *proto.GenerateSequenceRequest, stream proto.StatefulNumberGenerator_GenerateSequenceServer) error {
	if len(req.ConnectionID) < 1 {
		return status.Error(codes.InvalidArgument, "missing connection id")
	}

	if req.SequenceLength < 1 {
		return status.Error(codes.InvalidArgument, "invalid sequence length")
	}

	seed := time.Now().Unix()
	s.store.Set(req.ConnectionID, store.SessionState{Seed: seed, SequenceLength: req.SequenceLength})

	prng := rand.NewSource(seed).(rand.Source64)

	seq := make([]byte, 0, 1000)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for i := uint32(0); i < req.SequenceLength; i++ {
		num := uint32(prng.Uint64() >> 32)
		seq = strconv.AppendUint(seq, uint64(num), 10)

		gen := &proto.Generated{Number: num}
		if i == req.SequenceLength-1 {
			gen.FinalItem = true
			chSum := sha256.Sum256(seq)
			gen.Checksum = chSum[:]
		}

		err := stream.Send(gen)
		if err != nil {
			return status.Error(codes.Internal, errors.Wrap(err, "sending on stream").Error())
		}

		select {
		case <-stream.Context().Done():
			s.store.Expire(req.ConnectionID, time.Second*30)
			return nil
		case <-ticker.C:
			continue
		}
	}

	// The client is responsible for closing the connection once the sequence has been sent
	select {
	case <-stream.Context().Done():
		s.store.Delete(req.ConnectionID)
		return nil
	}
}

func (s *StatefulServer) ReconnectSequence(req *proto.ReconnectSequenceRequest, stream proto.StatefulNumberGenerator_ReconnectSequenceServer) error {
	if len(req.ConnectionID) < 1 {
		return status.Error(codes.InvalidArgument, "missing connection id")
	}

	state, ok := s.store.Get(req.ConnectionID)
	if !ok {
		return status.Error(codes.NotFound, "no connection found")
	}

	if req.LastReceivedIndex >= state.SequenceLength-1 {
		return status.Error(codes.InvalidArgument, "last received index is larger than sequence length")
	}

	// clear expiry
	s.store.Expire(req.ConnectionID, 0)

	prng := rand.NewSource(state.Seed).(rand.Source64)
	seq := make([]byte, 0, 1000)

	// rebuild previously sent sequence
	var i uint32
	for ; i <= req.LastReceivedIndex; i++ {
		num := uint32(prng.Uint64() >> 32)
		seq = strconv.AppendUint(seq, uint64(num), 10)
	}

	// now send the rest of the sequence
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for ; i < state.SequenceLength; i++ {
		num := uint32(prng.Uint64() >> 32)
		seq = strconv.AppendUint(seq, uint64(num), 10)

		gen := &proto.Generated{Number: num}
		if i == state.SequenceLength-1 {
			gen.FinalItem = true
			chSum := sha256.Sum256(seq)
			gen.Checksum = chSum[:]
		}

		err := stream.Send(gen)
		if err != nil {
			return status.Error(codes.Internal, errors.Wrap(err, "sending on stream").Error())
		}

		select {
		case <-stream.Context().Done():
			s.store.Expire(req.ConnectionID, time.Second*30)
			return nil
		case <-ticker.C:
			continue
		}
	}

	select {
	case <-stream.Context().Done():
		s.store.Delete(req.ConnectionID)
		return nil
	}
}
