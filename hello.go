package main

import (
	"fmt"
	"log"

	"example/hello/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Alex", "George", "Maria"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, message := range messages {
		fmt.Println(message)
	}
}
