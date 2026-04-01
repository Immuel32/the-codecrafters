// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: Emmanuel Elaigwu
// Squad:  The Gophers Group

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func Capitalize(word string) string {
	words := strings.Fields(word)
	for i, ch := range words {
		runes := []rune(strings.ToLower(ch))
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func snakeCase(s string) string {
	s = strings.ReplaceAll(s, "!", "")
	s = strings.ReplaceAll(s, ".", " ")
	s = strings.ToLower(strings.ReplaceAll(s, ",", "_"))
	return s
}

func titleCase(input string) string {
	smallwords := map[string]bool{
		"a": true, "an": true, "the": true, "and": true, "but": true, "or": true,
		"for": true, "nor": true, "on": true, "at": true, "to": true, "by": true,
		"in": true, "of": true, "up": true, "as": true, "is": true, "it": true,
	}

	words := strings.Fields(input)
	for i, word := range words {
		lower := strings.ToLower(word)
		if i == 0 || !smallwords[lower] {
			if len(lower) > 0 {
				words[i] = strings.ToUpper(string(lower[0])) + lower[1:]
			}
		} else {
			words[i] = lower
		}
	}
	return strings.Join(words, " ")
}

func reverseWord(s string) string {
	w := strings.Fields(s)
	for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
		w[i], w[j] = w[j], w[i]
	}
	return strings.Join(w, " ")
}

func main() {
	fmt.Println("==== String Transformer ====\n")
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter your text: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Thank You,.Goodbye!")
			break
		}

		fmt.Println("\nChoose a transformation:")
		fmt.Println("1) upper      >> All UPPERCASE")
		fmt.Println("2) lower      >> All lowercase")
		fmt.Println("3) capitalize >> Capitalize Each Word")
		fmt.Println("4) reverse    >> Reverse Word Order")

		fmt.Print("Enter choice (1-4): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var result string

		switch choice {
		case "1":
			result = ToUpper(input)
		case "2":
			result = ToLower(input)
		case "3":
			result = Capitalize(input)
		case "4":
			result = reverseWord(input)
		default:
			fmt.Println("Invalid choice. Returning original text.")
			result = input
		}

		fmt.Println("\nTransformed text:")
		fmt.Println(result)
	}
}
