package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println(quote.Go())

	message, err := greetings.Hello("Gladys")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	message2, err2 := greetings.Hello("")
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Println(message2)

}
