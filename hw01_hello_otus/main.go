package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	originalPhrase := "Hello, OTUS!"
	reversedPhrase := stringutil.Reverse(originalPhrase)

	fmt.Println(reversedPhrase)
}
