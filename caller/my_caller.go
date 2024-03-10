package main

import (
	"fmt"

	"learning.com/concepts"

	"rsc.io/quote"

	"log"
)

func main() {
	// Similar way of what we do in Java with Logger
	log.SetPrefix("concepts: ")
	log.SetFlags(0)

	fmt.Println("Hello from caller !!")

	msg := concepts.Print_my_str("Ash")
	fmt.Println(msg)

	msg, err := concepts.Print_my_str_with_error("")
	// nil is equivalent of null in Java here
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
	fmt.Println(quote.Go())
}
