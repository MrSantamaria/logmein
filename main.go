package main

import (
	"log"

	"github.com/MrSantamaria/logmein/ui"
)

func main() {

	_, _, err := ui.LoginWindow()
	if err != nil {
		log.Fatal(err)
	}
}
