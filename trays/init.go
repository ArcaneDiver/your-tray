package trays

import (
	"github.com/ArcaneDiver/your-tray/log"
	"github.com/pkg/errors"
	"time"
)

func Init(config *Config) {
	log.Log.Info("Initializing tray")
	log.Log.WithFields(map[string]interface{}{
		"name": 	config.Tray.Name,
		"icon":		config.Tray.Icon,
		"tooltip":	config.Tray.Tooltip,
	}).Debug()

	syncRegister := make(chan bool, 1)

	config.Tray.Init(syncRegister)

	<-syncRegister
	log.Log.Info("Tray initialized")

	for {
		for _, item := range config.Tray.Items {
			if !item.IsDynamic() {
				continue
			}

			if err := item.HandleTick(); err != nil {
				log.Log.WithField("item", item.Name).Error(errors.Wrap(err, "trays.Init"))
			}
		}

		time.Sleep(time.Second * 3)
	}

}
