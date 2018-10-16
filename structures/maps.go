package main

import "fmt"

func main() {
  fmt.Println("Maps: unordered collection of key-value pairs")

  // Definition
  // scores := make(map[string]int) // Roughly map[key_type]value_type
  // scores["qualifying"] = 3
  // scores["semi-finals"] = 2
  // scores["finals"] = 5

  // OR
  scores := map[string]int {
    "qualifying": 3,
    "semi-finals": 2,
    "finals": 5,
  }

  fmt.Println("Scores are:", scores)

  // Accessing map values
  fmt.Println("Finals score was:", scores["finals"])

  // Deleting map items
  delete(scores, "finals")
  fmt.Println("Deleted finals:", scores)


  // Multi-level maps
  person := map[string]map[string]string {
    "name": map[string]string {
      "first": "John",
      "last": "Mayor",
    },
  }
  fmt.Println("Miltilevel maps:", person)
}
