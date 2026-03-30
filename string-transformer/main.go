package main

import (
	"fmt"
	"strings"
)

func ToUpper(s string) string {
	w := strings.Fields(s)
	for i := range w {
		w[i] = strings.ToUpper(string(w[i][0])) + strings.ToUpper(w[i][1:])
	}
	return strings.Join(w, " ")
}
func ToLower(s string) string {
	w := strings.Fields(s)
	for i := range w {
		w[i] = strings.ToLower(string(w[i][0])) + strings.ToLoWer(w[i][1:])
	}
	return strings.Join(w, " ")
}

func applyAll(s string, funcs []func(string) string) string {
	for _, f := range funcs {
		s = f(s)
	}
	return s
}
func main() {
	text := "sentinel is online"
	text := "ALERT LEVEL FIVE DETECTED"
	transformations := []func(string) string{
		ToUpper,
		ToLower,
		reverseWord,
	}
	result := applyAll(text, transformations)
	fmt.Println(result)
}
