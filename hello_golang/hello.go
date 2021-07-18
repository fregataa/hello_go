package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
    // fmt.Println("Hello, Golang!")
	// fmt.Println(quote.Go())
	message := greetings.Hello("Lee")
	fmt.Println(message)
}