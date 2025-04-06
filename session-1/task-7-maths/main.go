package main

import "fmt"

// Basic calculator function
func calculate(a, b int) {
    // Addition
    fmt.Println("Adding numbers:")
    
    // Subtraction
    fmt.Println("\nSubtracting numbers:")
    
    // Multiplication
    fmt.Println("\nMultiplying numbers:")
    
    // Division (with check for divide by zero)
    if b != 0 {
        fmt.Println("\nDividing numbers:")
        //
        fmt.Println("\nCalculating remainder:")
        //
    } else {
        fmt.Println("Cannot divide by zero")
    }
}

func main() {
    // Example with two numbers
    fmt.Println("Calculator with 10 and 3:")
    calculate(10, 3)
    
    fmt.Println("\nCalculator with 15 and 5:")
    calculate(15, 5)
}

/* 
Task:
1. Use basic arithmetic operators to perform calculations
   - Addition (+)
   - Subtraction (-)
   - Multiplication (*)
   - Division (/)
   - Modulus (%)
*/
