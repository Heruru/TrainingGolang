package SesiSembilan

import (
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Ikhsan!")
}

func SesiSembilanRouting() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
