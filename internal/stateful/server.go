package stateful

import (
	proto "ably/protos/stateful"
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"time"
)

type SessionStore interface {
	//TODO
}

type StatefulServer struct {
	log    logrus.FieldLogger
	store  SessionStore
	server *grpc.Server
}

func New(log *logrus.Logger, store SessionStore) *StatefulServer {
	s := &StatefulServer{
		log:  log,
		store: store,
	}

	opts := []grpc_logrus.Option{
		grpc_logrus.WithDecider(func(fullMethodName string, err error) bool {
			switch fullMethodName {
			default:
				return true
			}
		})}

	s.server = grpc.NewServer(
		grpc_middleware.WithStreamServerChain(
			grpc_logrus.StreamServerInterceptor(logrus.NewEntry(log), opts...),
			grpc_recovery.StreamServerInterceptor(),
		),
	)

	proto.RegisterStatefulNumberGeneratorServer(s.server, s)
	return s
}

func (s *StatefulServer) Run(ctx context.Context, externalAddress string) error {
	return runServer(ctx, s.server, externalAddress)
}

func runServer(ctx context.Context, srv *grpc.Server, address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return errors.Wrapf(err, "failed to start listener on %s", address)
	}

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			tryGracefulStop(srv)
		}
	}(ctx)
	return srv.Serve(lis)
}

func tryGracefulStop(srv *grpc.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c := make(chan struct{})
	go func() {
		srv.GracefulStop()
		c <- struct{}{}
	}()

	select {
	case <-c:
	case <-ctx.Done():
		srv.Stop()
	}
	return
}
