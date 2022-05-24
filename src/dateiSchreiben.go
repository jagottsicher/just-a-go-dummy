package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("eineDatei.txt")
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}
	defer f.Close()

	unserInhalt := strings.NewReader("Hier kommt der Landgraf!\n")

	io.Copy(f, unserInhalt)
}
