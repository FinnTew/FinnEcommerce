package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type HashPwdUtil struct{}

func NewHashPwdUtil() *HashPwdUtil {
	return &HashPwdUtil{}
}

func (h *HashPwdUtil) GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("error generating hash: %v", err)
	}

	return string(hash), nil
}

func (h *HashPwdUtil) ValidatePwd(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return fmt.Errorf("invalid password: %v", err)
	}
	return nil
}
