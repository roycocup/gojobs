package main

import (
	"./lib"
	"github.com/Sirupsen/logrus"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

func main() {
	log := logrus.New()
	log.Info("Process started!")

	eChecker := lib.Error{}

	// read config file
	conf := config.NewConfig()
	err := conf.Load(file.NewSource(file.WithPath("./config.yml")))
	eChecker.Check(err)

	port := config.Get("site", "port").Int(9000)

	httpserver := lib.HTTPServer{}
	httpserver.Serve(port)
	// serve http on separate thread
	// init scrapers (one per source)
	// save data
}
