package product

import (
	"github.com/sirupsen/logrus"
	"toscale-test-task/firstService/internal/infrastructure/controllers/httpController"
	"toscale-test-task/firstService/internal/infrastructure/controllers/httpController/clientHandler"
)

type Controllers struct {
	HttpController httpController.HTTPHandler
}

func NewControllers(gateways *Gateways, entry *logrus.Entry) *Controllers {
	return &Controllers{
		HttpController: httpController.NewHttpController(
			clientHandler.NewClientHandler(gateways.Services.DataService, entry.WithField("location", "client-handlers"), gateways.GRPCGateway),
		),
	}
}
