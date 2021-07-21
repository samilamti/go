package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Hellos(names []string) (map[string]string, error) {
	log.SetPrefix("Hellos: ")
	log.SetFlags(log.Ldate + log.Ltime)

	messages := make(map[string]string)

	for _, name := range names {
		if name == "" {
			return nil, errors.New("empty name")
		}

		message := fmt.Sprintf(randomFormat(), name)
		messages[name] = message
	}

	return messages, nil
}

func init() {
	log.SetPrefix("greetings: ")
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
