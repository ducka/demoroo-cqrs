package grpc

import (
	"context"
	"net"

	"github.com/ducka/demoroo/internal/grpc/api"
	"github.com/ducka/demoroo/proto"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ServerSettings struct {
	HostAddress string
}

type ServerRunner struct {
	listener net.Listener
	server   *grpc.Server
}

func NewServerRunner(lifecycle fx.Lifecycle, logger *zap.Logger, branchSearchServiceServer proto.BranchSearchServiceServer, grpcServerSettings ServerSettings) (*ServerRunner, error) {
	lis, err := net.Listen("tcp", grpcServerSettings.HostAddress)

	if err != nil {
		logger.Fatal("failed to listen: %v", zap.Error(err))
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	proto.RegisterBranchSearchServiceServer(server, branchSearchServiceServer)

	// Register start/stop of the GRPC server
	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			server.Stop()
			return nil
		},
	})

	return &ServerRunner{
		listener: lis,
		server:   server,
	}, err
}

func (t *ServerRunner) Run() error {
	go func() {
		t.server.Serve(t.listener)
	}()

	return nil
}

func RegisterServer() fx.Option {
	return fx.Module(
		"grpc-server",
		fx.Provide(
			NewServerRunner,
			fx.Annotate(
				api.NewBranchSearchService,
				fx.As(new(proto.BranchSearchServiceServer)),
			),
		),
		fx.Invoke(func(serverRunner *ServerRunner) error {
			return serverRunner.Run()
		}),
	)
}
