package main

import (
 "fmt"
 "strings"
)

func main() {
   var X, S string
 fmt.Scan(&X, &S)

 index := strings.Index(X, S)

 if index == -1 {
  fmt.Println(-1)
 } else {
  fmt.Println(index)
 }
}

