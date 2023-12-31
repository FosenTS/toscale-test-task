package application

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"toscale-test-task/secondService/internal/application/config"
	"toscale-test-task/secondService/internal/application/product"
)

type App interface {
	Run(ctx context.Context, log *logrus.Entry) error
	runGRPCServer(log *logrus.Entry, config *config.GRPCConfig) error
}

type app struct {
	rpcConfig *config.GRPCConfig
	endpoints *product.Endpoints
}

func NewApp(ctx context.Context, log *logrus.Entry) (App, error) {
	log.Println("Start creating application")
	err := config.Env()
	if err != nil {
		log.Fatalln(fmt.Errorf("fatal read environments: %w", err))
		return nil, err
	}

	binanceConfig, err := config.Binance()
	if err != nil {
		log.Fatalln(fmt.Errorf("falat create binance config: %w", err))
		return nil, err
	}

	gRPCConfig, err := config.GRPC()
	if err != nil {
		log.Fatalln(fmt.Errorf("falat create gRPC config: %w", err))
		return nil, err
	}

	gateways := product.NewGateways(binanceConfig, log)
	controllers := product.NewControllers(gateways, log)

	endpoints := product.NewEndpoint(controllers)

	return &app{
		rpcConfig: gRPCConfig,
		endpoints: endpoints,
	}, nil
}

func (app *app) Run(ctx context.Context, log *logrus.Entry) error {
	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return app.runGRPCServer(log.WithField("location", "gRPC-Server"), app.rpcConfig)
	})

	return grp.Wait()
}

func (app *app) runGRPCServer(log *logrus.Entry, config *config.GRPCConfig) error {

	lis, err := net.Listen("tcp", config.Address)
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to listen: %w", err))
		return err
	}

	grpcServer := grpc.NewServer(grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_ctxtags.StreamServerInterceptor(),
		grpc_opentracing.StreamServerInterceptor(),
		grpc_logrus.StreamServerInterceptor(log),
		grpc_recovery.StreamServerInterceptor(),
	)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(log),
			grpc_recovery.UnaryServerInterceptor(),
		)))

	app.endpoints.RegisterServer(grpcServer)

	log.Infoln(fmt.Sprintf("gRPC server starting: %s", config.Address))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(fmt.Errorf("failed to serve: %w", err))
		return err
	}

	return nil
}
