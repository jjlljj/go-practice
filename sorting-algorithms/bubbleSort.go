package main

import "fmt" 

func bubbleSort(array []int) []int {

  for i := 0; i < len(array)-1; i++ {
    for j :=0; j < len(array)-1; j++ {
      if array[j] > array[j+1] {
        array[j], array[j+1] = array[j+1], array[j]
      }
    }
  }
  return array
}

func main() {
  var unsorted = []int{3,2,5,1,4}

    fmt.Println(bubbleSort(unsorted))
}

