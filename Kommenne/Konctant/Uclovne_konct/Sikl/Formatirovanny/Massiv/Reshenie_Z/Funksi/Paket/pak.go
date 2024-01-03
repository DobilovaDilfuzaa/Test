package main

import "fmt"

func multiplyAndPrint(a *int, b *int) {
    result := *a * *b
    fmt.Println(result)
}

func main() {
    var num1, num2 int
    fmt.Scanf("%d %d", &num1, &num2)

    multiplyAndPrint(&num1, &num2)
}

