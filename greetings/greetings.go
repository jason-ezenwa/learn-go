package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name:= range names {
		message, error:= Hello(name)

		if error != nil {
			return nil, error
		}

		messages[name] = message
	}

  darrin := messages["Darrin"]

	fmt.Println(darrin)

	return messages, nil
}

func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.

		randomInt := rand.Intn(len(formats))
    return formats[randomInt]
}
