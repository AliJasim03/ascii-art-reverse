package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"testing"
)

func BytesToStringSlice(byteSlice []byte) []string {
	trimmedSlice := bytes.TrimSuffix(byteSlice, []byte("\n"))
	lines := bytes.Split(trimmedSlice, []byte("\n"))
	stringSlice := make([]string, len(lines))

	for i, line := range lines {
		stringSlice[i] = string(line)
	}

	return stringSlice
}

func AreStringSlicesEqual(slice1, slice2 []string) error {
	if len(slice1) != len(slice2) {

		return errors.New("Length does not match")
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			msg := fmt.Sprintf("Text does not match at line %d\n", i)
			return errors.New(msg)
		}
	}
	return nil
}

// Args:
//
//	text (string): The input text to pass to the program.
//
// Returns:
//
//	err: An error encountered during program execution, if any.
//	[]string: The output of the program as a string slice, where each element is 1 line.
func RunProgram(text string, banner string) (err error, output []string) {
	cmd := exec.Command("go", "run", ".", text,banner)

	o, e := cmd.CombinedOutput()
	if e != nil {
		return e, nil
	}
	return nil, BytesToStringSlice(o)
}

func TestAudit1(t *testing.T) {
	ans := []string{
		` _              _   _          `,
		`| |            | | | |         `,
		`| |__     ___  | | | |   ___   `,
		`|  _ \   / _ \ | | | |  / _ \  `,
		`| | | | |  __/ | | | | | (_) | `,
		`|_| |_|  \___| |_| |_|  \___/  `,
		`                               `,
		`                               `,
	}
	err, txt := RunProgram("hello", "standard")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"hello\" standard")
	}
}

func TestAudit2(t *testing.T) {
	ans := []string{
		`                                                                                        `,
		`_|                _| _|                                                     _|       _| `,
		`_|_|_|     _|_|   _| _|   _|_|         _|      _|      _|   _|_|   _|  _|_| _|   _|_|_| `,
		`_|    _| _|_|_|_| _| _| _|    _|       _|      _|      _| _|    _| _|_|     _| _|    _| `,
		`_|    _| _|       _| _| _|    _|         _|  _|  _|  _|   _|    _| _|       _| _|    _| `,
		`_|    _|   _|_|_| _| _|   _|_|             _|      _|       _|_|   _|       _|   _|_|_| `,
		`                                                                                        `,
		`                                                                                        `,		
	}
	err, txt := RunProgram("hello world", "shadow")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"hello world\" shadow")
	}
}

func TestAudit3(t *testing.T) {
	ans := []string{
		`                                                                       `,
		`                       --                       o                      `,
		`     o                o  o                      |                      `,
		`o-o     o-o o-o         /        o-O-o o-o o-o -o-       o  o o-o o  o `,
		`|  | | |    |-'        /         | | | |-' |-'  |        |  | | | |  | `,
		`o  o |  o-o o-o       o--o       o o o o-o o-o  o        o--O o-o o--o `,
		`                                                            |          `,
		`                                                         o--o          `,
	}
	err, txt := RunProgram("nice 2 meet you", "thinkertoy")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"nice 2 meet you\" thinkertoy")
	}
}

// func TestAudit4(t *testing.T) {
// 	ans := []string{
// 		`                                                                `,
// 		`                                ___                             `,
// 		` _   _    ___    _   _         ( _ )          _ __ ___     ___  `,
// 		`| | | |  / _ \  | | | |        / _ \/\       | '_`` _ \   / _ \ `,
// 		`| |_| | | (_) | | |_| |       | (_>  <       | | | | | | |  __/ `,
// 		` \__, |  \___/   \__,_|        \___/\/       |_| |_| |_|  \___| `,
// 		` __/ /                                                          `,
//    	    `|___/                                                           `,
// 	}
	
// 	err, txt := RunProgram("you & me", "standard")
// 	if err != nil {
// 		t.Errorf("Couldn't run the program: %s\n", err)
// 	}
// 	areEqual := AreStringSlicesEqual(ans, txt)
// 	if areEqual != nil {
// 		t.Log(areEqual)
// 		t.Error("Invalid output for \"you & me\" standard")
// 	}
// }

