package main

import (
	"log"

	_ "github.com/MrSantamaria/logmein/ui"
)

func main() {

	err := login_window()
	if err != nil {
		log.Fatal(err)
	}
}
