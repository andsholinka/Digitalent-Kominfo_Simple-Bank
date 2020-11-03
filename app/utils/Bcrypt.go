package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashGenerator is func for
func HashGenerator(str string) (string, error) {
	hashedString, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedString), nil
}

// HashComparator is func for
func HashComparator(hashedByte []byte, passwordByte []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedByte, passwordByte)
	if err != nil {
		return err
	}

	return nil
}
