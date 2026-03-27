package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var num, base string

		fmt.Println("Enter a number to convert")
		fmt.Scan(&num)
		if num == "exit" {
			fmt.Println("Goodbye!")
			break
		}
		fmt.Print("Enter it base(hex/bin/dec): ")
		fmt.Scan(&base)

		var decimal int64
		var err error

		switch base {
		case "hex":
			decimal, err = strconv.ParseInt(num, 16, 64)
		case "bin":
			decimal, err = strconv.ParseInt(num, 2, 64)
		case "dec":
			decimal, err = strconv.ParseInt(num, 10, 64)
		default:
			fmt.Println("Unknown base. Use 'hex', 'bin', or 'dec'")
			continue
		}

		if err != nil {
			fmt.Println("Invalid Number")
			continue
		}
		fmt.Println("Decimal:", decimal)
		fmt.Println("Binary:", strconv.FormatInt(decimal, 2))
		fmt.Println("Hexdicimal", strconv.FormatInt(decimal, 16))
	}
}
