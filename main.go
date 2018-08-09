package main

import (
	"strconv"

	"./lib"
	"github.com/Sirupsen/logrus"
	"github.com/tsuru/config"
)

func main() {
	log := logrus.New()
	log.Info("Process started!")
	// read config file
	config.ReadConfigFile("./config.yml")
	portStr, _ := config.GetString("site_port")
	port, _ := strconv.Atoi(portStr)

	httpserver := lib.HTTPServer{}
	httpserver.Serve(port)
	// serve http on separate thread
	// init scrapers (one per source)
	// save data
}
