package main

import (
  "fmt"
  "math/rand"
  "time"
)

func dungeon(size int) [][]string {
  var dungeon [][]string
  
  for i:=0; i < size; i++ {
    var row []string

    for j:=0; j<size; j++ {
      row = append(row, cell(dungeon, row, i, j, size))
    }
    
    fmt.Println(row)
    dungeon = append(dungeon, row)
  }

  return dungeon
}


func cell(dungeon [][]string, row []string, i int, j int, size int) string {
  var possible = []string{"M", "t", "#", "-"}

  if (j > 0 && row[j-1] == "M") ||
    i > 0 && (
      (j > 0 && dungeon[i-1][j-1] == "M") ||
      dungeon[i-1][j] == "M" ||
      (j < size-1 && dungeon[i-1][j+1] == "M")) {
    possible = possible[1:]
  }  

  rand.Seed(time.Now().UnixNano())
  var random int = rand.Intn(len(possible))
  
  return possible[random]
}

func main() {
  fmt.Println(dungeon(8))
}
