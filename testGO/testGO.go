package main

import (
	"exemple/tutoGO/module"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("Attention: ")
	log.SetFlags(0)

	message2, err2 := module.Hello("semmy")
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(message2)

	// If an error was returned, print it to the console and
	// exit the program.

	message, err := module.Hello("")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(message + "ok")
	}
}
