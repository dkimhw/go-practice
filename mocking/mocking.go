package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
	for i := countdownStart; i >= 1; i-- {
		fmt.Fprint(out, strconv.Itoa(i)+"\n")
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
