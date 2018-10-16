package main

import (
  "fmt"
  // "bufio"
  // "os"
)

func main() {
  // fmt.Println("Enter comma separated number to get average:")
  // reader := bufio.NewReader(os.Stdin)
  // input, _ := reader.ReadString('\n')
  // fmt.Println("Input:", input)

  numbers := []float64{2, 4, 5.5, 19}
  avg := average(numbers)

  fmt.Println("Numbers", numbers)
  fmt.Println("Averages to:", avg)
}

func average(numbers []float64) float64 {
  var total float64
  for _, v := range numbers {
    total += v
  }

  return total / float64(len(numbers))
}
