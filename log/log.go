package log

import "github.com/sirupsen/logrus"

var Log = logrus.New()

func Init() {
	Log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	Log.SetLevel(logrus.DebugLevel)
}
