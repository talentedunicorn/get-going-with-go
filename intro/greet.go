package main

// Adding the string formatting lib
import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main() {
  fmt.Println("Welcome to the greeting app")

  reader := bufio.NewReader(os.Stdin)
  fmt.Printf("What's your name: ")
  name, _ := reader.ReadString('\n')
  name = strings.TrimSpace(name)

  if len(name) <= 1 {
    fmt.Println("No name passed, are you anonymous?")
  } else {
    fmt.Println("Hello,", name)
  }
}
