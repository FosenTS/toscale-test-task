package application

import (
	"context"
	"fmt"
	"github.com/AubSs/fasthttplogger"
	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"toscale-test-task/firstService/internal/application/config"
	"toscale-test-task/firstService/internal/application/product"
	"toscale-test-task/firstService/protoMessages"
)

type App interface {
	Run(ctx context.Context, log *logrus.Entry) error
	runHTTP(log *logrus.Entry) error
}

type app struct {
	endpoint   *product.Endpoints
	httpConfig *config.HTTPConfig
	log        *logrus.Entry
}

func NewApp(ctx context.Context, log *logrus.Entry) (App, error) {
	log.Println("Start creating application")
	err := config.Env()
	if err != nil {
		log.Fatalln(fmt.Errorf("fatal reading env config: %w", err))
		return nil, err
	}

	gormConfig, err := config.Gorm()
	if err != nil {
		log.Errorln(fmt.Errorf("cannot read gorm config: %w", err))
		return nil, err
	}

	gormDB, err := createGormConnection(
		gormConfig,
		log.WithField("location", "gorm"))
	if err != nil {
		return nil, err
	}

	httpConfig, err := config.Http()
	if err != nil {
		log.Errorln(fmt.Errorf("error creating http config: %w", err))
		return nil, err
	}

	grpcConfig, err := config.GRPC()
	if err != nil {
		log.Errorln(fmt.Errorf("error creating gRPC config: %w", err))
		return nil, err
	}
	clientGRPC, err := createGRPCConnection(grpcConfig, log)
	if err != nil {
		log.Errorln(fmt.Errorf("error create gRPC client connection: %w", err))
		return nil, err
	}

	storages, err := product.NewStorages(gormDB)
	if err != nil {
		err := fmt.Errorf("error creating storages: %w", err)
		log.Errorln(err)
		return nil, err
	}

	services := product.NewServices(storages)

	gateways := product.NewGateways(services, clientGRPC, log)

	controllers := product.NewControllers(gateways, log)

	endpoints := product.NewEndpoints(controllers)
	return &app{
		endpoint:   endpoints,
		httpConfig: httpConfig,
		log:        log.WithField("location", "app"),
	}, nil
}

func (app *app) Run(ctx context.Context, log *logrus.Entry) error {
	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return app.runHTTP(log.WithField("location", "runHTTP"))
	})

	return grp.Wait()
}

func (app *app) runHTTP(log *logrus.Entry) error {
	app.log.Infoln("Configure api routes")

	r := router.New()
	app.endpoint.Controllers.HttpController.HandlerRouter(r.Group("/api"))
	list := r.List()
	for key, routesGroup := range list {
		for _, route := range routesGroup {
			log.Infoln(fmt.Sprintf("%s : %s", key, route))
		}
	}

	err := fasthttp.ListenAndServe(app.httpConfig.HttpServer, fasthttplogger.CombinedColored(r.Handler))
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error start listen and serve http server: %w", err))
		return err
	}
	return nil
}

func createGormConnection(cfg *config.GormConfig, log *logrus.Entry) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", cfg.Host, cfg.User, cfg.Password, cfg.DatabaseName, cfg.Port, cfg.SSLMode, cfg.TimeZone.String())
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(fmt.Errorf("error connect gorm to database: %w", err))
		return nil, err
	}

	log.Infoln("Successful connection")

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnection)
	sqlDB.SetConnMaxLifetime(cfg.ConnectionMaxLifeTime)

	log.Infoln("Gorm has been successfully configured")

	return db, nil
}

func createGRPCConnection(cfg *config.GRPCConfig, log *logrus.Entry) (protoMessages.KlineServiceClient, error) {
	conn, err := grpc.Dial(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(fmt.Errorf("did not connect: %w", err))
		return nil, err
	}

	c := protoMessages.NewKlineServiceClient(conn)

	return c, nil
}
