package parser

import (
	"os/exec"
	"bytes"
	"log"
	"strings"
)

type Parser struct {
	bin string
	config_area string
	config_contact string
	config_price string
	config_type string
}

func NewParser() *Parser {
	p := new(Parser)

	p.bin = "/Users/newuser/web/rent-tomita/bin/tomita"
	p.config_area = "/Users/newuser/web/rent-tomita/tomita/area/config.proto"
	p.config_contact = "/Users/newuser/web/rent-tomita/tomita/contact/config.proto"
	p.config_price = "/Users/newuser/web/rent-tomita/tomita/price/config.proto"
	p.config_type = "/Users/newuser/web/rent-tomita/tomita/type/config.proto"

	return p
}

func (parser Parser) ParseArea(arg string) string {
	command := exec.Command(parser.bin, parser.config_area)
	var out bytes.Buffer
	command.Stdin = strings.NewReader(arg)
	command.Stdout = &out

	err := command.Run()
	command.Start()

	if err != nil {
		log.Fatal(err)
	}

	return out.String()
}

func (parser Parser) ParseContact(arg string) string {
	command := exec.Command(parser.bin, parser.config_contact)
	var out bytes.Buffer
	command.Stdin = strings.NewReader(arg)
	command.Stdout = &out

	err := command.Run()

	if err != nil {
		log.Fatal(err)
	}

	return out.String()
}

func (parser Parser) ParsePrice(arg string) string {
	command := exec.Command(parser.bin, parser.config_price)
	var out bytes.Buffer
	command.Stdin = strings.NewReader(arg)
	command.Stdout = &out

	err := command.Run()

	if err != nil {
		log.Fatal(err)
	}

	return out.String()
}

func (parser Parser) ParseType(arg string) string {
	command := exec.Command(parser.bin, parser.config_type)
	var out bytes.Buffer
	command.Stdin = strings.NewReader(arg)
	command.Stdout = &out

	err := command.Run()

	if err != nil {
		log.Fatal(err)
	}

	return out.String()
}
