package httpController

import (
	"github.com/fasthttp/router"
)

type HTTPHandler interface {
	HandlerRouter(*router.Group)
}

type httpController struct {
	clientHandler HTTPHandler
}

func NewHttpController(clientHandler HTTPHandler) HTTPHandler {
	return &httpController{clientHandler: clientHandler}
}

func (hC *httpController) HandlerRouter(r *router.Group) {
	api := r.Group("/kline")

	hC.clientHandler.HandlerRouter(api)
}
