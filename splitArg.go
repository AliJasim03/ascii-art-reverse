package asciiArt

func SplitArg(arg string) []string {
	chars := []rune(arg)
	var counter int = 0
	var words []string
	words = append(words, "")
	for index, char := range chars {
		if index != len(chars) {
			if (chars[index] == '\\') && (chars[index+1] == 'n') {
				words = append(words, "")
				counter++
				// words[counter] = ""
			}
		}
		words[counter] = words[counter] + string(char)
	}
	return words
}
