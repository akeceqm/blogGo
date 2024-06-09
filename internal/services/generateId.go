package services

func GenerateId() string {
	return GenerateString(idLenght, idSet)
}
