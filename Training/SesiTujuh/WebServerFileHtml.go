package SesiTujuh

import (
	"html/template"
	"log"
	"net/http"
)

func htmlFile(w http.ResponseWriter, r *http.Request) {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.ParseFiles("./SesiTujuh/templates/base.html")
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos file",
			"My blog file",
		},
	}

	err = t.Execute(w, data)
	check(err)
}

func SesiTujuhWebServerHtmlFile() {
	http.HandleFunc("/html/file", htmlFile)  // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}
