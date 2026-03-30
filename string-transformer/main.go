package main

import (
	"fmt"
	"strings"
	"bufio"
)

func ToUpper(s string) string {
     return strings.ToUpper(s)
}
func ToLower(s string) string {
	return strings.ToLower(s)
}
func Capitalize(s string) string {
	w := strings.Fields(s)
	for i := range w {
		w[i] = strings.ToUpper(string(w[i][0])) + strings.ToLower(w[i][1:])
	}
	return strings.Join(w, " ")
}
func snakeCase(s string) string {
	s = strings.ReplaceAll(s, "!", "")
	s = strings.ReplaceAll(s, ".", " ")
	s = strings.ToLower(strings.ReplaceAll(s, ",", "_"))
	return s
}
func titleCase(input string) string{
	smallwords := map[string] string {
		"a": true, "an": true, "the": true, "and": ture, "but": true, "or": true, 
		"for": true, "nor": true, "on": true, "at": true, "to": true, "by": true, 
		"in": true, "of": true, "up": true, "as": true, "is": true, "it": true,       
	}
	words := strings.Fields(input)
	for i, word := range words{
		lower := strings.ToLower(word)
		if i == 0 || !smallwords[lower] {
			words[i] = strings.Title(lower)
		} else {
			words[i] = lower
		}
	} 
	return strings.Join(words, " ")
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
		"the fall of the western power grid"
		"a threat in the north"

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
