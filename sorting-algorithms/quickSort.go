package main

import "fmt"

func quickSort(array []int) []int {
  if len(array) <= 1 {
    return array
  }

  var pivot int
  var l []int
  var r []int
  pivot, array = array[0], array[1:]

  for i := 0; i < len(array); i++ {
    if array[i] < pivot {
      l = append(l, array[i])
    } else {
      r = append(r, array[i])
    }
  }

  l, r = quickSort(l), quickSort(r)
  return append(append(l, pivot), r...)
}

func main() {
  var unsorted = []int{9,3,5,4,6,2,1,8,7}

  fmt.Println(quickSort(unsorted))
}
