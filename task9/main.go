package main

import (
	"fmt"
	"unicode"
)

func UnpackString(s string) (string, error) {
	runes := []rune(s)
	var prevRune rune
	var isEsc, isOk bool
	var out []rune
	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if !unicode.IsDigit(r) {
			if prevRune != 0 {
				out = append(out, prevRune)
			}
			if r == '\\' {
				isEsc = true
			} else {
				prevRune = r
			}
			isOk = true
		} else {
			if isEsc {
				isEsc = false
				prevRune = r
			} else if prevRune != 0 {
				for i := 0; i < int(r-'0'); i++ {
					out = append(out, prevRune)
				}
				prevRune = 0
			}
		}
	}
	if prevRune != 0 {
		out = append(out, prevRune)
	}
	if !isOk && len(s) > 0 {
		return "", fmt.Errorf("string contains digids only")
	}
	return string(out), nil
}

func main() {
	//s := "a4bc2d5e"
	s := []rune{'q', 'w', 'e', '\\', '4', '\\', '5'}
	fmt.Println(UnpackString(string(s)))
}
