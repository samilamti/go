package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHellosName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hellos([]string{name})
	if !want.MatchString(msg[name]) || err != nil {
		t.Fatalf(`Hello("%v") = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHellosEmpty(t *testing.T) {
	msg, err := Hellos([]string{""})
	if msg[""] != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
