package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	// Remove time stamp and source information
	log.SetFlags(0)

	message, err := greetings.Hello("John")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
