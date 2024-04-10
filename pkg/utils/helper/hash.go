package helper

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err, "problem at hashing")
	}
	return string(HashedPassword)
}

func CompairPassword(hashedPassword string,plainPassword string) error {
	fmt.Println("==",hashedPassword, plainPassword)

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(plainPassword))

	if err != nil {
		return errors.New("password do not match")
	}

	return nil
}