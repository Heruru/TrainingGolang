package SesiTujuh

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
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

func html(w http.ResponseWriter, r *http.Request) {
	const tpl = `
        <!DOCTYPE html>
        <html>
            <head>
                <meta charset="UTF-8">
                <title>{{.Title}}</title>
            </head>
            <body>
                {{$len := len .Items}}
                {{if lt $len 3}}
                    bisa
                {{else}}
                    tidak
                {{end}}
                {{range .Items}}
                    <div>
                        {{ . }}
                    </div>
                {{else}}
                    <div>
                        <strong>Tidak ada menu</strong>
                    </div>
                {{end}}
            </body>
        </html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("index").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(w, data)
	check(err)
}

func SesiTujuhWebServerHtml() {
	http.HandleFunc("/", index)              // set router
	http.HandleFunc("/html", html)           // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}
