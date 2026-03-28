***Conditions***



Condition

Conditional statements are used to perform different actions based on different conditions.

***Types of conditional statement***
1. if statement : if statement are use to specify a block of Go code to be executed if a condition is true.
2. if else statement : The else statement is use to specify a block of code to be executed if the condition is false.
3. else if statement : The else if statement is use to specify a new condition if the first condition is false.
4. nested if statement : Used when the condition is true to true


***Code Snippet***
```
func main() {
    Age := 20
    if Age >= 10 {
        fmt.Println("Age is greater then 10")
        if Age > 18{
            fmt.Println("Age is still more than 18")
        } else { 
            fmt.Println("Age is less than 10")
        }
    }

}
```