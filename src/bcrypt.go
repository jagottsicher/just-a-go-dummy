package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	passwortKlartext := `pa55w0rd123öäü\/&%世界,`
	passwortKlartextAlsSOB := []byte(passwortKlartext)
	sliceOfByte, _ := bcrypt.GenerateFromPassword(passwortKlartextAlsSOB, bcrypt.MinCost)
	fmt.Println(sliceOfByte)
	fmt.Println(string(sliceOfByte))
}
