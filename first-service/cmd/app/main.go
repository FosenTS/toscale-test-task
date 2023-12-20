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
	app, err := application.NewApp(ctx, log.WithField("location", "application"))
	if err != nil {
		log.Fatalln("fatal error creating application")
		return
	}
	err = app.Run(ctx, log.WithField("location", "runner"))
	if err != nil {
		log.Fatalln("fatal run application")
		return
	}
}
