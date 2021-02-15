package main

import (
	"flag"
	"github.com/ArcaneDiver/your-tray/log"
	"github.com/ArcaneDiver/your-tray/trays"
	"github.com/sirupsen/logrus"
)

func main() {
	configPath 	:= flag.String("config", "/etc/your-tray/config.yaml", "Path to the configuration")
	logLevel 	:= flag.String("level", "debug", "Log levels: error, warn, info, debug")
	flag.Parse()

	log.Init(*logLevel)

	configs, err := trays.Parse(*configPath)
	if err != nil {
		log.Log.Error(err)
		logrus.Exit(1)
	}

	trays.Init(configs)
}