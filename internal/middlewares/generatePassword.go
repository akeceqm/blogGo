package middlewares

func GeneratePassword() string {
	return GenerateString(passwordLenght, passwordSet)
}
