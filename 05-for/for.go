package main

import "fmt"

// Loving the simplicity of not having for/do/while
func main() {
  i := 1

  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }

  for {
    fmt.Println("loop")
    break
  }
}
