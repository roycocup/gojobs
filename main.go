package main

import (
	"os"

	"./lib"
	"github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

func main() {
	log := logrus.New()
	log.Info("Process started!")

	eChecker := lib.Error{}

	// read config file
	err := config.Load(file.NewSource(file.WithPath("./config.yml")))
	eChecker.Check(err)

	port, err := cfg.Get("site.port")
	spew.Dump(port, err)
	os.Exit(100)

	// httpserver := lib.HTTPServer{}
	// httpserver.Serve(port)
	// serve http on separate thread
	// init scrapers (one per source)
	// save data
}
