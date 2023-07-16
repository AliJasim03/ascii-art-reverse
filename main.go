package main

import (
	"asciiArt/readTxtFile"
	"fmt"
	"os"
	"strings"
)

func main() {
	// check if the argument is not equals 2

	if len(os.Args) != 2 {
		fmt.Println("Please provide one phrase to art")
		return
	}
	//read from the first argument
	arg := os.Args[1]
	if arg == "" {
		return
	}

	//split the argument into an array of string
	words := strings.Split(arg, "\\n")

	//read from the lines of the file
	var lines []string = asciiArt.GetAsciiArt()

	//loop throught out the array of strings

	for _, word := range words {

		//cheack if it is empty, if it is print new line
		if word == "" {
			fmt.Println()
			continue
		}

		//convert the stirng into array of rune
		chars := []rune(word)

		//print the 8 lines
		for i := 0; i < 8; i++ {
			for _, char := range chars {
				targetLine := 10 + (9 * int(char-33))
				fmt.Print(lines[targetLine+i])
			}
			fmt.Println()
		}
	}
}
