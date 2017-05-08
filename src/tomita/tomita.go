package tomita

import (
	"os/exec"
	"bytes"
	"log"
	"strings"
)

type Tomita struct {
	bin string
	config string
}

func NewTomita(bin string, config string) *Tomita {
	p := new(Tomita)

	p.bin = bin
	p.config = config

	return p
}

func (tomita Tomita) Parse(text string) string {
	command := exec.Command(tomita.bin, tomita.config)
	var Stdout bytes.Buffer
	var Stderr bytes.Buffer
	command.Stdin = strings.NewReader(text)
	command.Stdout = &Stdout
	command.Stderr = &Stderr

	err := command.Run()
	command.Start()

	if err != nil {
		log.Fatal(err)
	}

	return Stdout.String()
}