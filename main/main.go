package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1]
	
	chars := []rune(arg)
	if arg == ""{
		return
	}
	//read from the first argument
	readFile, err := os.Open("../standard.txt")
	if err != nil {
		fmt.Println("There is an error on the sample file") //print and error msg if there is any
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile) //creating a new scanenr on the file object
	
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	for i := 0; i < 8; i++ {
	for _,char := range chars{
/* 		if char == '\\'{
				// fmt.Println()
			
			continue
		} */
	targetLine := 10+(9*int(char-33))
		fmt.Print(lines[targetLine+i])
	}
	fmt.Println()

}

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
	readFile.Close() //close the readfile object ... system resources are freed up and that the file is properly released.

}
