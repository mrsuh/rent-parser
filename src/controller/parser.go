package controller

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
	"rent-parser/src/parser/price"
	parsetype "rent-parser/src/parser/type"
)

type Response struct {
	Type int `json:"type"`
	Price int `json:"price"`
}

func Parse(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)

	body := string(ctx.PostBody())

	chan_type := make(chan int)
	chan_price := make(chan int)

	go parsetype.Parse(body, chan_type)
	go price.Parse(body, chan_price)

	response := Response{<-chan_type, <-chan_price}
	json_res, _ := json.Marshal(response)
	ctx.SetBody(json_res)
}
