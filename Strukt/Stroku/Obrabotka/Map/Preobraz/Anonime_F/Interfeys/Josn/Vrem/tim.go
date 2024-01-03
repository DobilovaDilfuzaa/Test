package main


import (
    "fmt"
    "time"
)

func main() {
    var s string
    fmt.Scan(&s)
    parseTime, _ := time.Parse(time.RFC3339,s)
    fmt.Println(parseTime.Format(time.UnixDate))
}
    
    
 