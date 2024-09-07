package main

import (
	"errors"
	"fmt"
	"testing"
	"testing/quick"
)

var (
	cases = []struct {
		Arabic uint16
		Roman  string
	}{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 6, Roman: "VI"},
		{Arabic: 7, Roman: "VII"},
		{Arabic: 8, Roman: "VIII"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 14, Roman: "XIV"},
		{Arabic: 18, Roman: "XVIII"},
		{Arabic: 20, Roman: "XX"},
		{Arabic: 39, Roman: "XXXIX"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 47, Roman: "XLVII"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 100, Roman: "C"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 1006, Roman: "MVI"},
		{Arabic: 798, Roman: "DCCXCVIII"},
	}

	testInputs = []struct {
		Arabic uint16
		Roman  string
	}{
		{Arabic: 4000, Roman: ""},
		{Arabic: 0, Roman: ""},
	}
)

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got, _ := ConvertToRoman(test.Arabic)
			want := test.Roman

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}

	for _, testInput := range testInputs {

		t.Run("test input validation for convert to roman", func(t *testing.T) {

			_, err := ConvertToRoman(testInput.Arabic)
			want := errors.New(errorMsg).Error()

			if err.Error() != want {
				t.Errorf("got error %q want %q", err, want)
			}

		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

// Can't have negative intergers - use uint16
// What if we could take these rules that we know about our domain
// and somehow exercise them against our code?
// Property based tests help you do this by throwing random data at your code
// and verifying the rules you describe always hold true
func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		// Roman numerals can't be bigger than 3999
		if arabic > 3999 {
			return true
		}

		t.Log("testing", arabic)

		roman, _ := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)

		return fromRoman == arabic
	}

	// quick.Check a function that it will run against a number of random inputs
	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
