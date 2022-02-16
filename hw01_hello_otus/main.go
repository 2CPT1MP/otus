package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {

	var originalPhrase string = "Hello, OTUS!"

	var reversedPhrase string = stringutil.Reverse(originalPhrase)
	fmt.Println(reversedPhrase)
}