func TestAudit5(t *testing.T) {
	ans := []string{
		`                       `,
		`  _|   _|_|   _|_|_|   `,
		`_|_| _|    _|       _| `,
		`  _|     _|     _|_|   `,
		`  _|   _|           _| `,
		`  _| _|_|_|_| _|_|_|   `,
		`                       `,
		`                       `,
	}
	err, txt := RunProgram("123", "shadow")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"123\" shadow")
	}
}

func TestAudit6(t *testing.T) {
	ans := []string{
		`         o o    `,
		`    o  / | | \  `,
		`   /  o       o `,
		`  o   |       | `,
		` /    o       o `,
		`o      \     /  `,
		`                `,
		`                `,
	}
	err, txt := RunProgram("/(\")", "thinkertoy")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"/(\")\" thinkertoy")
	}
}

func TestAudit7(t *testing.T) {
	ans := []string{
			`                                                                                                                                                                                                                                                              `,
			`  _|_|   _|_|_|     _|_|_| _|_|_|   _|_|_|_| _|_|_|_|   _|_|_| _|    _| _|_|_|       _| _|    _| _|       _|      _| _|      _|   _|_|   _|_|_|     _|_|     _|_|_|     _|_|_| _|_|_|_|_| _|    _| _|      _| _|          _| _|      _| _|      _| _|_|_|_|_| `,
			`_|    _| _|    _| _|       _|    _| _|       _|       _|       _|    _|   _|         _| _|  _|   _|       _|_|  _|_| _|_|    _| _|    _| _|    _| _|    _|   _|    _| _|           _|     _|    _| _|      _| _|          _|   _|  _|     _|  _|         _|   `,
			`_|_|_|_| _|_|_|   _|       _|    _| _|_|_|   _|_|_|   _|  _|_| _|_|_|_|   _|         _| _|_|     _|       _|  _|  _| _|  _|  _| _|    _| _|_|_|   _|  _|_|   _|_|_|     _|_|       _|     _|    _| _|      _| _|    _|    _|     _|         _|         _|     `,
			`_|    _| _|    _| _|       _|    _| _|       _|       _|    _| _|    _|   _|   _|    _| _|  _|   _|       _|      _| _|    _|_| _|    _| _|       _|    _|   _|    _|       _|     _|     _|    _|   _|  _|     _|  _|  _|     _|  _|       _|       _|       `,
			`_|    _| _|_|_|     _|_|_| _|_|_|   _|_|_|_| _|         _|_|_| _|    _| _|_|_|   _|_|   _|    _| _|_|_|_| _|      _| _|      _|   _|_|   _|         _|_|  _| _|    _| _|_|_|       _|       _|_|       _|         _|  _|     _|      _|     _|     _|_|_|_|_| `,
			`                                                                                                                                                                                                                                                              `,
			`                                                                                                                                                                                                                                                              `,
	}
	err, txt := RunProgram("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "shadow")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"ABCDEFGHIJKLMNOPQRSTUVWXYZ0\" shadow")
	}
}

func TestAudit8(t *testing.T) {
	ans := []string{
			`o o         | |                                                  `,
			`| |  | |   -O-O-      O          o  / \  o | o                 o `,
			`    -O-O- o | |   o  /    o     /  o   o  \|/   |             /  `,
			`     | |   -O-O-    /    /|    o   |   | --O-- -o-           o   `,
			`    -O-O-   | | o  /  o o-O-  /    o   o  /|\   |    o-o    /    `,
			`     | |   -O-O-  O       |  o      \ /  o | o     o     O o     `,
			`            | |                                    |             `,
			`                                                                 `,
	
	}
	err, txt := RunProgram("\"#$%&/()*+,-./", "thinkertoy")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"\"#$%&/()*+,-./\"")
	}
}

func TestAudit9(t *testing.T) {
	ans := []string{
			`          o                                              `,
			`o-O-o  o  |           o       o         o                `,
			`  |    |              |       |         | /  o           `,
			`  |   -o-   o-o       o   o   o o-o o-o OO     o-o  o--o `,
			`  |    |     \         \ / \ /  | | |   | \  | |  | |  | `,
			`o-O-o  o    o-o         o   o   o-o o   o  o | o  o o--O `,
			`                                                       | `,
			`                                                    o--o `,

	}
	err, txt := RunProgram("It's Working", "thinkertoy")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"It's Working\"", "thinkertoy")
	}
}