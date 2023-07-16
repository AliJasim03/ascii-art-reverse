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
func RunProgram(text string) (err error, output []string) {
	cmd := exec.Command("go", "run", ".", text)

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
	err, txt := RunProgram("hello")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"hello\"")
	}
}

func TestAudit2(t *testing.T) {
	ans := []string{
		` _    _   ______   _        _         ____   `,
		`| |  | | |  ____| | |      | |       / __ \  `,
		`| |__| | | |__    | |      | |      | |  | | `,
		`|  __  | |  __|   | |      | |      | |  | | `,
		`| |  | | | |____  | |____  | |____  | |__| | `,
		`|_|  |_| |______| |______| |______|  \____/  `,
		`                                             `,
		`                                             `,
	}
	err, txt := RunProgram("HELLO")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"HELLO\"")
	}
}

func TestAudit3(t *testing.T) {
	ans := []string{
		` _    _          _        _                 _    _           __  __           _   _  `,
		`| |  | |        | |      | |               | |  | |         |  \/  |         | \ | | `,
		`| |__| |   ___  | |      | |   ___         | |__| |  _   _  | \  / |   __ _  |  \| | `,
		"|  __  |  / _ \\ | |      | |  / _ \\        |  __  | | | | | | |\\/| |  / _` | | . ` | ",
		`| |  | | |  __/ | |____  | | | (_) |       | |  | | | |_| | | |  | | | (_| | | |\  | `,
		`|_|  |_|  \___| |______| |_|  \___/        |_|  |_|  \__,_| |_|  |_|  \__,_| |_| \_| `,
		`                                                                                     `,
		`                                                                                     `,
	}
	err, txt := RunProgram("HeLlo HuMaN")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"HeLlo HuMaN\"")
	}
}

func TestAudit4(t *testing.T) {
	ans := []string{
		`     _    _          _   _                         _______   _                           `,
		` _  | |  | |        | | | |                ____   |__   __| | |                          `,
		`/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  `,
		`| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ `,
		`| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ `,
		`|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| `,
		`                                                                                         `,
		`                                                                                         `,
	}
	err, txt := RunProgram("1Hello 2There")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"1Hello 2There\"")
	}
}

func TestAudit5(t *testing.T) {
	ans := []string{
		` _    _          _   _          `,
		`| |  | |        | | | |         `,
		`| |__| |   ___  | | | |   ___   `,
		`|  __  |  / _ \ | | | |  / _ \  `,
		`| |  | | |  __/ | | | | | (_) | `,
		`|_|  |_|  \___| |_| |_|  \___/  `,
		`                                `,
		`                                `,
		` _______   _                           `,
		`|__   __| | |                          `,
		`   | |    | |__     ___   _ __    ___  `,
		`   | |    |  _ \   / _ \ | '__|  / _ \ `,
		`   | |    | | | | |  __/ | |    |  __/ `,
		`   |_|    |_| |_|  \___| |_|     \___| `,
		`                                       `,
		`                                       `,
	}
	err, txt := RunProgram("Hello\\nThere")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"Hello\nThere\"")
	}
}

func TestAudit6(t *testing.T) {
	ans := []string{
		` _    _          _   _          `,
		`| |  | |        | | | |         `,
		`| |__| |   ___  | | | |   ___   `,
		`|  __  |  / _ \ | | | |  / _ \  `,
		`| |  | | |  __/ | | | | | (_) | `,
		`|_|  |_|  \___| |_| |_|  \___/  `,
		`                                `,
		`                                `,
		``,
		` _______   _                           `,
		`|__   __| | |                          `,
		`   | |    | |__     ___   _ __    ___  `,
		`   | |    |  _ \   / _ \ | '__|  / _ \ `,
		`   | |    | | | | |  __/ | |    |  __/ `,
		`   |_|    |_| |_|  \___| |_|     \___| `,
		`                                       `,
		`                                       `,
	}
	err, txt := RunProgram("Hello\\n\\nThere")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"Hello\n\nThere\"")
	}
}

func TestAudit7(t *testing.T) {
	ans := []string{
		`   __  _    _          _   _                                _______   _                                    _  _    __    `,
		`  / / | |  | |        | | | |                 ___          |__   __| | |                                 _| || |_  \ \   `,
		` | |  | |__| |   ___  | | | |   ___          ( _ )            | |    | |__     ___   _ __    ___        |_  __  _|  | |  `,
		`/ /   |  __  |  / _ \ | | | |  / _ \         / _ \/\          | |    |  _ \   / _ \ | '__|  / _ \        _| || |_    \ \ `,
		`\ \   | |  | | |  __/ | | | | | (_) |       | (_>  <          | |    | | | | |  __/ | |    |  __/       |_  __  _|   / / `,
		` | |  |_|  |_|  \___| |_| |_|  \___/         \___/\/          |_|    |_| |_|  \___| |_|     \___|         |_||_|    | |  `,
		`  \_\                                                                                                              /_/   `,
		`                                                                                                                         `,
	}
	err, txt := RunProgram("{Hello & There #}")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"{Hello & There #}\"")
	}
}

