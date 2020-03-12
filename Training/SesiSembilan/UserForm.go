package SesiSembilan

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	// "github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"tawesoft.co.uk/go/dialog"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> Home")

	if r.Method == "GET" {
		year, month, _ := time.Now().Date()
		var strMonth = strconv.Itoa(int(month))
		if len(strMonth) == 1 {
			strMonth = "0" + strMonth
		}
		var periode string = strconv.Itoa(year) + strMonth

		var articles = GetArticle(periode)
		fmt.Println(articles)

		t, _ := template.ParseFiles("./SesiSembilan/templates/home.html")
		t.Execute(w, articles)

	} else {
		var menu = r.FormValue("submit")
		if menu == "login" {
			http.Redirect(w, r, "/login", 301)
		} else if menu == "home" {
			http.Redirect(w, r, "/home", 301)
		} else if menu == "about" {
			http.Redirect(w, r, "/about", 301)
		} else if menu == "contactus" {
			http.Redirect(w, r, "/contactus", 301)
		}
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> About")

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./SesiSembilan/templates/about.html")
		t.Execute(w, nil)

	} else {
		var menu = r.FormValue("submit")
		if menu == "login" {
			http.Redirect(w, r, "/login", 301)
		} else if menu == "home" {
			http.Redirect(w, r, "/home", 301)
		} else if menu == "about" {
			http.Redirect(w, r, "/about", 301)
		} else if menu == "contactus" {
			http.Redirect(w, r, "/contactus", 301)
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> Login")

	fmt.Println("# Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./SesiSembilan/templates/login.html")
		t.Execute(w, nil)
	} else {
		var menu = r.FormValue("submit")
		if menu == "login" {
			http.Redirect(w, r, "/login", 301)
		} else if menu == "home" {
			http.Redirect(w, r, "/home", 301)
		} else if menu == "about" {
			http.Redirect(w, r, "/about", 301)
		} else if menu == "contactus" {
			http.Redirect(w, r, "/contactus", 301)
		} else {
			r.ParseForm()
			var status bool
			status = true

			if len(r.Form["inputUsername"][0]) == 0 {
				dialog.Alert("Username field is empty")
				fmt.Println(">> Username field is empty")
				status = false
			}

			if len(r.Form["inputPassword"][0]) == 0 {
				dialog.Alert("Password field is empty")
				fmt.Println(">> Password field is empty")
				status = false
			}

			var loginData = LoginData{
				UserName: r.Form["inputUsername"][0],
				Password: r.Form["inputPassword"][0],
			}

			// InsertDB()
			// QueryDB()

			var statusLogin bool
			var isadmin string

			statusLogin, isadmin = ValidatioLogin(loginData.UserName, loginData.Password)

			if !statusLogin {
				dialog.Alert("Error Validation")
				fmt.Println(">> Error Validation")
				status = false
			} else {
				fmt.Println(">> Validation OK")

				expiration := time.Now().Add(365 * 24 * time.Hour)
				cookie1 := http.Cookie{Name: "Username", Value: loginData.UserName, Expires: expiration}
				cookie2 := http.Cookie{Name: "Password", Value: loginData.Password, Expires: expiration}
				cookie3 := http.Cookie{Name: "IsAdmin", Value: isadmin, Expires: expiration}
				http.SetCookie(w, &cookie1)
				http.SetCookie(w, &cookie2)
				http.SetCookie(w, &cookie3)

				http.Redirect(w, r, "/mainmenu", 301)
			}

			if status == false {
				t, _ := template.ParseFiles("./SesiSembilan/templates/login.html")
				t.Execute(w, loginData)
			}
		}
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("# Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./SesiSembilan/templates/register.html")
		t.Execute(w, nil)
	} else {
		var menu = r.FormValue("submit")
		if menu == "login" {
			http.Redirect(w, r, "/login", 301)
		} else if menu == "home" {
			http.Redirect(w, r, "/home", 301)
		} else if menu == "about" {
			http.Redirect(w, r, "/about", 301)
		} else if menu == "contactus" {
			http.Redirect(w, r, "/contactus", 301)
		} else {
			r.ParseForm()

			var status bool
			status = true

			if len(r.Form["inputUsername"][0]) == 0 {
				dialog.Alert("Username field is empty")
				fmt.Println(">> Username field is empty")
				status = false
			}

			if len(r.Form["inputPassword"][0]) == 0 {
				dialog.Alert("Password field is empty")
				fmt.Println(">> Password field is empty")
				status = false
			}

			if len(r.Form["inputRepeatPassword"][0]) == 0 {
				dialog.Alert("Repeate Password field is empty")
				fmt.Println(">> Repeate Password field is empty")
				status = false
			}

			if r.Form["inputPassword"][0] != r.Form["inputRepeatPassword"][0] {
				dialog.Alert("Password not match")
				fmt.Println(">> Password not match")
				status = false
			}

			var loginData = LoginData{
				UserName: r.Form["inputUsername"][0],
				Password: r.Form["inputPassword"][0],
			}

			if status == false {
				t, _ := template.ParseFiles("./SesiSembilan/templates/register.html")
				t.Execute(w, loginData)
			}

			RegisterUser(loginData.UserName, loginData.Password)

			if status == true {
				t, _ := template.ParseFiles("./SesiSembilan/templates/register.html")
				t.Execute(w, nil)
			}
		}
	}
}

