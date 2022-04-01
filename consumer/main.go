package main

import (
  "fmt"
  "github.com/adrienDog/go118-seed/datastruct/set"
  api "github.com/adrienDog/go118-seed/api/go/services"
)

func main() {
  set := set.NewSimpleSet[string]()

  set.Add("lala")

  fmt.Printf("Set contains %d elements\n", set.Size())

  req := &api.GetGreetingsRequest{}
  fmt.Println(req)
}
