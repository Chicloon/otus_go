package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	c := "Hello, OTUS!"
	fmt.Println(PrintReverse(c))
}

func PrintReverse(s string) string {
	res := stringutil.Reverse(s)
	return res
}
