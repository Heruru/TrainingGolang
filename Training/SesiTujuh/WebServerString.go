package SesiTujuh

import (
	"fmt"
	"log"
	"net/http"
)

func indexSring(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hallo World \n") // send data to client side
	fmt.Fprintf(w, "It's Works!!")   // send data to client side
}

func SesiTujuhWebServerString() {
	http.HandleFunc("/hello", indexSring)    // set router
	err := http.ListenAndServe(":9090", nil) // set listen port to 9090
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}
