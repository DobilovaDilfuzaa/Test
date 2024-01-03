
package main

import (
    
    "fmt"


)
func main(){
	var a int
	fmt.Println(a)
}
/*

func main() {
    value1, value2, operation := readTask()
    typeValue1 := fmt.Sprintf("%T", value1)
    typeValue2 := fmt.Sprintf("%T", value2)
    if typeValue1 == "float64" && typeValue2 == "float64" {
        switch operation {
        case "-":
            result := value1 - value2
            fmt.Printf("%.4f", float64(result))
        case "+":
            result := value1 + value2
            fmt.Printf("%.4f", float64(result))
        case "*":
            result := value1 * value2
            fmt.Printf("%.4f", float64(result))
        case "/":
            result := value1 / value2
            fmt.Printf("%.4f", float64(result))
        }
        os.Exit(0)
    }
    if typeValue1 != "float64" || typeValue2 != "float64" {
        if typeValue1 != "float64" {
            typeValue := value1
            fmt.Printf("value=%v: %T", typeValue, typeValue)
            os.Exit(0)
        }
        if typeValue2 != "float64" {
            typeValue := value2
            fmt.Printf("value=%v: %T", typeValue, typeValue)
            os.Exit(0)
        }
    }
    fmt.Println("неизвестная операция")
    os.Exit(0)
}*/