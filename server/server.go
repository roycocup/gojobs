package server

import (
	"net/http"

	"io/ioutil"

	"github.com/Sirupsen/logrus"
)

func Serve(port int) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile("web/index.html")
		if err != nil {
			logrus.Fatal(err.Error())
		}

		w.Write(data)
	})

	logrus.Info("Http server started!")
	logrus.Fatal(http.ListenAndServe(":8282", nil))
}
