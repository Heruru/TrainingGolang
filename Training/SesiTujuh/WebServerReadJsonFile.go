package SesiTujuh

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func htmlStudent(w http.ResponseWriter, r *http.Request) {
	// Open our jsonFile
	jsonFile, err := os.Open("./SesiTujuh/templates/student.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened student.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal([]byte(byteValue), &users)

	const tpl = `
        <!DOCTYPE html>
        <html>
            <head>
                <meta charset="UTF-8">
                <title>Student</title>
            </head>
            <body>
                {{range .Users}}
                    <div>
						User Job: {{ .Job }} 
						<br>
						User Age: {{ .Age }}
						<br>
						User Name: {{ .Name }}
						<br>
						Facebook Url: {{ .Social.Facebook }}
						<br>
						Twitter Url: {{ .Social.Twitter }}
						<br>
						</br>
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

	err = t.Execute(w, users)
	check(err)
}

func student(w http.ResponseWriter, r *http.Request) {
	// Open our jsonFile
	jsonFile, err := os.Open("./SesiTujuh/templates/student.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened student.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal([]byte(byteValue), &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Job: " + users.Users[i].Job)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
		fmt.Println("Twitter Url: " + users.Users[i].Social.Twitter)
	}
}

func SesiTujuhWebServerReadJsonFile() {
	http.HandleFunc("/student", student)          // set router
	http.HandleFunc("/student/html", htmlStudent) // set router
	err := http.ListenAndServe(":9090", nil)      // set listen port
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}
