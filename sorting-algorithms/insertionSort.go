package main

import "fmt"

func insertionSort(array []int) []int {

  for i :=0; i < len(array)-1; i++ {
    for j := i; j > 0; j-- {
      if array[j] < array[j-1] {
        array[j], array[j-1] = array[j-1], array[j]
      }
    }
  }
  return array
}

func main() {
  var unsorted = []int{5,3,2,7,1,4,6}

  fmt.Println(insertionSort(unsorted))
}
