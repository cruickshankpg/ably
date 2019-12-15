package stateful

import (
	proto "ably/protos/stateful"
)

func (s *StatefulServer) GenerateSequence(*proto.GenerateSequenceRequest, proto.StatefulNumberGenerator_GenerateSequenceServer) error {
	panic("implement me")
}

func (s *StatefulServer) ReconnectSequence(*proto.ReconnectSequenceRequest, proto.StatefulNumberGenerator_ReconnectSequenceServer) error {
	panic("implement me")
}