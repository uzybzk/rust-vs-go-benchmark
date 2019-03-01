package main

import (
    "fmt"
    "time"
)

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciIterative(n int) int {
    if n <= 1 {
        return n
    }
    
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}

func main() {
    n := 40
    
    // Recursive version
    start := time.Now()
    result := fibonacci(n)
    recursiveDuration := time.Since(start)
    
    fmt.Printf("Go Recursive Fibonacci(%d) = %d\n", n, result)
    fmt.Printf("Time: %v\n", recursiveDuration)
    
    // Iterative version
    start = time.Now()
    result = fibonacciIterative(n)
    iterativeDuration := time.Since(start)
    
    fmt.Printf("Go Iterative Fibonacci(%d) = %d\n", n, result)
    fmt.Printf("Time: %v\n", iterativeDuration)
}