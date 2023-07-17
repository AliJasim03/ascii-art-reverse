package main

import (
	"asciiArt/readTxtFile"
	"fmt"
	"os"
	"strings"
)

func main() {
	// check if the argument is not equals 2 or 3
	argsLength := len(os.Args)

	if argsLength != 2 && argsLength != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		return
	}
	//read from the first argument
	arg := os.Args[1]
	if arg == "" {
		return
	}

	//split the argument into an array of string
	words := strings.Split(arg, "\\n")
	banner := "standard"

	if argsLength == 3 {
		banner = os.Args[2]
	}
	//read from the lines of the file
	var lines []string = asciiArt.GetAsciiArt(banner)

	if lines == nil {
		fmt.Println("Please choose a valid banner option from the list")
		fmt.Println("shadow, standard or thinkertoy")
		return
	}
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