func mainmenu(w http.ResponseWriter, r *http.Request) {
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

	strAdmin, err := r.Cookie("IsAdmin")
	if err != nil {
		fmt.Println(err)
	}

	if r.Method == "GET" {
		var loginData = LoginData{
			UserName: strUsername.Value,
			Password: strPassword.Value,
			IsAdmin:  strAdmin.Value,
		}

		fmt.Println(loginData)

		if strAdmin.Value == "0" {
			t, _ := template.ParseFiles("./SesiSembilan/templates/menuuser.html")
			t.Execute(w, loginData)
		} else {
			var user = GetUser()
			// fmt.Println(user)

			t, _ := template.ParseFiles("./SesiSembilan/templates/menuadmin.html")
			t.Execute(w, user)
		}
	} else {
		var menu = r.FormValue("submit")
		if menu == "logout" {
			fmt.Println(">> Logout")
			c1 := &http.Cookie{
				Name:     "Username",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
			}

			c2 := &http.Cookie{
				Name:     "Password",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
			}

			c3 := &http.Cookie{
				Name:     "IsAdmin",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
			}

			http.SetCookie(w, c1)
			http.SetCookie(w, c2)
			http.SetCookie(w, c3)

			http.Redirect(w, r, "/home", 301)
		} else if menu == "mainmenu" {
			http.Redirect(w, r, "/mainmenu", 301)
		} else if menu == "article" {
			http.Redirect(w, r, "/article", 301)
		} else if menu == "update" {
			UpdateUser(r.Form["inputUsername"][0], r.Form["inputPassword"][0], r.Form["inputAdmin"][0])

			var loginData = LoginData{
				UserName: r.Form["inputUsername"][0],
				Password: r.Form["inputPassword"][0],
				IsAdmin:  r.Form["inputAdmin"][0],
			}

			t, _ := template.ParseFiles("./SesiSembilan/templates/menuuser.html")
			t.Execute(w, loginData)
		} else {
			var strUser = r.FormValue("submit")
			// fmt.Println(strUser)

			var user = GetUserbyName(strUser)
			var loginData = LoginData{
				UserName: user[0].UserName,
				Password: user[0].Password,
				IsAdmin:  user[0].IsAdmin,
			}

			t, _ := template.ParseFiles("./SesiSembilan/templates/menuuser.html")
			t.Execute(w, loginData)
		}
	}
}

func article(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> Article")

	if r.Method == "GET" {
		var articles = GetArticle("")
		// var article = Article{
		// 	Subject: r.Form["inputJudul"][0],
		// 	News:    r.Form["inputIsi"][0],
		// 	Periode: r.Form["inputPeriode"][0],
		// }

		t, _ := template.ParseFiles("./SesiSembilan/templates/article.html")
		t.Execute(w, articles)
	} else {
		var menu = r.FormValue("submit")
		if menu == "logout" {
			fmt.Println(">> Logout")
			c1 := &http.Cookie{
				Name:     "Username",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
			}

			c2 := &http.Cookie{
				Name:     "Password",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
			}

			c3 := &http.Cookie{
				Name:     "IsAdmin",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
			}

			http.SetCookie(w, c1)
			http.SetCookie(w, c2)
			http.SetCookie(w, c3)

			http.Redirect(w, r, "/home", 301)
		} else if menu == "mainmenu" {
			http.Redirect(w, r, "/mainmenu", 301)
		} else if menu == "article" {
			http.Redirect(w, r, "/article", 301)
		} else if menu == "add" {
			fmt.Println(">> Add Article")

			AddArticle(r.Form["inputJudul"][0], r.Form["inputIsi"][0], r.Form["inputPeriode"][0])

			var articles = GetArticle("")
			// var article = Article{
			// 	Subject: r.Form["inputJudul"][0],
			// 	News:    r.Form["inputIsi"][0],
			// 	Periode: r.Form["inputPeriode"][0],
			// }

			t, _ := template.ParseFiles("./SesiSembilan/templates/article.html")
			t.Execute(w, articles)
		}
	}
}

