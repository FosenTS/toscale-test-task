package product

import (
	"toscale-test-task/first-service/internal/infrastructure/controllers/httpController"
	"toscale-test-task/first-service/internal/infrastructure/controllers/httpController/clientHandler"
)

type Controllers struct {
	HttpController httpController.HTTPHandler
}

func NewControllers(gateways *Gateways) *Controllers {
	return &Controllers{
		HttpController: httpController.NewHttpController(
			clientHandler.NewClientHandler(gateways.Services.DataService),
		),
	}
}
