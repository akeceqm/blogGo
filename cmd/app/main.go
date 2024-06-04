package main

import "post/internal/services/mainApp"

type UserHandler interface {
	AuthorizationUser() error
	RegistrationUser() error
}

func main() {
	mainApp.MainApp("Добро пожаловать в  Блог")
}
