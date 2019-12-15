package stateless

import (
	proto "ably/protos/stateless"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

func (s *StatelessServer) GenerateSequence(req *proto.GenerateSequenceRequest, stream proto.StatelessNumberGenerator_GenerateSequenceServer) error {
	// If this is a reconnect the client will provide a number to start from. Otherwise generate a new random one
	toSend := req.RestartFrom
	if toSend == 0 {
		toSend = uint32(rand.Int31n(0xff))
	} else {
		toSend = toSend << 1
	}

	err := stream.Send(&proto.Generated{Number: toSend})
	if err != nil {
		return errors.Wrap(err, "sending on stream")
	}
	// TODO handle the integer overflowing
	toSend = toSend << 1

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	select {
	case <-stream.Context().Done():
		return nil
	case <-ticker.C:
		err := stream.Send(&proto.Generated{Number: toSend})
		if err != nil {
			return errors.Wrap(err, "sending on stream")
		}
		toSend = toSend << 1
	}

	return nil
}
