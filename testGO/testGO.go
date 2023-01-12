package main

import (
	"exemple/tutoGO/module"
	"fmt"
	"log"
	"time"
)

func main() {

	// test   variable
	var conferenceName = "La conference de Semmy Guiose il y a"
	const conferenceTicket = 50
	var remainingTickets = 50

	fmt.Println(conferenceName, "ticket disponible", remainingTickets, "sur", conferenceTicket)
	fmt.Printf("%v ticket disponible %v sur %v \n", conferenceName, remainingTickets, conferenceTicket)
	//saisie clavier
	//	module.Saisie()

	fmt.Println(time.Now())

	// log and erreur
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
