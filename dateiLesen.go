package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("eineDatei.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	gelesenerInhalt, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(gelesenerInhalt))
}
