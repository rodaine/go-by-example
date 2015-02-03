package main

import "fmt"
import "time"

// No fall through with comma-delimited cases is $$$
func main() {

  i := 2
  fmt.Print("write ", i, " as ")
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend!")
  default:
    fmt.Println("It's a weekday T.T")
  }

  t := time.Now()
  switch { // LLOOOOOVVVVVEEEEEEEEE
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  default:
    fmt.Println("It's after noon")
  }
}
