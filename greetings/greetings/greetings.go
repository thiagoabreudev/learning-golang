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
