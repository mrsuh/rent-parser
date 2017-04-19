package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"errors"
)

func Read() (map[string]interface{}, error) {
	args := os.Args
	if len(args) < 2 {
		return nil, errors.New("Config file is required")
	}

	data, read_err := ioutil.ReadFile(args[1])
	if read_err != nil {
		return nil, read_err
	}

	var conf map[string]interface{}

	parse_err := yaml.Unmarshal(data, &conf)
	if parse_err != nil {
		return nil, parse_err
	}

	return conf, nil
}