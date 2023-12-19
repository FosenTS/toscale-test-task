package application

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
	config "toscale-test-task/first-service/internal/application/config"

	"golang.org/x/sync/errgroup"
)

type App interface {
	Run(ctx context.Context, log *logrus.Entry) error
	runHTTP(log *logrus.Entry) error
}

type app struct {
	db         *gorm.DB
	httpConfig *config.HTTPConfig
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

	gormConn, err := createGormConnection(
		gormConfig,
		log.WithField("location", "gorm"))
	if err != nil {
		return nil, err
	}

	httpConfig, err := config.Http()
	if err != nil {
		log.Errorln("error creating http config")
		return nil, err
	}

	return &app{
		db:         gormConn,
		httpConfig: httpConfig,
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
	err := fasthttp.ListenAndServe(app.httpConfig.HttpServer, fasthttp.TimeoutHandler(func(ctx *fasthttp.RequestCtx) {
		select {
		case <-ctx.Done():
			// ctx.Done() is only closed when the server is shutting down.
			log.Println("context cancelled")
			return
		case <-time.After(10 * time.Second):
			log.Println("process finished ok")
		}
	}, time.Second*2, "timeout"))
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error start listen and serve http server: %w", err))
		return err
	}
	return nil
}

func createGormConnection(cfg config.GormConfig, log *logrus.Entry) (*gorm.DB, error) {
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
