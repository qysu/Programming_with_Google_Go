package main

import (
  "fmt"
  "strings"
)

func main()  {
  var input string
  var lower string

  fmt.Printf("Enter a string: ")
  fmt.Scan(&input)

  lower = strings.ToLower(input)

  if strings.HasPrefix(lower, "i") && strings.ContainsAny(lower, "a") && strings.HasSuffix(lower, "n") {
    fmt.Printf("Found!\n")
  } else {
    fmt.Printf("Not Found!\n")
  }
}
