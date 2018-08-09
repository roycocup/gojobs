package main

import (
	"./lib"
	"github.com/Sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Info("Process started!")
	// read config file
	config := lib.Config{}
	config.Read("./config.yml")

	httpserver := lib.HTTPServer{}
	httpserver.Serve(8282)
	// serve http on separate thread
	// init scrapers (one per source)
	// save data
}
