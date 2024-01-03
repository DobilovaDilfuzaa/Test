package main

import "fmt"

func main(){
	var a, b, sum int

	fmt.Scan(&a, &b)

	squareA := a *a
	squareB := b * b

	sum = squareA + squareB
	fmt.Println(sum)
}