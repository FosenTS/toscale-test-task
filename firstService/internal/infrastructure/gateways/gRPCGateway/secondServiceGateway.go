package gRPCGateway

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"toscale-test-task/firstService/internal/domain/storage/dto"
	"toscale-test-task/firstService/protoMessages"
)

type GRPCGateway interface {
	KlineRequest(ctx context.Context, symbol, interval string) ([]*dto.KlineCreate, error)
}

type gRPCGateway struct {
	client protoMessages.KlineServiceClient
	log    *logrus.Entry
}

func NewGRPCGateway(client protoMessages.KlineServiceClient, log *logrus.Entry) GRPCGateway {
	return &gRPCGateway{client: client, log: log}
}

func (g *gRPCGateway) KlineRequest(ctx context.Context, symbol, interval string) ([]*dto.KlineCreate, error) {
	r, err := g.client.Kline(ctx, &protoMessages.KlineRequest{Symbol: symbol, Interval: interval})
	if err != nil {
		g.log.Errorln(fmt.Errorf("fatal sending gRPC request: %w", err))
		return nil, err
	}
	klines := make([]*dto.KlineCreate, 0)
	for _, item := range r.GetItems() {
		klines = append(
			klines,
			dto.NewKlineCreate(
				symbol,
				interval,
				item.KlineCloseTime,
				item.OpenPrice,
				item.HighPrice,
				item.LowPrice,
				item.ClosePrice,
				item.Volume,
				item.KlineCloseTime,
				item.QuoteAssetVolume,
				item.NumberOfTrades,
				item.TakerBuyBaseAssetVolume,
				item.TakerBuyQuoteAssetVolume,
			),
		)
	}

	return klines, nil
}
