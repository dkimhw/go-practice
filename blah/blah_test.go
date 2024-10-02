package blah_test

import (
	"testing"
)

func TestBlah(t *testing.T) {
	want := "Blah"
	got := blah.getBlah()

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}
