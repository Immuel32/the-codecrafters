package main

import "fmt"

func main() {
	var num1, num2 int

	for {
		fmt.Println("Enter the first number")
		fmt.Scan(&num1)

		fmt.Println("Enter the second number")
		fmt.Scan(&num2)

		fmt.Println("Select an operator")
		fmt.Println("1 - Addition")
		fmt.Println("2 - Sutraction")
		fmt.Println("3 - Multiplication")
		fmt.Println("4 - Division")
		fmt.Println("5 - Exit")
		fmt.Println("6 - Help")
		fmt.Println()

		var choice int
		fmt.Scan(&choice)

		if choice == 6 {

			fmt.Println("1 - Addition")
			fmt.Println("2 - Subtraction")
			fmt.Println("3 - Multiplication")
			fmt.Println("4 - Division")
			fmt.Println("5 - Exit")
			fmt.Println("6 - Help")
			fmt.Println()

		}
		if choice == 5 {
			fmt.Println("Thanks, Goodbye!!")
			break
		}

		switch choice {
		case 1:
			fmt.Println(num1 + num2)
		case 2:
			fmt.Println(num1 - num2)
		case 3:
			fmt.Println(num1 * num2)
		case 4:
			if num2 != 0 {
				fmt.Println(num1 / num2)
			} else {
				fmt.Println("Cannot divide by zero")
			}
		}

	}

}
