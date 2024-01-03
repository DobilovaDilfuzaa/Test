package main

import (
    "errors"
     "fmt"
    
)

func divide(a, b int) (int, error) {
     if b == 0 {
      return 0, errors.New("Ошибка:")
     }
     return a / b, nil
}

func main() {
     var x, y int
     fmt.Scan(&x, &y)

     result, err := divide(x, y)
     if err != nil {
     fmt.Println("ошибка")
     } else {
     fmt.Println(result)
     }
}
