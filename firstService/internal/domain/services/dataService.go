package services

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"toscale-test-task/firstService/internal/domain/entity"
	"toscale-test-task/firstService/internal/domain/storage"
	"toscale-test-task/firstService/internal/domain/storage/dto"
)

type DataService interface {
	StoreKline(ctx context.Context, klines []*dto.KlineCreate) ([]*entity.Kline, error)
}

type dataService struct {
	klinesStorage storage.KlinesStorage

	log *logrus.Entry
}

func NewDataService(klinesStorage storage.KlinesStorage) DataService {
	return &dataService{klinesStorage: klinesStorage}
}

func (dS *dataService) StoreKline(ctx context.Context, klinesCreate []*dto.KlineCreate) ([]*entity.Kline, error) {
	klines := make([]*entity.Kline, 0)

	err := dS.klinesStorage.StoreMany(ctx, klinesCreate)
	if err != nil {
		err := fmt.Errorf("error storing klines: %w", err)
		dS.log.Errorln(err)
		return nil, err
	}
	for _, kline := range klinesCreate {
		klines = append(
			klines,
			kline.ToEntity(),
		)
	}

	return klines, nil
}
