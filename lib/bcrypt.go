package lib

import "golang.org/x/crypto/bcrypt"

type IHasher interface {
	HashPassword(password string) string
	PasswordIsCorrect(hashedPassword string, password string) bool
}

type Hasher struct{}

// HashPassword implements HasherInterface.
func (*Hasher) HashPassword(password string) string {
	var hashedPasswordBytes []byte
	var hashedPassword string

	if len(password) > 0 {
		hashedPasswordBytes, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		hashedPassword = string(hashedPasswordBytes)
	}

	return hashedPassword
}

// PasswordIsCorrect implements HasherInterface.
func (*Hasher) PasswordIsCorrect(hashedPassword string, password string) bool {
	passwordIsCorrect := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return passwordIsCorrect == nil
}

func NewHasher() IHasher {
	return &Hasher{}
}
