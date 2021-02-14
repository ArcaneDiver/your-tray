package trays

import (
	"github.com/ArcaneDiver/your-tray/log"
)

func Init(config *Config) {
	log.Log.Info("Initializing tray")
	log.Log.WithFields(map[string]interface{}{
		"name": 	config.Tray.Name,
		"icon":		config.Tray.Icon,
		"tooltip":	config.Tray.Tooltip,
	}).Debug()

	config.Tray.Init()
}
