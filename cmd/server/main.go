package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/microdimmer/testing_api/internal/app/rest_server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/rest_server.toml", "path to config file")
	// flag.StringVar(&configPath, "config-path", "../../configs/rest_server.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := rest_server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	serv := rest_server.New(config)

	if err := serv.Run(); err != nil {
		log.Fatal(err)
	}
}
