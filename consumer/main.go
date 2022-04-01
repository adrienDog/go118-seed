package main

import (
  "fmt"
  "github.com/adrienDog/go118-seed/datastruct/set"
)

func main() {
  set := set.NewSimpleSet[string]()

  set.Add("lala")

  fmt.Printf("Set contains %d elements\n", set.Size())
}