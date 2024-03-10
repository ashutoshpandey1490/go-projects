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

	// msg, err := concepts.Print_my_str_with_error("")
	// // nil is equivalent of null in Java here
	// if err != nil {
	// 	log.Fatal(err)
	// }

	create_name_map()

	fmt.Println(msg)
	fmt.Println(quote.Go())

	defer_concept()
}

func create_name_map() {
	names := []string{"ash", "soo", "xyz"}
	mapped_names := concepts.Map_names(names)
	fmt.Println(mapped_names)
}

func defer_concept() {
	for i := 0; i < 9; i++ {
		// defer statements will in stack and LIFO order will be followed
		defer fmt.Println(i)
		fmt.Println("Done")
	}
}
