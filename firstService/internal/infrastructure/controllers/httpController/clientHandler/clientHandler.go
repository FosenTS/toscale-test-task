package clientHandler

import (
	"encoding/json"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"toscale-test-task/firstService/internal/domain/services"
	"toscale-test-task/firstService/internal/infrastructure/controllers/httpController"
	"toscale-test-task/firstService/internal/infrastructure/gateways/gRPCGateway"
)

type ClientHandler struct {
	dataService services.DataService
	log         *logrus.Entry
	gateway     gRPCGateway.GRPCGateway
}

func NewClientHandler(dataService services.DataService, log *logrus.Entry, gateway gRPCGateway.GRPCGateway) httpController.HTTPHandler {
	return &ClientHandler{dataService: dataService, log: log, gateway: gateway}
}

func (cH *ClientHandler) HandlerRouter(g *router.Group) {
	g.GET("/conn", cH.clientRequest)
}

func (cH *ClientHandler) clientRequest(ctx *fasthttp.RequestCtx) {
	fmt.Println("Client Connected")
	args := ctx.QueryArgs()
	symbol := string(args.Peek("symbol"))
	interval := string(args.Peek("interval"))

	klinesC, err := cH.gateway.KlineRequest(ctx, symbol, interval)
	if err != nil {
		err = fmt.Errorf("error gRPC request: %w", err)
		cH.log.Errorln(err)
		ctx.Error("error gRPC request", fasthttp.StatusInternalServerError)
		return
	}

	klines, err := cH.dataService.StoreKline(ctx, klinesC)
	if err != nil {
		err = fmt.Errorf("error storing kline: %w", err)
		cH.log.Errorln(err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	responce, err := json.Marshal(klines)
	if err != nil {
		err = fmt.Errorf("error marshaling json: %w", err)
		cH.log.Errorln(err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

	ctx.Success("application/json", responce)
}
