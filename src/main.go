package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"controller"
	"config"
	"os"
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/parse":
		controller.Parse(ctx)
		break
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

func main() {

	conf, err := config.Read()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fasthttp.ListenAndServe(fmt.Sprintf("%s:%s", conf["server_host"], conf["server_port"]), requestHandler)
}
