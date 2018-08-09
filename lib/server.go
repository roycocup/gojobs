package lib

import (
	"fmt"
	"html"
	"net/http"

	"github.com/Sirupsen/logrus"
)

type HTTPServer struct{}

func (this HTTPServer) Serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	logrus.Fatal(http.ListenAndServe(":8080", nil))
}
