// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: Emmanuel Elaigwu
// Squad:  The Gophers Group

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 string
	var num2 string

	for {
		fmt.Println("Enter the first number")
		fmt.Scan(&num1)
		num1, err := strconv.ParseFloat(num1, 64)
		if err != nil {
			fmt.Println("Invalid Input")
			continue
		}

		fmt.Println("Enter the second number")
		fmt.Scan(&num2)
		num2, err := strconv.ParseFloat(num2, 64)
		if err != nil {
			fmt.Println("Invalid Input")
			continue
		}

		fmt.Println("Select an operator: ")
		fmt.Println("1: Addition|| 2: Sutraction|| 3: Multiplication|| 4: Division| |5: Exit|| 6: Help")

		var choice string
		fmt.Scan(&choice)

		if choice == "6" {

			fmt.Println("1 - Addition")
			fmt.Println("2 - Subtraction")
			fmt.Println("3 - Multiplication")
			fmt.Println("4 - Division")

		}
		if choice == "5" {
			fmt.Println("Thanks, Goodbye!!")
			break
		}

		if choice != "1" && choice != "2" && choice != "3" && choice != "4" && choice != "5" && choice != "6" {
			fmt.Println("...Not a valid input,. Try again!!")
		}

		switch choice {
		case "1":
			fmt.Println(num1 + num2)
		case "2":
			fmt.Println(num1 - num2)
		case "3":
			fmt.Println(num1 * num2)
		case "4":
			if num2 != 0 {
				fmt.Println(num1 / num2)
			} else {
				fmt.Println("Cannot divide by zero")
			}
		}

	}

}
