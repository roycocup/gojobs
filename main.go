package main

import (
	"./lib"
	"github.com/Sirupsen/logrus"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

var (
	conf config.Config
	log *logrus.Logger
	eChecker *lib.Error
)

func main() {
	log = logrus.New()
	log.Info("Process started!")
	eChecker = lib.NewErrorHelper()

	// read config file
	loadConfig("./config.yml")

	port := config.Get("site", "port").Int(9000)

	httpserver := lib.HTTPServer{}
	httpserver.Serve(port)
	// serve http on separate thread
	// init scrapers (one per source)
	// save data
}


func loadConfig(configFile string){
	conf = config.NewConfig()
	err := conf.Load(file.NewSource(file.WithPath(configFile)))
	eChecker.Check(err)
}