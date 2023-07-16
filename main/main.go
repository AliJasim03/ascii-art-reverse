package main

import (
	"fmt"
	"os"
	"asciiArt"
)

func main() {

	if len(os.Args) != 2{
		fmt.Println("Please provide one argument to art")
		return
	}
	//read from the first argument
	arg := os.Args[1]
	if arg == "" {
		return
	}
	asciiArt.Print()
	words := asciiArt.SplitArg(arg)
	var lines []string = asciiArt.GetAsciiArt()

	for _, word := range words {
		if word == ""{
			continue
		}
		if word == "\\n" {
			fmt.Println()
			continue
		}
		chars := []rune(word)
		for i := 0; i < 8; i++ {
			for _, char := range chars {
				targetLine := 10 + (9 * int(char-33))
				fmt.Print(lines[targetLine+i])
			}
			fmt.Println()
		}
	}
}

// Not effieciat solution xd

//	for _,char := range chars{
// 	fmt.Println(char)

// 	targetLine := 10+(9*int(char-33))
// 	for i := 0; i < 8; i++ {
// 		fmt.Println(lines[targetLine+i])
// 	}

// }
// //loop until the end of the file
// for fileScanner.Scan() {
// 	lineNumber++
// 	if lineNumber == targetLine {
// 		for i := 0; i < 8; i++ {
// 			// fmt.Println(fileScanner.Text())
// 			fileScanner.Scan()
// 		}
// 		break
// 	}
// }
