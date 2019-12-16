package stateless

import (
	proto "ably/protos/stateless"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		err := stream.Send(&proto.Generated{Number: toSend})
		if err != nil {
			return status.Error(codes.Internal, errors.Wrap(err, "sending on stream").Error())
		}
		//TODO handle integer overflow
		toSend = toSend << 1

		select {
		case <-stream.Context().Done():
			return nil
		case <-ticker.C:
			continue
		}
	}
}
