package serverController

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"toscale-test-task/secondService/internal/infrastructure/gateway/binanceGateway"
	"toscale-test-task/secondService/protoMessages"
)

type GRPCController interface {
	RegisterServer(grpc *grpc.Server)
}

type Server struct {
	protoMessages.UnimplementedKlineServiceServer

	binanceGateway binanceGateway.BinanceGateway

	log *logrus.Entry
}

func NewServer(binanceGateway binanceGateway.BinanceGateway, log *logrus.Entry) GRPCController {
	return &Server{
		binanceGateway: binanceGateway,
		log:            log,
	}
}

func (s *Server) RegisterServer(grpc *grpc.Server) {
	protoMessages.RegisterKlineServiceServer(grpc, s)
}

func (s *Server) Kline(ctx context.Context, in *protoMessages.KlineRequest) (*protoMessages.KlineResponce, error) {
	symbol := in.GetSymbol()
	interval := in.GetInterval()

	klines, err := s.binanceGateway.GetKlines(ctx, symbol, interval)
	if err != nil {
		return nil, err
	}

	resp := &protoMessages.KlineResponce{
		Items: make([]*protoMessages.KlineItem, 0),
	}

	for _, kline := range klines {
		resp.Items = append(
			resp.Items,
			&protoMessages.KlineItem{
				KlineOpenTime:            kline.KlineOpenTime,
				OpenPrice:                kline.OpenPrice,
				HighPrice:                kline.HighPrice,
				LowPrice:                 kline.LowPrice,
				ClosePrice:               kline.ClosePrice,
				Volume:                   kline.Volume,
				KlineCloseTime:           kline.KlineCloseTime,
				QuoteAssetVolume:         kline.QuoteAssetVolume,
				NumberOfTrades:           kline.NumberOfTrades,
				TakerBuyBaseAssetVolume:  kline.TakerBuyBaseAssetVolume,
				TakerBuyQuoteAssetVolume: kline.TakerBuyQuoteAssetVolume,
			},
		)
	}

	return resp, nil
}
