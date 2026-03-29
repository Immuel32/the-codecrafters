package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		var num string
		fmt.Println()
		fmt.Println("Enter a number to convert:")
		fmt.Scan(&num)
		num = strings.TrimSpace(num)
		if num == "exit" {
			fmt.Println("Thank You,.Goodbye!")
			break
		}
		fmt.Println()
		fmt.Println("Select the base of the number:\n ")
		fmt.Println("1 - Hex")
		fmt.Println("2 - Bin")
		fmt.Println("3 - Dec")
		fmt.Println()

		var baseChoice int
		fmt.Scan(&baseChoice)

		var decimal int64
		var err error

		switch baseChoice {
		case 1:
			decimal, err = strconv.ParseInt(num, 16, 64)
			if err != nil {
				fmt.Println("Is not valid hex")
				fmt.Println()
				continue
			}
		case 2:
			decimal, err = strconv.ParseInt(num, 2, 64)
			if err != nil {
				fmt.Println()
				fmt.Println("Is not valid binary")
				continue
			}
		case 3:
			decimal, err = strconv.ParseInt(num, 10, 64)
			if err != nil {
				fmt.Println()
				fmt.Println("Is not valid decimal")
				fmt.Println()
				continue
			}
		default:
			fmt.Println("Invalid input.,Please enter")
			fmt.Println(1, 2, 3)
			fmt.Println()
			continue
		}

		switch baseChoice {
		case 1, 2:
			fmt.Println("Decimal:", decimal)
		case 3:
			fmt.Println()
			fmt.Println("Binary:", strconv.FormatInt(decimal, 2))
			fmt.Println("Hexdicimal", strings.ToUpper(strconv.FormatInt(decimal, 16)))
			fmt.Println()
		}
	}
}
