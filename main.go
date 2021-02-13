package main

import (
	"fmt"
	"github.com/ArcaneDiver/your-tray/config"
	"github.com/ArcaneDiver/your-tray/log"
	"github.com/getlantern/systray"
	"github.com/sirupsen/logrus"
)

func main() {
	log.Init()

	configs, err := config.Parse("config.yaml")
	if err != nil {
		log.Log.Error(err)
		logrus.Exit(1)
	}

	fmt.Print(configs)
}

func onStart() {
	systray.SetTitle("Lallo")
	systray.SetTooltip("Lallo claccami")
	mClick := systray.AddMenuItem("Lelloo", "toolatpp")

	go func() {
		for {
			select {
				case <- mClick.ClickedCh: {
					log.Log.Info("Clicckckkkk")
				}
			}
		}
	}()
}

func onExit() {
	log.Log.Info("Outttt")
}