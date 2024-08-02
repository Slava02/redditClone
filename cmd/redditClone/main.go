package main

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"redditClone/internal/app"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	log.SetOutput(ioutil.Discard)

	cfg := app.MustLoad()
	//app.PrintConf(cfg)

	app.Run(*cfg)
}
