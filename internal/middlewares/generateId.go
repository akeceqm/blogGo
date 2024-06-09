package middlewares

func GenerateId() string {
	return GenerateString(idLenght, idSet)
}
