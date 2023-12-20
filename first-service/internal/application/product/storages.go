package product

import (
	"gorm.io/gorm"
	"toscale-test-task/first-service/internal/domain/storage"
	"toscale-test-task/first-service/internal/infrastructure/repository"
)

type Storages struct {
	Klines storage.KlinesStorage
}

func NewStorages(gorm *gorm.DB) *Storages {
	return &Storages{
		Klines: repository.NewKlinesRepository(gorm),
	}
}
