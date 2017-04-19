package controller

import (
	"github.com/valyala/fasthttp"
	"parser"
	"fmt"
)

func parseType(par *parser.Parser, body string, req chan string) {
	r := par.ParseType(body)
	req <- r
}

func parseArea(par *parser.Parser, body string, req chan string) {
	r := par.ParseArea(body)
	req <- r
}

func parseContact(par *parser.Parser, body string, req chan string) {
	r := par.ParseContact(body)
	req <- r
}

func parsePrice(par *parser.Parser, body string, req chan string) {
	r := par.ParsePrice(body)
	req <- r
}
func Parse(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)

	body := string(ctx.PostBody())

	chan_type := make(chan string)
	chan_contact := make(chan string)
	chan_area := make(chan string)
	chan_price := make(chan string)

	p := parser.NewParser()

	go parseType(p, body, chan_type)
	go parseContact(p, body, chan_contact)
	go parseArea(p, body, chan_area)
	go parsePrice(p, body, chan_price)

	// then override already written body
	ctx.SetBody([]byte(fmt.Sprint(<-chan_type, <-chan_contact, <-chan_area, <-chan_price)))
}
