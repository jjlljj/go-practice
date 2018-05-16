//DESIRED OUTPUT
//[ [ '0', '-', '0', '-', '0', '-', '0', '-', '0' ],
//  [ '-', '0', '-', '0', '-', '0', '-', '0', '-' ],
//  [ '-', '-', '0', '-', '0', '-', '0', '-', '-' ],
//  [ '-', '-', '-', '0', '-', '0', '-', '-', '-' ],
//  [ '-', '-', '-', '-', '0', '-', '-', '-', '-' ] ]

package main

import "fmt"

func createTable(size int) [][]string {
  var table [][]string
  width := size*2 - 1

  for i := 0; i < size; i++ {
    var row []string

    for j := 0; j < width; j++ {
      if j  >= 2*size-i || j < i {
        row = append(row, "-")
      } else if j%2 == 0 && i%2 == 0 {
        row = append(row, "0")
      } else if j%2 == 1 && i%2 == 1 {
        row = append(row, "0")
      } else {
        row = append(row, "-")
      }
    }
    fmt.Println(row)
    table = append(table, row)
  }
  return table
}

func main() {
  fmt.Println(createTable(4))
}
