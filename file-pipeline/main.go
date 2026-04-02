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
func ReplaceToDo(line string) string {
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
func Lower(line string) string {
	return strings.ToUpper(line)
}
func reverseWord(s string) string {
	w := strings.Fields(s)
	for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
		w[i], w[j] = w[j], w[i]
	}
	return strings.Join(w, " ")
}
func main() {
	if len(os.Args)-1 != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		fmt.Println("input file and output file cannot be the same")
		return
	}

	in, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("File not found: %s\n", inputFile)
		return
	}
	defer in.Close()

	scanner := bufio.NewScanner(in)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	linesRead := len(lines)
	linesRemoved := 0

	if linesRead == 0 {
		fmt.Println("Input file is empty. Nothing to process.")
		return
	}

	transformations := []func(string) string{
		TrimWhiteSpaces,
		ReplaceToDo,
		AllCapsToTitle,
		Lower,
		reverseWord,
	}

	var processed []string
	for _, line := range lines {
		for _, fn := range transformations {
			line = fn(line)
		}
		processed = append(processed, line)
	}
	for i := range processed {
		processed[i] = fmt.Sprintf("%03d. %s", i+1, processed[i])
	}
	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Cannot write to output: %s\n", outputFile)
		return
	}
	defer out.Close()

	fmt.Fprintln(out, "Gopher's Sentinel Field Report - Processed")

	for _, line := range processed {
		fmt.Fprintln(out, line)
	}

	fmt.Fprintln(out, "\n--- Summary ---")
	fmt.Fprintf(out, "Lines read    : %d\n", linesRead)
	fmt.Fprintf(out, "Lines written : %d\n", len(processed))
	fmt.Fprintf(out, "Lines removed : %d\n", linesRemoved)
	fmt.Fprintf(out, "Rules applied : TrimWhitespace, ReplaceTODO, AllCapsToTitle, Lower, reverseWord\n")

	fmt.Println("\n--- Terminal Summary ---")
	fmt.Printf("✦ Lines read    : %d\n", linesRead)
	fmt.Printf("✦ Lines written : %d\n", len(processed))
	fmt.Printf("✦ Lines removed : %d\n", linesRemoved)
	fmt.Printf("✦ Rules applied : TrimWhitespace, ReplaceTODO, AllCapsToTitle, lower, reverseWord\n")
}
