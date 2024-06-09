package services

func GeneratePassword() string {
	return GenerateString(passwordLenght, passwordSet)
}
