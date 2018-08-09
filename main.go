package main

import (
	"./lib"
	"github.com/Sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Info("Process started!")
	// read config file
	httpserver := lib.HTTPServer{}
	httpserver.Serve()
	// serve http on separate thread
	// init scrapers (one per source)
	// save data
}
