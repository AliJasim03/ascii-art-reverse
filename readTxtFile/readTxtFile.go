package asciiArt

import (
	"bufio"
	"fmt"
	"os"
)

func GetAsciiArt(s string) []string {
	s = "Ascii_files/" + s + ".txt"
	readFile, err := os.Open(s)
	if err != nil {
		fmt.Println("There is an error on the banner file") //print and error msg if there is any
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
