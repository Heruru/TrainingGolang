package SesiDelapan

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func readJson() StsDisplay {
	var status Sts
	var statusDisplay StsDisplay

	// Open our jsonFile
	jsonFile, err := os.Open("./SesiDelapan/templates/hasil.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return statusDisplay
	}

	fmt.Println("Successfully Opened hasil.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return statusDisplay
	}

	// fmt.Println("json byte", string(byteValue))

	json.Unmarshal(byteValue, &status)

	statusDisplay = StsDisplay{
		Status: status.Status,
	}

	if status.Status.Water < 5 {
		statusDisplay.StatusWater = "Aman"
	} else if status.Status.Water >= 5 && status.Status.Water <= 9 {
		statusDisplay.StatusWater = "Siaga"
	} else {
		statusDisplay.StatusWater = "Bahaya"
	}

	if status.Status.Wind < 6 {
		statusDisplay.StatusWind = "Aman"
	} else if status.Status.Wind >= 6 && status.Status.Wind <= 15 {
		statusDisplay.StatusWind = "Siaga"
	} else {
		statusDisplay.StatusWind = "Bahaya"
	}

	return statusDisplay
}

func writeJson() {
	var status = Sts{
		Status: Status{
			Water: random(1, 100),
			Wind:  random(1, 100),
		},
	}

	file, err := json.MarshalIndent(status, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile("./SesiDelapan/templates/hasil.json", file, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully Update hasil.json")

	water := strconv.Itoa(status.Status.Water)
	wind := strconv.Itoa(status.Status.Wind)

	fmt.Println("Water: " + water)
	fmt.Println("Wind: " + wind)
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func ShowHasil(w http.ResponseWriter, r *http.Request) {
	status := readJson()
	// fmt.Printf("%+v\n", status)

	const tpl = `
        <!DOCTYPE html>
        <html>
            <head>
				<meta charset="UTF-8">
				<meta http-equiv="refresh" content="5">				
                <title>Status</title>
            </head>
            <body>
                    <div>
						Water: {{ .Status.Water }} 
						<br>
						Status:  {{ .StatusWater }} 
						<br>
						<br>
						Wind: {{ .Status.Wind }}
						<br>
						Status:  {{ .StatusWind }}
						<br>
						</br>
                    </div>
            </body>
        </html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("index").Parse(tpl)
	check(err)

	err = t.Execute(w, status)
	check(err)
}

func SesiDelapanReviewandExercise() {
	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				writeJson()
			}
		}
	}()

	go func() {
		http.HandleFunc("/status", ShowHasil)    // set router
		err := http.ListenAndServe(":9090", nil) // set listen port
		if err != nil {
			log.Fatal("Error running service: ", err)
		}
	}()

	time.Sleep(1 * time.Minute)
	ticker.Stop()
	done <- true
	fmt.Println("Stopped")
}
