package models

type User struct {
	Name     string
	Email    string
	Password string
	Ip       string
}

type IPResponse struct {
	IP string
}
