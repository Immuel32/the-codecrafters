package main

import (
	"fmt"
)

func applyAll(s string, funcs []func(string) string) string {
	for _, f := range funcs {
		s = f(s)
	}
	return s
}
func main() {
	text := "hello word"
	transformations := []func(string) string{
		capWords,
		reverseWord,
	}
	result := applyAll(text, transformations)
	fmt.Println(result)
}
