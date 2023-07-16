package asciiArt
// import "fmt"
func SplitArg(arg string) []string {
	chars := []rune(arg)
	var counter int = 0
	var words []string
	words = append(words, "")
	for index, char := range chars {
		if index != len(chars) {
			if (chars[index] == '\\') && (chars[index+1] == 'n') {
				words = append(words, "\\n")
				words = append(words, "")
				counter+=2
				continue
			}
			if index != 0 {
			if (chars[index-1] == '\\') && (chars[index] == 'n'){
				continue
			}
		}
		}
		words[counter] = words[counter] + string(char)
	}
	return words
}
