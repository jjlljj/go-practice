package main

import "fmt"

var unsorted = []int{9,3,5,4,6,2,1,8,7}

func merge(array []int) []int {
  if len(array) <= 1 {
    return array
  }

  var center int= len(array)/2 
  var a []int = array[0:center]
  var b []int = array[center:]

  return join(merge(a), merge(b))
}

func join(a, b []int) []int {
  var sorted []int

  for len(a) > 0 && len(b) > 0 {
    var x int
    if a[0] < b[0] {
      x, a = a[0], a[1:]
    } else {
      x, b = b[0], b[1:]
    }
    sorted = append(sorted, x)
  }

  if len(a) > 0 {
    sorted = append(sorted, a...)
  } else {
    sorted = append(sorted, b...)
  }

  return sorted
}

func main() {
  fmt.Println(merge(unsorted))
}
