package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  fmt.Print("Enter a phrase: ")
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')
  // fmt.Sscanf("%s", &input)
  fmt.Println(input)
}
