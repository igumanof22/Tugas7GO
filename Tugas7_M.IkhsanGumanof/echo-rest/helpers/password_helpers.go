package helpers

import "golang.org/x/crypto/bcrypt"

//HashPassword ...
func HashPassword(password string) (string, error) {

	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(pass), err
}

func CheckPasswordHash(password, hash string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {

		return false, nil

	}

	return true, nil
}
