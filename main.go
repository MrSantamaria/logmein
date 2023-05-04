package main

import (
	"log"

	_ "github.com/MrSantamaria/logmein/ui"
)

func main() {

	err := LoginWindow()
	if err != nil {
		log.Fatal(err)
	}
}
