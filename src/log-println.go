package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("keineDatei.txt")
	if err != nil {
		// fmt.Println("Ein Fehler trat auf:", err)
		log.Println("Ein Fehler trat auf:", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
}
