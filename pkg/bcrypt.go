package pkg

import "golang.org/x/crypto/bcrypt"

func BcryptHash(password string) (hashedPass string, err error) {
	passByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	hashedPass = string(passByte)
	return
}

func BcryptCompare(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
