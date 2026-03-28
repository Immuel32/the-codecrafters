***LOOPS***



Loops 

A loop is use if you want to run the same code over and over again, each time with a different value.

***Types of Loops***
There is one type of loop in go
1. for loop


***The continue statement***

The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.

***The break statement***

The break statement is used to break/terminate the loop execution.

***Code Snippet***
```
package main
import ("fmt")

func main() {
  for i:=0; i < 2; i++ {
    fmt.Println(i)
  }
}
```
