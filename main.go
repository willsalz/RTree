package main

import (
  "fmt"
  "github.com/paulmach/go.geo"
)

func main() {
  fmt.Printf("Hello, World!\n")
  p := geo.NewPoint(38, -122)
  fmt.Printf("%f\n", p[0])
}
