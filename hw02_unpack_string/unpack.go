package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {

	var stringBuilder strings.Builder
	var previousChar rune

	for _, char := range str {

		if !unicode.IsDigit(char) {

			if previousChar == 0 {
				previousChar = char
			} else {
				stringBuilder.WriteRune(previousChar)
				previousChar = char
			}

		} else if previousChar != 0 {
			repeatedChar := strings.Repeat(string(previousChar), int(char-'0'))
			stringBuilder.WriteString(repeatedChar)

			previousChar = 0
		} else {
			return "", ErrInvalidString
		}

	}

	if previousChar != 0 {
		stringBuilder.WriteRune(previousChar)
	}

	return stringBuilder.String(), nil
}

func main() {
	fmt.Println(Unpack("a4bc2d5e"))
	fmt.Println(Unpack("abcd"))
	fmt.Println(Unpack("3abc"))
	fmt.Println(Unpack("45"))
	fmt.Println(Unpack("aaa10b"))
	fmt.Println(Unpack("aaa0b"))
	fmt.Println(Unpack(""))
	fmt.Println(Unpack("d\n5abc"))

}
