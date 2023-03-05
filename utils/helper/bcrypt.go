package helper

import "golang.org/x/crypto/bcrypt"

func PasswordHashing(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err.Error()
	}
	return string(hashedPassword)
}
