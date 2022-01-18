package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var text string = `
It is my 'Daktilo' function. You can write slowly your words.
It is an example for this function. 
Hello World! You are amazing!
	`
	// Usage 1: Daktilo(text(string), milliseconds(int))
	// Usage 2: Daktilo(text(string))
	Daktilo(text, 50)
}

func Daktilo(txt string, args ...int) {
	var durum bool = true
	milsec := 75
	if len(args) == 1 {
		milsec = args[0]
	} else {
		fmt.Println("[Error] Please enter a maximum of 2 arguments. \nExample 1: Daktilo('Hello World', 100)\nExample 2: Daktilo('Hello World')")
		durum = false
	}
	if durum {
		timing := time.Duration(milsec)    // millisecond
		txtchars := strings.Split(txt, "") // txtchars = array
		for _, character := range txtchars {
			fmt.Print(character)
			time.Sleep(time.Millisecond * timing)
		}
	}
}
