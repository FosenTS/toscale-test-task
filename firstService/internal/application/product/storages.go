package product

import (
	"gorm.io/gorm"
	"toscale-test-task/firstService/internal/domain/storage"
	"toscale-test-task/firstService/internal/infrastructure/repository"
)

type Storages struct {
	Klines storage.KlinesStorage
}

func NewStorages(gorm *gorm.DB) (*Storages, error) {
	klinesRepository := repository.NewKlinesRepository(gorm)
	err := klinesRepository.Migrate()
	if err != nil {
		return nil, err
	}
	return &Storages{
		Klines: klinesRepository,
	}, nil
}
