package storage

import (
	"context"
	"toscale-test-task/firstService/internal/domain/storage/dto"
)

type KlinesStorage interface {
	Migrate() error
	StoreMany(ctx context.Context, klines []*dto.KlineCreate) error
}
