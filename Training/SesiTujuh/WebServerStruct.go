package SesiTujuh

import (
	"html/template"
	"log"
	"net/http"
)

func indexStruct(w http.ResponseWriter, r *http.Request) {
	daniel := struct {
		Name string
		Age  int
	}{"Daniel Sudibyo", 23}
	tmplt, err := template.New("index").Parse("Nama saya {{.Name}} dan umur saya {{.Age}} tahun")
	if err != nil {
		log.Fatal(err)
	}

	err = tmplt.Execute(w, daniel) // send data to client side
	if err != nil {
		log.Fatal(err)
	}
}

func SesiTujuhWebServerStruct() {
	http.HandleFunc("/", indexStruct)        // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}
