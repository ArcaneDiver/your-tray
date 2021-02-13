package main

import (
	"github.com/ArcaneDiver/your-tray/log"
	"github.com/getlantern/systray"
)

func main() {
	log.Init()
	systray.Run(onStart, onExit)
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