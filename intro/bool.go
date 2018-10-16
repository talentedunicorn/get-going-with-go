package main

import "fmt"

func main() {
  var exp bool = (true && false) || (false && true) || !(false && false)
  fmt.Println("Answer", exp)
}
