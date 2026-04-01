// CodeCrafters — Operation Gopher Protocol
// Module: Go-Reloaded (prt: fixArticle)
// Author: Emmanuel Elaigwu
// Squad:  The Gophers Group

package main

import (
	"fmt"
	"strings"
)

func fixArticle(word string) string {
	words := strings.Fields(word)
	vowels := "aeiouhAEIOUH"
	for i := 0; i < len(words)-1; i++ {
		if (words[i] == "a" || words[i] == "A") && strings.ContainsRune(vowels, rune(words[i+1][0])) {
			if words[i] == "A" {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(fixArticle("There is no greater agony than bearing a untold story inside you."))
}
