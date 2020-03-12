package SesiSembilan

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
	// "github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"tawesoft.co.uk/go/dialog"
)

func login1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("# Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./SesiSembilan/templates/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		if len(r.Form["username"][0]) == 0 {
			dialog.Alert("Username field is empty")
			fmt.Println(">> Username field is empty")
			return
		}

		if len(r.Form["password"][0]) == 0 {
			dialog.Alert("Password field is empty")
			fmt.Println(">> Password field is empty")
			return
		}

		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("username")); !m {
			dialog.Alert("Error No Alphabet")
			fmt.Println(">> Error No Alphabet")
			return
		}

		_, err := strconv.Atoi(r.Form.Get("password"))
		if err != nil {
			dialog.Alert("Error when convert to number, it may not a number")
			fmt.Println(">> Error when convert to number, it may not a number")
			return
		}

		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("password")); !m {
			dialog.Alert("Error negative number")
			fmt.Println(">> Error negative number")
			return
		}

		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			dialog.Alert("email not valid")
			fmt.Println(">> email not valid")
			return
		}

		if !isFruit1(r) {
			dialog.Alert("Error Fruit")
			fmt.Println(">> Error Fruit")
			return
		}

		if !isGender1(r) {
			dialog.Alert("Error Gender")
			fmt.Println(">> Error Gender")
			return
		}

		var loginData = LoginData{
			UserName: r.Form["username"][0],
			Password: r.Form["password"][0],
		}

		// InsertDB()
		QueryDB1()

		if !ValidatioLogin1(loginData.UserName, loginData.Password) {
			fmt.Println(">> Error Validation")
			return
		}

		fmt.Println(">> Validation OK")

		// fmt.Println("Username: ", loginData.UserName)
		// fmt.Println("Password: ", loginData.Password)

		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie1 := http.Cookie{Name: "Username", Value: loginData.UserName, Expires: expiration}
		cookie2 := http.Cookie{Name: "Password", Value: loginData.Password, Expires: expiration}
		http.SetCookie(w, &cookie1)
		http.SetCookie(w, &cookie2)

		// http.Redirect(w, r, "http://www.google.com", 301)
		http.Redirect(w, r, "/mainmenu", 301)
	}
}

func mainmenu1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> Main Menu")

	// r.ParseForm()
	// fmt.Printf("%+v", r)

	strUsername, err := r.Cookie("Username")
	if err != nil {
		fmt.Println(err)
	}

	strPassword, err := r.Cookie("Password")
	if err != nil {
		fmt.Println(err)
	}

	var loginData = LoginData{
		UserName: strUsername.Value,
		Password: strPassword.Value,
	}

	// fmt.Println("Username: ", loginData.UserName)
	// fmt.Println("Password: ", loginData.Password)

	t, _ := template.ParseFiles("./SesiSembilan/templates/mainmenu.html")
	t.Execute(w, loginData)
}

func errorMsg1(w http.ResponseWriter, r *http.Request) {
	var err Error

	t, _ := template.ParseFiles("./SesiSembilan/templates/error.html")
	t.Execute(w, err)
}

func isFruit1(r *http.Request) bool {
	fruits := []string{"apple", "banana", "pear"}

	for _, v := range fruits {
		if v == r.Form.Get("fruit") {
			return true
		}
	}
	return false
}

func isGender1(r *http.Request) bool {
	gender := []string{"1", "2"}

	for _, v := range gender {
		if v == r.Form.Get("gender") {
			return true
		}
	}
	return false
}

func InsertDB1() {
	db, err := sql.Open("mysql", "root:root@tcp(10.0.12.103:3306)/goles?charset=utf8")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("insert into `userinfo` values (?, ?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec("heru", "123", "IT", "2020-01-01")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	defer db.Close()
}

func QueryDB1() {
	db, err := sql.Open("mysql", "root:root@tcp(10.0.12.103:3306)/goles?charset=utf8")
	checkErr(err)

	// query
	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var username string
		var password string
		var department string
		var created string
		err = rows.Scan(&username, &password, &department, &created)
		checkErr(err)

		fmt.Println(username)
		fmt.Println(password)
		fmt.Println(department)
		fmt.Println(created)
	}

	defer rows.Close()
	defer db.Close()
}

func ValidatioLogin1(strUsername string, strPassword string) bool {
	var status bool
	status = false

	db, err := sql.Open("mysql", "root:root@tcp(10.0.12.103:3306)/goles?charset=utf8")
	checkErr(err)

	var query = "select * from userinfo where username ='" + strUsername + "' and password ='" + strPassword + "'"
	// fmt.Println(query)

	// query
	rows, err := db.Query(query)
	checkErr(err)

	for rows.Next() {
		var username string
		var password string
		var department string
		var created string
		err = rows.Scan(&username, &password, &department, &created)
		checkErr(err)

		status = true
	}

	defer rows.Close()
	defer db.Close()

	return status
}

func checkErr1(err error) {
	if err != nil {
		panic(err)
	}
}

func SesiSembilanUserForm1() {
	http.HandleFunc("/login", login)         // set router
	http.HandleFunc("/mainmenu", mainmenu)   // set router
	http.HandleFunc("/error", errorMsg1)     // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}
