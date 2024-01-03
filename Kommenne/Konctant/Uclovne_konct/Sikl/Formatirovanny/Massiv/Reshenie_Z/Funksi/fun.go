package main

import "fmt"

func fibonacci(n int) int {
    if n <= 0 {
        return -1 
    } else if n == 1 || n == 2 {
        return 1
    } else {
        a, b := 1, 1
        for i := 2; i < n; i++ {
            a, b = b, a+b
        }
        return b
    }
}

func main() {
    var n int
    fmt.Scan(&n)

    result := fibonacci(n)
    fmt.Println(result)
}



