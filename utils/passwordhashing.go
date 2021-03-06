package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) string {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		fmt.Println("failed to hash the password")

	}
	return string(hashpassword)
}

func Verify_password(givenpassword string, storedpassword string) bool {
	flag := true
	err := bcrypt.CompareHashAndPassword([]byte(storedpassword), []byte(givenpassword))
	if err != nil {
		flag = false
	}
	return flag
}
