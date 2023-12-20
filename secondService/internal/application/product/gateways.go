package product

import (
	"github.com/sirupsen/logrus"
	"toscale-test-task/secondService/internal/application/config"
	"toscale-test-task/secondService/internal/infrastructure/gateway/binanceGateway"
)

type Gateways struct {
	Binance binanceGateway.BinanceGateway
}

func NewGateways(configBinance *config.BinanceConfig, entry *logrus.Entry) *Gateways {
	return &Gateways{Binance: binanceGateway.NewBinanceGateway(configBinance, entry.WithField("location", "binance-gateway"))}
}
