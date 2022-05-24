package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("keineDatei.txt")
	if err != nil {
		//		fmt.Println("Ein Fehler trat auf:", err)
		//		log.Println("Ein Fehler trat auf:", err)
		log.Fatalln(err)
		//		panic(err)
	}
}

func foo() {
	fmt.Println("Weil durch log.Fatalln(err) os.Exit(1) aufgerufen wird, dird diese verzögerte Funktion nie aufgerufen")
}
