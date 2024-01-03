

package main

import (
 "fmt"
 "time"
)

func work(x int) int {
 time.Sleep(1 * time.Second) 
 return x * x
}

func main() {
 cache := make(map[int]int)
 input := []int{3, 1, 5, 2, 3, 5, 3, 0, 3, 4}

 for _, num := range input {
  result, found := cache[num]
  if !found {
   result = work(num)
   cache[num] = result
  }
  fmt.Print(result, " ")
 }

 fmt.Println("time limit ok")
}