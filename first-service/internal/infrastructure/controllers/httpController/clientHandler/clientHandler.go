package clientHandler

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"toscale-test-task/first-service/internal/domain/services"
	"toscale-test-task/first-service/internal/infrastructure/controllers/httpController"
)

type ClientHandler struct {
	dataService services.DataService
}

func (cH *ClientHandler) HandlerRouter(g *router.Group) {
	g.GET("/conn", cH.clientRequest)
}

func NewClientHandler(dataService services.DataService) httpController.HTTPHandler {
	return &ClientHandler{dataService: dataService}
}

func (cH *ClientHandler) clientRequest(ctx *fasthttp.RequestCtx) {
	fmt.Println("Client Connected")
	args := ctx.QueryArgs()
	arg := args.Peek("arg")
	fmt.Println(arg)
}
