package main

import (
	"github.com/valyala/fasthttp"
	"project/controller"
	"fmt"
	"log"
	"project/config"
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

	//str := "89992144342 89992144347"
	//fmt.Println("init")
	//parser.ParseContact(str)
}
