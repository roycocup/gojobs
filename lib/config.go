package lib

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
)

type Config struct{}

func (this *Config) Read(configFileName string) {
	file, err := os.Open(configFileName)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	spew.Dump(file)
}
