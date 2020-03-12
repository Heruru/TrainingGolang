package SesiSembilan

type LoginData struct {
	UserName string
	Password string
	IsAdmin  string
}

type User struct {
	Username string
}

type Article struct {
	Subject string
	News    string
	Periode string
}

type Error struct {
	Error string
}
