package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}

/*
Commands:

go build // generate hello binary
./hello // execute binary

You can discover the install path by running the go list command, as in the following example:

go list -f '{{.Target}}'

*/