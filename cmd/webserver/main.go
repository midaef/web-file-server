package main

import (
	"flag"
	"log"
	"packages/internal/app/webserver"
)

var configName string

func init() {
	flag.StringVar(&configName, "config-name", "config.json", "config name")
}

func main() {
	flag.Parse()
	config := webserver.NewConfig()
	config.ConfigInit(configName)
	if err := webserver.Run(config); err != nil {
		log.Fatal(err)
	}
}
