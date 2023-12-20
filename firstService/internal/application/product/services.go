package product

import "toscale-test-task/firstService/internal/domain/services"

type Services struct {
	DataService services.DataService
}

func NewServices(repos *Storages) *Services {
	return &Services{
		DataService: services.NewDataService(repos.Klines),
	}
}
