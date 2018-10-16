package main

import "fmt"

var (
  firstName = "James"
  lastName = "Brown"
)

func main() {
  var fullName [2]string
  fullName[0] = firstName
  fullName[1] = lastName

  fmt.Println(fullName)
  fmt.Println("Length is:", len(fullName))


  var floats = [5]float64{4,6,8,10,12}
  fmt.Println("An array of 5 floats:", floats)

  // Slices
  var slice []float64
  fmt.Println("A slice:", slice)

  slice = make([]float64, 3)
  fmt.Println("Make slice with size 3:", slice)

  // Appending
  sliceAppended := append(slice, 3, 4, 5)
  fmt.Println("Append 3, 4 and 5 to slice:", sliceAppended)

  sliceFromArray := floats[:3]
  fmt.Println("Slice from floats array:", sliceFromArray)

  // Copying slices
  var sliceOri = []int{2,3,4}
  sliceCopy := make([]int, 2)
  fmt.Println("SliceOri:", sliceOri)
  copy(sliceCopy, sliceOri)
  fmt.Println("SliceCopy:", sliceCopy)
}
