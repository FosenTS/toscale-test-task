package application

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	config "toscale-test-task/first-service/internal/application/config"
)

type App interface {
	Run(ctx context.Context) error
}

type app struct {
	db *gorm.DB
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

	gorm, err := createGormConnection(
		gormConfig,
		log.WithField("location", "gorm"))
	if err != nil {
		return nil, err
	}

	log.Println(gorm)

	return nil, nil
}

func createGormConnection(cfg config.GormConfig, log *logrus.Entry) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", cfg.Host, cfg.User, cfg.Password, cfg.DatabaseName, cfg.Port, cfg.SSLMode, cfg.TimeZone.String())
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(fmt.Errorf("error connect gorm to database: %w", err))
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnection)
	sqlDB.SetConnMaxLifetime(cfg.ConnectionMaxLifeTime)

	return db, nil
}
