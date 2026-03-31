package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----------------------------------------------\n---- CodeCrafters - Operation Gopher Protocol\n---- Module: File Pipeline\n---- Author: Emmanuel Elaigwu\n---- Squad: The Gophers\n-------------------------------------------\n")
}

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
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println ("====String Transformer====")
	fmt.Print("Enter your text")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Println("Choose a transformation: ")
	fmt.Println("1") upper ()
	fmt.Println("2") lower ()
	fmt.Println("3") capitalize()



}
func applyAll(s string, funcs []func(string) string) string {
	for _, f := range funcs {
		s = f(s)
	}
	return s
}

	result := applyAll(text, transformations)
	fmt.Println(result)
}
