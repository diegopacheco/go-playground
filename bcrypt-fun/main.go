package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)
	// Hash password with bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

func doPasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err == nil
}

func main() {
	var hashedPassword, err = hashPassword("password1")
	if err != nil {
		println(fmt.Println("Error hashing password"))
		return
	}
	fmt.Println("Password Hash:", hashedPassword)
	fmt.Println("Password Match:", doPasswordsMatch(hashedPassword, "password1"))
}
