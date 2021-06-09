package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}


/*
Add a Hellos function whose parameter is a slice of names rather than a single name. Also, you change one of its return types from a string to a map so you can return names mapped to greeting messages.
Have the new Hellos function call the existing Hello function. This helps reduce duplication while also leaving both functions in place.
Create a messages map to associate each of the received names (as a key) with a generated message (as a value). In Go, you initialize a map with the following syntax: make(map[key-type]value-type). You have the Hellos function return this map to the caller. For more about maps, see Go maps in action on the Go blog.
Loop through the names your function received, checking that each has a non-empty value, then associate a message with each. In this for loop, range returns two values: the index of the current item in the loop and a copy of the item's value. You don't need the index, so you use the Go blank identifier (an underscore) to ignore it. For more, see The blank identifier in Effective Go.
*/

func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with
		// the name.

		messages[name] = message
	}

	return messages, nil

}

/*
Go executes init functions automatically at program startup, after global
variables have been initialized. For more about init functions, see Effective Go.
*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
Note that randomFormat starts with a lowercase letter, making it accessible
only to code in its own package (in other words, it's not exported).
*/
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Wel met!",
	}

	return formats[rand.Intn(len(formats))]
}
