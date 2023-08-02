package helpers

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(userPassword string, requestPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(requestPassword))
	return err == nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
