package main

import (
	"github.com/cruzzan/treat-your-plant/internal/pkg/store/postgres"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	log.SetLevel(logrus.DebugLevel)

	log.Info("Hello, i have nothing to show you yet :(")

	dbUrl := "We don't have one yet..."

	_, err := postgres.New(dbUrl, 10, log)

	if err != nil {
		log.WithError(err).Fatal("Could not set up storage")
	}
}
