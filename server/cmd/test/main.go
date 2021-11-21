package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Out = os.Stdout

	log.WithField("a", "b").Info("c")
	log.Info("Ads")
}
