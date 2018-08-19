package main

import (
	"./errors"
	"./server"
	"github.com/Sirupsen/logrus"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

var (
	conf config.Config
	log  *logrus.Logger
)

func main() {
	log = logrus.New()
	log.Info("Process started!")

	// read config file
	loadConfig("./config.yml")

	port := config.Get("site", "port").Int(9000)
	// serve http on separate thread
	server.Serve(port)

	// init scrapers (one per source)

	// save data
}

func loadConfig(configFile string) {
	conf = config.NewConfig()
	err := conf.Load(file.NewSource(file.WithPath(configFile)))
	errors.Check(err)
}
