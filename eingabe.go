package main

import "fmt"

func main() {
	var antwort string
	fmt.Print("Name: ")
	_, err := fmt.Scan(&antwort)
	if err != nil {
		fmt.Println("Fehler:", err)
	}
	fmt.Println("Name:", antwort)
}
