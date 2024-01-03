
package main

import (
 "fmt"
 "math"
)

func main() {
 var i float64
 fmt.Scan(&i)

 if i <= 0 {
  fmt.Printf("число %.2f не подходит\n", i)
 } else if i > 10000 {
  fmt.Printf("%.2e\n", i)
 } else {
  res := math.Pow(i, 2)
  res = math.Floor(res*10000) / 10000 
  fmt.Printf("%.4f\n", res)
 }
}