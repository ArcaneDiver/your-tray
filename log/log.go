package log

import "github.com/sirupsen/logrus"

var Log = logrus.New()

func Init(level string) {
	Log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		Log.Panicf("Level: %s is not a valid log level", level)
	}

	Log.SetLevel(parsedLevel)
}
