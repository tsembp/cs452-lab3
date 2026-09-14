package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

var randomFormats = []string{
	"Hi, %v. Welcome!",
	"Great to see you, %v!",
	"Hello, %v!",
}

// Hello returns a random greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	format := randomFormats[rand.Intn(len(randomFormats))]
	return fmt.Sprintf(format, name), nil
}

// Hellos returns greetings for multiple people.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}
