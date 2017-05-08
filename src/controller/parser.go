package controller

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
	"rent-parser/src/parser/contact"
	"rent-parser/src/parser/area"
	"rent-parser/src/parser/price"
	parsetype "rent-parser/src/parser/type"
)

type Response struct {
	Type int `json:"type"`
	Contact []string `json:"phone"`
	Area float32 `json:"area"`
	Price int `json:"price"`
}

func Parse(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)

	body := string(ctx.PostBody())

	chan_type := make(chan int)
	chan_contact := make(chan []string)
	chan_area := make(chan float32)
	chan_price := make(chan int)

	go parsetype.Parse(body, chan_type)
	go contact.Parse(body, chan_contact)
	go area.Parse(body, chan_area)
	go price.Parse(body, chan_price)

	response := Response{<-chan_type, <-chan_contact, <-chan_area, <-chan_price}
	json_res, _ := json.Marshal(response)
	ctx.SetBody(json_res)
}
