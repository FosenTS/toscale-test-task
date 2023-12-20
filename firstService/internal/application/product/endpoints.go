package product

type Endpoints struct {
	*Controllers
}

func NewEndpoints(controllers *Controllers) *Endpoints {
	return &Endpoints{Controllers: controllers}
}
