package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

func main() {
  var numbers []float64
  fmt.Println("Enter comma separated number to get average:")
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')

  inputMap := strings.Split(input, ",")

  for _, i := range inputMap {
    cleanedString := strings.TrimSpace(i)
    number, _ := strconv.ParseFloat(cleanedString, 64)
    numbers = append(numbers, number)
  }

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
