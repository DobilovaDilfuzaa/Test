package main
import (
	"fmt"
	"strconv"
   )
   
   func main() {
	fn := func(num uint) uint {
	 filtered := ""
	 strNum := strconv.FormatUint(uint64(num), 10)
	 for _, char := range strNum {
	  digit, _ := strconv.Atoi(string(char))
	  if digit%2 == 0 && digit != 0 {
	   filtered += string(char)
	  }
	 }
	 if filtered == "" {
	  return 100
	}
	result, _ := strconv.ParseUint(filtered, 10, 64)
	return uint(result)
	}
	fmt.Println(fn(727178))
   }