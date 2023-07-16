package asciiArt

import (
	"bufio"
	"fmt"
	"os"
)

func GetAsciiArt() []string {
	readFile, err := os.Open("readTxtFile/standard.txt")
	if err != nil {
		fmt.Println("There is an error on the sample file") //print and error msg if there is any
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile) //creating a new scanenr on the file object

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close() //close the readfile object ... system resources are freed up and that the file is properly released.

	return lines
}
