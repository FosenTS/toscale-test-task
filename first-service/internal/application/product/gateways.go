package product

type Gateways struct {
	*Services
}

func NewGateways(services *Services) *Gateways {
	return &Gateways{Services: services}
}
