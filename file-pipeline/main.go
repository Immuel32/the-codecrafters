// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: Gophers
// ───────────────────────────────────────────
// Input line types:
// Number of lines: 20
// Normal report lines
// Lines in ALL CAPS
// Lines in all lowercase
// Lines starting with TODO:
// Lines with extra leading/trailing spaces

// Transformation rules (in order):
// 1. Trim all leading and trailing whitespace
// 2. Replace TODO: with ✦ ACTION:
// 3. Convert ALL CAPS lines to Title Case
// 4. Convert all lowercase lines to uppercase
// 5. Reverse the words in any line that contains the word REVERSE

// Output format:
// Header: Yes, Exact Text: "Gopher's Sentinel Field Report - Processed"
// Line numbering format : "1."
// Summary block: yes
//
//	 	Fields :
//			✦ Lines read    :
//			✦ Lines written :
//			✦ Lines removed :
//			✦ Rules applied : [our 5 rules]
//
// Terminal summary fields:
//
//	✦ Lines read    :
//	✦ Lines written :
//	✦ Lines removed :
//	✦ Rules applied : [our 5 rules]
//
// ═══════════════════════════════════════════
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func TrimWhiteSpaces(line string) string {
	return strings.TrimSpace(line)
}
func ReplaceToDo(line string) string{
	return strings.ReplaceAll(line, "TODO:", "Action:")

}
func AllCapsToTitle(line string) string {
	if line == strings.ToUpper(line) && len(line) > 0 {
		words := strings.Fields(strings.ToLower(line))
		for i := range words {
			words[i] = strings.Title(words[i])

		}
        return strings.Join(words, " ")
	}
	return line
}
func Lower(line string) string{
	return strings.ToUpper(line)
}
func reverseWord(s string) string{
	 w := strings.Fields(s)
	for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
		w[i], w[j] = w[j], w[i]
	}
	return strings.Join(w, " ")
}
func main() {
	if len(os.Args)-1 != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>" )
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		fmt.Println("input file and output file cannot be the same")
	}
	return
}
in, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("✗ File not found: %s\n", inputFile)
		return
	}
	defer in.Close()

	scanner := bufio.NewScanner(in)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	lineRead := len(line)
	linesRemove = 0 

	if lineRead == 0{
		fmt.Println("Inputline is empty. Nothing to Process")
	}

	var Processed []string

	for _, line range line {
		line = TrimWhiteSpaces(lines)
		lines = ReplaceTODO(lines)
		lines = AllCapsToTitle(lines)
		continue
	}
	Processed = append(Processed, line)