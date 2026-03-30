package main

import (
	"fmt"
	"strings"
)
func main() {
	var s, string string
	fmt.Println("The String Transformer")
	fmt.Scan(&s)
}

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
		w[i] = strings.ToLower(string(w[i][0])) + strings.ToLower(w[i][1:])
	}
	return strings.Join(w, " ")
}
func capFirstLetter(s string) string {
	w := strings.Fields(s)
	for i := range w {
		w[i] = strings.ToUpper(string(w[i][0])) + strings.ToLower(w[i][1:])
	}
	return strings.Join(w, " ")
}
func replaceWord(s string) string {
	w := strings.Fields(s)
	s = strings.ReplaceAll(s, "!", "")
	s = strings.ToLower(strings.ReplaceAll(s, "", "_"))
	return s
}
func reverseWord(s string) string{
	w := strings.Fields(s)
	for i, j := 0, len(w)-1; i < j; i, j = i + j, j - i {
		w[i], w[j] = w[j], w[i]
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
	text := (
		"sentinel is online"
        "ALERT LEVEL FIVE DETECTED" 
		 "director adaeze okonkwo" 
		"Operation Gopher Protocol" 

	)

	transformations := []func(string) string{
		ToUpper,
		ToLower,
		capFirstLetter,
		reverseWord,
	}
	result := applyAll(text, transformations)
	fmt.Println(result)
}
