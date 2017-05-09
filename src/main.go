package main

import (
	"github.com/valyala/fasthttp"
	config "github.com/mrsuh/cli-config"
	"rent-parser/src/controller"
	"fmt"
	"log"
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/parse":

		if !ctx.IsPost() {
			ctx.Error("Method not allowed", fasthttp.StatusMethodNotAllowed)
			break
		}

		controller.Parse(ctx)

		break
	default:
		ctx.Error("Not found", fasthttp.StatusNotFound)
	}
}

func main() {

	conf_instance := config.GetInstance()

	err := conf_instance.Init()

	if err != nil {
		log.Fatal(err)
	}

	conf := conf_instance.Get()

	fmt.Println("init")

	server_err := fasthttp.ListenAndServe(fmt.Sprintf("%s:%s", conf["server.host"], conf["server.port"]), requestHandler)

	if server_err != nil {
		log.Fatal(server_err)
	}
}
