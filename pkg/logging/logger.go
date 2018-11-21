package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var root = logrus.New()

func init() {
	if os.Getenv("DEBUG") != "" {
		root.SetLevel(logrus.DebugLevel)
	}
}

func GetRoot() *logrus.Logger {
	return root
}
