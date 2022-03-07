package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	c := "Hello, OTUS!"
	printReverse(c)
}

func printReverse(s string) {
	fmt.Println(stringutil.Reverse(s))
}
