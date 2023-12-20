package repository

import (
	"gorm.io/gorm"
	"toscale-test-task/first-service/internal/domain/storage"
)

type KlinesRepository storage.KlinesStorage

type klinesRepository struct {
	gorm *gorm.DB
}

func NewKlinesRepository(gorm *gorm.DB) KlinesRepository {
	return &klinesRepository{gorm: gorm}
}
