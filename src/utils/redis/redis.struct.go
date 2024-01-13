package redis

type User struct {
	id       int64
	name     string
	email    string
	isActive int
	dob      string
	pob      string
	phoneint string
	address  string
	agency   *[]Agency
}

type Agency struct {
	id                int64
	name              string
	email             string
	noHp              string
	address           string
	logo              string
	domain            string
	subs              *[]Subs
	role              *[]Role
	specialPermission *[]Permission
}

type Role struct {
	id         int64
	name       string
	permission *[]Permission
}

type Permission struct {
	id          int64
	name        string
	description string
}

type Subs struct {
	id    int64
	name  string
	logo  string
	price string
	app   App
}

type App struct {
	id      int64
	name    string
	image   string
	service *[]Service
	setting *[]Setting
}

type Setting struct {
	id    int64
	key   string
	value string
}

type Service struct {
	id       int64
	name     string
	schema   Schema
	hostname string
	username string
	password string
	database string
	port     string
	engine   string
}

type Schema struct {
	id   int64
	name string
}
