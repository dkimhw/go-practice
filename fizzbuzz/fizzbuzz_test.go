package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("should return correct fizzbuzz values", func(t *testing.T) {
		n := 3
		want := []string{"1", "2", "Fizz"}
		got := FizzBuzz(n)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test more variations", func(t *testing.T) {
		n := 5
		want := []string{"1", "2", "Fizz", "4", "Buzz"}
		got := FizzBuzz(n)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test even longer variations", func(t *testing.T) {
		n := 15
		want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
		got := FizzBuzz(n)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
