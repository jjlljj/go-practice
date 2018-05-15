package main

import "fmt"

func dungeon(size int) [][]string {
  var dungeon [][]string
  
  for i:=0; i < size; i++ {
    var row []string

    for j:=0; j<size; j++ {
      row = append(row, cell(dungeon, row, i, j))
    }

    dungeon = append(dungeon, row)
  }

  return dungeon
}


func cell(dungeon [][]string, row []string, i int, j int) string {
  return "m"
}

func main() {
  fmt.Println(dungeon(4))
}
