package product

import (
	"github.com/sirupsen/logrus"
	"toscale-test-task/secondService/internal/infrastructure/controller/gRPC/serverController"
)

type Controllers struct {
	GRPCController serverController.GRPCController
}

func NewControllers(gateways *Gateways, entry *logrus.Entry) *Controllers {
	return &Controllers{GRPCController: serverController.NewServer(gateways.Binance, entry.WithField("location", "server-gRPS-controller"))}
}
