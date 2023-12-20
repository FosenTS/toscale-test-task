package product

import "google.golang.org/grpc"

type Endpoints struct {
	*Controllers
}

func NewEndpoint(controllers *Controllers) *Endpoints {
	return &Endpoints{Controllers: controllers}
}

func (e *Endpoints) RegisterServer(grpc *grpc.Server) {
	e.GRPCController.RegisterServer(grpc)
}
