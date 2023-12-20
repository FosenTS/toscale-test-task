package repository

import (
	"context"
	"gorm.io/gorm"
	"toscale-test-task/firstService/internal/domain/entity"
	"toscale-test-task/firstService/internal/domain/storage"
	"toscale-test-task/firstService/internal/domain/storage/dto"
)

type KlinesRepository storage.KlinesStorage

type klinesRepository struct {
	gorm *gorm.DB
}

func NewKlinesRepository(gorm *gorm.DB) KlinesRepository {
	return &klinesRepository{gorm: gorm}
}

func (kR *klinesRepository) Migrate() error {
	err := kR.gorm.AutoMigrate(&entity.Kline{})
	if err != nil {
		return err
	}

	return nil
}

func (kR *klinesRepository) StoreMany(ctx context.Context, klines []*dto.KlineCreate) error {
	kR.gorm.Table("klines").Create(klines)
	return nil
}
