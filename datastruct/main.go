package main

import (
  "github.com/adrienDog/go118-seed/datastruct/set"
)

func main() {
  s := set.NewSimpleSet[string]()
  s.Add("lala")
}