func TestAudit8(t *testing.T) {
	ans := []string{
		` _              _   _                 _______   _                                            _                           _  `,
		`| |            | | | |               |__   __| | |                                 _        | |                  ____   | | `,
		`| |__     ___  | | | |   ___            | |    | |__     ___   _ __    ___        / |       | |_    ___         |___ \  | | `,
		`|  _ \   / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \       | |       | __|  / _ \          __) | | | `,
		`| | | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/       | |       \ |_  | (_) |        / __/  |_| `,
		`|_| |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|       |_|        \__|  \___/        |_____| (_) `,
		`                                                                                                                            `,
		`                                                                                                                            `,
	}
	err, txt := RunProgram("hello There 1 to 2!")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"hello There 1 to 2!\"")
	}
}

func TestAudit9(t *testing.T) {
	ans := []string{
		` __  __           _____            _____                              _        _    _____   _               _   _  `,
		`|  \/  |         |  __ \   _____  |_   _|            /\       ___    | |      (_)  / ____| | |             | \ | | `,
		`| \  / |   __ _  | |  | | |___ /    | |    _ __     /  \     ( _ )   | |       _  | (___   | |__     ___   |  \| | `,
		"| |\\/| |  / _` | | |  | |   |_ \\    | |   | '__|   / /\\ \\    / _ \\/\\ | |      | |  \\___ \\  | '_ \\   / _ \\  | . ` | ",
		`| |  | | | (_| | | |__| |  ___) |  _| |_  | |     / ____ \  | (_>  < | |____  | |  ____) | | |_) | | (_) | | |\  | `,
		`|_|  |_|  \__,_| |_____/  |____/  |_____| |_|    /_/    \_\  \___/\/ |______| |_| |_____/  |_.__/   \___/  |_| \_| `,
		`                                                                                                                   `,
		`                                                                                                                   `,
	}
	err, txt := RunProgram("MaD3IrA&LiSboN")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"MaD3IrA&LiSboN\"")
	}
}

func TestAudit11(t *testing.T) {
	ans := []string{
		`   __  _  __     /\/| `,
		`  / / | | \ \   |/\/  `,
		` | |  | |  | |        `,
		`/ /   | |   \ \       `,
		`\ \   | |   / /       `,
		` | |  | |  | |        `,
		`  \_\ | | /_/         `,
		`      |_|             `,
	}
	err, txt := RunProgram("{|}~")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"{|}~\"")
	}
}

func TestAudit12(t *testing.T) {
	ans := []string{
		` ___  __       ___   /\                  _          `,
		`|  _| \ \     |_  | |/\|                ( )         `,
		`| |    \ \      | |                     |/    __ _  `,
		"| |     \\ \\     | |                          / _` | ",
		"| |      \\ \\    | |                         | (_| | ",
		"| |_      \\_\\  _| |                          \\__,_| ",
		"|___|         |___|       ______                    ",
		`                         |______|                   `,
	}
	err, txt := RunProgram("[\\]^_ 'a")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"[\\]^_ 'a\"")
	}
}

func TestAudit13(t *testing.T) {
	ans := []string{
		` _____     _____   ____   `,
		`|  __ \   / ____| |  _ \  `,
		`| |__) | | |  __  | |_) | `,
		`|  _  /  | | |_ | |  _ <  `,
		`| | \ \  | |__| | | |_) | `,
		`|_|  \_\  \_____| |____/  `,
		`                          `,
		`                          `,
	}
	err, txt := RunProgram("RGB")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \"RGB\"")
	}
}

func TestAudit14(t *testing.T) {
	ans := []string{
		`           __          __     ___             `,
		` _   _    / /  ______  \ \   |__ \     ____   `,
		`(_) (_)  / /  |______|  \ \     ) |   / __ \  `,
		"        < <    ______    > >   / /   / / _` | ",
		` _   _   \ \  |______|  / /   |_|   | | (_| | `,
		`(_) ( )   \_\          /_/    (_)    \ \__,_| `,
		`    |/                                \____/  `,
		`                                              `,
	}
	err, txt := RunProgram(":;<=>?@")
	if err != nil {
		t.Errorf("Couldn't run the program: %s\n", err)
	}
	areEqual := AreStringSlicesEqual(ans, txt)
	if areEqual != nil {
		t.Log(areEqual)
		t.Error("Invalid output for \":;<=>?@\"")
	}
}
