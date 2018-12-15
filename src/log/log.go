package log

import (
	"bytes"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Init ...
func Init(level string) {
	lv := string(bytes.ToLower([]byte(level)))
	switch lv {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano})
	logrus.SetOutput(os.Stdout)
}
