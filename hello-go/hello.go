package main

import (
	"fmt"

	"log"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("Main: ")
	fmt.Println(quote.Go())

	messages, err := greetings.Hellos([]string{"Banana-man", ""})
	log.Printf("Got some stuff from greetings")

	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
