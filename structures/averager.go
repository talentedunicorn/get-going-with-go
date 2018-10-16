package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

func main() {
  fmt.Println("Enter comma separated number to get average:")
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')

  inputMap := strings.Split(input, ",")
  var numbers []float64
  for _, n := range inputMap {
    number, _ := strconv.ParseFloat(n, 64)
    numbers = append(numbers, number)
  }

  // numbers := []float64{2, 4, 5.5, 19}
  avg := average(numbers)

  fmt.Println("Numbers", numbers)
  fmt.Println("Averages to:", avg)
}

func average(numbers []float64) (res float64) {
  var total float64
  for _, v := range numbers {
    total += v
  }

  res = total / float64(len(numbers))
  return res
  // Or
  // return total / float64(len(numbers))
}
