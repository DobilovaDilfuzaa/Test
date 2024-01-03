package main

import "fmt"

func main(){
	var number int
	fmt.Scan(&number)

	lastDigit := number % 100 / 10
	fmt.Println(lastDigit)
}
