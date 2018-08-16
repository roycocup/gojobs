package lib

import (
	"github.com/Sirupsen/logrus"
)

/*
Custom error class
*/
type Error struct{}

func NewErrorHelper() *Error{
	return &Error{}
}

func (client *Error) Check(err error) {
	if err != nil {
		logrus.Warn(err.Error())
	}
}

func (client *Error) CheckWithMessage(err error, message string) {
	if err != nil {
		logrus.WithError(err).Warn(message)
	}
}

func (client *Error) Fatal(err error) {
	if err != nil {
		logrus.Fatal(err.Error())
	}
}

func (client *Error) FatalWithMessage(err error, message string) {
	if err != nil {
		logrus.WithError(err).Fatal(message)
	}
}
