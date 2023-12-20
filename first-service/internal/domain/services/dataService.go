package services

import (
	"toscale-test-task/first-service/internal/domain/storage"
)

type DataService interface {
}

type dataService struct {
	klinesStorage storage.KlinesStorage
}

func NewDataService(klinesStorage storage.KlinesStorage) DataService {
	return &dataService{klinesStorage: klinesStorage}
}
