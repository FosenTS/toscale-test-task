package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"toscale-test-task/first-service/internal/application"
)

func main() {
	log := logrus.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := application.NewApp(ctx, log.WithField("location", "application"))
	if err != nil {
		log.Fatalln("fatal error creating application")
		return
	}
}
