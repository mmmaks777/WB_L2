package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func Unpack(str string) (string, error) {
	var result []rune
	var prevRune rune
	escape := false

	for idx, r := range str {
		if escape {
			result = append(result, r)
			prevRune = r
			escape = false
			continue
		}

		if string(r) == `\` {
			escape = true
			continue
		}

		if unicode.IsDigit(r) {
			if idx == 0 || prevRune == 0 {
				return "", errors.New("incorrect string")
			}

			cnt, err := strconv.Atoi(string(r))
			if err != nil {
				return "", err
			}

			for i := 0; i < cnt-1; i++ {
				result = append(result, prevRune)
			}

			prevRune = 0
		} else {
			result = append(result, r)
			prevRune = r
		}
	}

	if escape {
		return "", errors.New("incorrect string")
	}

	return string(result), nil
}

func main() {
	fmt.Println(Unpack("a4bc2d5e"))
	fmt.Println(Unpack("abcd"))
	fmt.Println(Unpack(""))
	fmt.Println(Unpack(`qwe\\5`))
}
