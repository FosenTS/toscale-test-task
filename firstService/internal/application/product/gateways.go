package product

import (
	"github.com/sirupsen/logrus"
	"toscale-test-task/firstService/internal/infrastructure/gateways/gRPCGateway"
	"toscale-test-task/firstService/protoMessages"
)

type Gateways struct {
	*Services
	gRPCGateway.GRPCGateway
}

func NewGateways(services *Services, grpcClient protoMessages.KlineServiceClient, entry *logrus.Entry) *Gateways {
	return &Gateways{
		Services:    services,
		GRPCGateway: gRPCGateway.NewGRPCGateway(grpcClient, entry.WithField("location", "grpc-gateway")),
	}
}
