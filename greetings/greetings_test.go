package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`) // regex for strings that include the name
	message, err := Hello(name)
	// Test if the message have the name or we got a non nil err
	if !want.MatchString(message) || err != nil {
		t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}
