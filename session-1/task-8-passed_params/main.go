package main

import "fmt"

// Function that demonstrates pass by value
func modifyValue(x int) {
    // 1. set the value in x to be 100

    // 2.print x
    fmt.Println("Inside function:")
}

// Function that demonstrates pass by reference
func modifyReference(x *int) {
    // 1. set the value in x to be 100

    // 2. print x
    fmt.Println("Inside function:")
}

func main() {
    // Pass by value example
    fmt.Println("Pass by value example:")
    number := 42
    fmt.Println("Before:", number)
    modifyValue(number)
    fmt.Println("After:", number)
    
    // Pass by reference example
    fmt.Println("\nPass by reference example:")
    number = 42
    fmt.Println("Before:", number)
    modifyReference(&number)
    fmt.Println("After:", number)
}

/* 
Task:
1. Create two functions that show the difference between:
   - Passing a value (makes a copy)
   - Passing a reference (can modify original)

*/
