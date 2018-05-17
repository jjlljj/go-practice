package main

import (
  "fmt"
  "math/rand"
  "time"
)

func dungeon(size int) [][][]string {
  var dungeon [][][]string
  
  rand.Seed(time.Now().UnixNano())
  var downX int = rand.Intn(size)
  var downY int = rand.Intn(size)
  var upX int = size * 2
  var upY int = size * 2

  for z:=0; z < size; z++ {
    dungeon = append(dungeon, level(size, downX, downY, upX, upY, z))
    upX = downX
    upY = downY

    if downX == upX {
      downX = rand.Intn(size)
    }

    if downY == upY {
      downY = rand.Intn(size)
    }
  }

  return dungeon
}

func level(size, downX, downY, upX, upY, z int) [][]string {
  var level [][]string
  
  for y:=0; y < size; y++ {
    var row []string

    for x:=0; x<size; x++ {
      cell := cell(level, row, x, y, size, downX, downY, upX, upY, z)
      row = append(row, cell)
    }
    
    level = append(level, row)
  }

    fmt.Println(level)
  return level
}


func cell(level [][]string, row []string, x, y, size, downX, downY, upX, upY, z int) string {
  if x == downX && y == downY && z != size-1{
    return "d"
  }

  if x == upX && y == upY {
    return "u"
  }

  var possible = []string{"M", "t", "#", "-"}

  if (x > 0 && row[x-1] == "M") ||
    y > 0 && (
      (x > 0 && level[y-1][x-1] == "M") ||
      level[y-1][x] == "M" ||
      (x < size-1 && level[y-1][x+1] == "M")) {
    possible = possible[1:]
  }  

  var random int = rand.Intn(len(possible))
  
  return possible[random]
}

func main() {
  fmt.Println(dungeon(4))
}
