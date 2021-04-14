package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
func init() {
	rand.Seed(time.Now().UnixNano())
}
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! well met!",
	}
	return formats[rand.Intn(len(formats))]
}
