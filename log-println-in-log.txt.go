package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("keineDatei.txt")
	if err != nil {
		//		fmt.Println("Ein Fehler trat auf:", err)
		log.Println("Ein Fehler trat auf:", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
	defer f2.Close()

	fmt.Println("Prüfen Sie die Datei log.txt.")
}
