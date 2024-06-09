package middlewares

<<<<<<< Updated upstream
import "golang.org/x/crypto/bcrypt"
=======
import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) []byte {

	if password == "" {
		fmt.Println("password пустая ")

	}
>>>>>>> Stashed changes

func PasswordHash(password string) []byte {
	passwordHash := []byte(password)
	cost := 12
	hash, _ := bcrypt.GenerateFromPassword(passwordHash, cost)
	return hash
}
