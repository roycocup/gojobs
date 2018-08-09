package lib

import (
	"fmt"
	"html"
	"net/http"

	"github.com/Sirupsen/logrus"
)

type HTTPServer struct{}

func (this *HTTPServer) Serve(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	logrus.Info("Http server started!")
	logrus.Fatal(http.ListenAndServe(":8282", nil))
}
