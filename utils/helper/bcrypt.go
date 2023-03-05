package helper

import "golang.org/x/crypto/bcrypt"

func PasswordHashing(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return []byte(err.Error()), err
	}
	return hashedPassword, nil
}

func ComparePassword(hashedPassword, password []byte) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
