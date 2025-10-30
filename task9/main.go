package main

import (
	"fmt"
	"unicode"
)

func UnpackString(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	runes := []rune(s)
	var out []rune
	var prev rune
	var esc, nonDigidFound bool

	for _, r := range runes {
		switch {
		case esc:
			out = append(out, r)
			esc = false
		case r == '\\':
			if prev != 0 {
				out = append(out, prev)
				prev = 0
			}
			esc = true
			nonDigidFound = true
		case unicode.IsDigit(r):
			if prev == 0 {
				return "", fmt.Errorf("digid without prev char")
			}
			count := int(r - '0')
			for i := 0; i < count; i++ {
				out = append(out, prev)
			}
			prev = 0
		default:
			if prev != 0 {
				out = append(out, prev)
			}
			prev = r
			nonDigidFound = true
		}
	}
	if esc {
		return "", fmt.Errorf("escape char at the end")
	}

	if prev != 0 {
		out = append(out, prev)
	}

	if !nonDigidFound {
		return "", fmt.Errorf("string contains digids only")
	}

	return string(out), nil
}

func main() {
	//s := "a4bc2d5e"
	s := []rune{'q', 'w', 'e', '\\', '4', '\\', '5'}
	fmt.Println(UnpackString(string(s)))
}
