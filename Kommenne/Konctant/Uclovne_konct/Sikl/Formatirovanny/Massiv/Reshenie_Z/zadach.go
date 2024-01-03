package main

import "fmt"

func main() {
    var number, sum, digit int
    fmt.Scan(&number)

    
    for number > 0 {
        digit = number % 10 
        sum += digit        
        number /= 10        
    }

    fmt.Println(sum) 
}
