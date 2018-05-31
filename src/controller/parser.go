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

	if !parsetype.PreValid(body) {
		response := Response{parsetype.WRONG, -1}
		json_res, _ := json.Marshal(response)
		ctx.SetBody(json_res)

		return
	}

	chanType := make(chan int)
	chanPrice := make(chan int)

	go parsetype.Parse(body, chanType)
	go price.Parse(body, chanPrice)

	response := Response{<-chanType, <-chanPrice}
	json_res, _ := json.Marshal(response)
	ctx.SetBody(json_res)
}
