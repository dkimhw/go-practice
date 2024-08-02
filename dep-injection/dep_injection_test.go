package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// The Buffer type from the bytes package implements the Writer interface
	buffer := bytes.Buffer{}
	Greet(&buffer, "Patri")

	got := buffer.String()
	want := "Hello, Patri"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
