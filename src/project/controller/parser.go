package controller

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
	"project/parser/contact"
	"project/parser/price"
	"project/parser/area"
	parsetype "project/parser/type"
)

type Response struct {
	Type []string `json:"type"`
	Contact []string `json:"contact"`
	Area []string `json:"area"`
	Price []string `json:"price"`
}

func Parse(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)

	body := string(ctx.PostBody())

	chan_type := make(chan []string)
	chan_contact := make(chan []string)
	chan_area := make(chan []string)
	chan_price := make(chan []string)

	go parsetype.Parse(body, chan_type)
	go contact.Parse(body, chan_contact)
	go area.Parse(body, chan_area)
	go price.Parse(body, chan_price)

	response := Response{<-chan_type, <-chan_contact, <-chan_area, <-chan_price}
	json_res, _ := json.Marshal(response)
	ctx.SetBody(json_res)
}
