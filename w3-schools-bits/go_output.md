***OUTPUT***


Output

Output is the method the machine communicate withe the outside world 

***Forms of Output***
There three forms of output in go
- Println()
- Printf()
- Print()

Println() function
This function print output with white space added between the arguments and with a newline


Printf() function
This function first format its argument base on the formating verb giving before printing the output

The Print() function
This function prints its arguments with their default format.

***Formatting Verbs***
formatting verbs are use along side with the printf() for printig out arguments

***Types***

%v : Prints the value in the default format
%T : Print the Type of the value
%s : Prints the value as plain string
%q : Prints the value as a double-quoted string
%#v : Print the value in go syntax format

***Code Snippet***
```
func main() {
    greet := "Hello, World!"
    fmt.Printf("%v\n", greet)
    fmt.Printf("%T", greet)

    //Output: Hello, World!
    //Output: string
}
```