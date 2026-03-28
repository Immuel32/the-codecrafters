***Variable***


Variables are like containers that store values for use or referencing back to later


***Variable Types***

1. Int: Which stands for integer store whole numbers either positive or negative e.g 345 or -478
2. Float32: which stores floating point numbers that is numbers with decimals e.g 3.142
3. String: They store text base characters e.g "name" and which are enclose in double quotes
4. Bool: which stands for boolean they store values whch is either true or false


***Declaring or creating variables***

Variables can be declared in two wsys

1. Using the var key word: in this case you declare the variable type along with its value

2. Using the := sign in this case there is no need to declare the variable type along its value 


***Code Snippet***
```
func main() {
    var num int
    num = 42
    fmt.Println("Number:", num)
}

func main() {
    num := 42
    fmt.Println(num)
}
```


***Variable declaration without initialization***

in go when you declare a variable without initializing them the value will be set to their default value.

int = 0
string = ""
float = 0.0
bool = false

***Naming Variables***
Variable are named using the following styles

1. Camel Case e.g

    myVariableName

2. Pascal Case e.g

    MyVariableName 

3. Snake Case e.g

    My_Variabl_Name 

***Naming Rules***
- variable name must start with a letter or an underscore character
- variable name cannot start with a digit
- variable name cannot contain space
- variable name must not contain go keyword