package main

import "fmt"

var unsorted = []int{5,3,2,7,1,4,6}

func insertion(array []int) []int {

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
  fmt.Println(insertion(unsorted))
}
