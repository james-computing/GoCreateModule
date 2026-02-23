package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}
	var message string = fmt.Sprintf("Hello %v", name)
	return message, nil
}