func AddArticle(strJudul string, strIsi string, strPeriode string) {
	db, err := sql.Open("mysql", "root:root@tcp(10.0.12.103:3306)/goles?charset=utf8")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("insert into `content` values (?, ?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec(nil, strJudul, strIsi, strPeriode)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	dialog.Alert("Sukses")

	defer db.Close()
}

func GetArticle(periode string) (articles []Article) {
	db, err := sql.Open("mysql", "root:root@tcp(10.0.12.103:3306)/goles?charset=utf8")
	checkErr(err)

	fmt.Println(periode)

	var query string = "select subject,news,periode from content "
	if query != "" && periode != "" {
		query = query + "where periode = '" + periode + "'"
	}

	// query
	rows, err := db.Query(query)
	checkErr(err)

	for rows.Next() {
		var subject string
		var news string
		var periode string
		err = rows.Scan(&subject, &news, &periode)
		checkErr(err)

		var article = Article{
			Subject: subject,
			News:    news,
			Periode: periode,
		}

		articles = append(articles, article)
	}

	defer rows.Close()
	defer db.Close()

	return articles
}

func GetUser() (user []User) {
	db, err := sql.Open("mysql", "root:root@tcp(10.0.12.103:3306)/goles?charset=utf8")
	checkErr(err)

	// query
	rows, err := db.Query("select username from userinfo")
	checkErr(err)

	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		checkErr(err)

		var usr = User{
			Username: username,
		}

		user = append(user, usr)
	}

	defer rows.Close()
	defer db.Close()

	return user
}

func GetUserbyName(strUser string) (user []LoginData) {
	db, err := sql.Open("mysql", "root:root@tcp(10.0.12.103:3306)/goles?charset=utf8")
	checkErr(err)

	// query
	rows, err := db.Query("select username,password,isadmin from userinfo where username ='" + strUser + "'")
	checkErr(err)

	for rows.Next() {
		var username string
		var password string
		var isadmin string
		err = rows.Scan(&username, &password, &isadmin)
		checkErr(err)

		var usr = LoginData{
			UserName: username,
			Password: password,
			IsAdmin:  isadmin,
		}

		user = append(user, usr)
	}

	defer rows.Close()
	defer db.Close()

	return user
}

func ValidatioLogin(strUsername string, strPassword string) (status bool, isadmin string) {
	status = false
	isadmin = "0"

	db, err := sql.Open("mysql", "root:root@tcp(10.0.12.103:3306)/goles?charset=utf8")
	checkErr(err)

	var query = "select username,password,isadmin from userinfo where username ='" + strUsername + "' and password ='" + strPassword + "'"
	// fmt.Println(query)

	// query
	rows, err := db.Query(query)
	checkErr(err)
	// fmt.Println("rows ", rows)

	for rows.Next() {
		var username string
		var password string
		err = rows.Scan(&username, &password, &isadmin)
		checkErr(err)
		// fmt.Println(username, password, isadmin)

		status = true
	}

	defer rows.Close()
	defer db.Close()

	return status, isadmin
}

func RegisterUser(strUsername string, strPassword string) {
	currentTime := time.Now()
	var strCurrentTime = currentTime.Format("2006-01-02")

	db, err := sql.Open("mysql", "root:root@tcp(10.0.12.103:3306)/goles?charset=utf8")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("insert into `userinfo` values (?, ?, ?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec(strUsername, strPassword, "User", "0", strCurrentTime)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	dialog.Alert("Sukses")

	defer db.Close()
}

func UpdateUser(strUsername string, strPassword string, strIsAdmin string) {
	// currentTime := time.Now()
	// var strCurrentTime = currentTime.Format("2006-01-02")

	db, err := sql.Open("mysql", "root:root@tcp(10.0.12.103:3306)/goles?charset=utf8")
	checkErr(err)

	// update
	stmt, err := db.Prepare("update `userinfo` set username=?, password=?, isadmin=? where username=?")
	checkErr(err)

	res, err := stmt.Exec(strUsername, strPassword, strIsAdmin, strUsername)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	dialog.Alert("Sukses")

	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func SesiSembilanUserForm() {
	http.HandleFunc("/home", home)           // set router
	http.HandleFunc("/about", about)         // set router
	http.HandleFunc("/login", login)         // set router
	http.HandleFunc("/mainmenu", mainmenu)   // set router
	http.HandleFunc("/register", register)   // set router
	http.HandleFunc("/article", article)     // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}
