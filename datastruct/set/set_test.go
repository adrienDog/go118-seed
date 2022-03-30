package set

import (
  "testing"
)

func TestGenericSetCanBeConstructed(t *testing.T) {
  set := NewSimpleSet[string]()

  if (set == nil) {
    t.Fatal("set should not be nil")
  }
}

type Set[T comparable] interface {
  Add(element T)
  Has(element T) bool
}

func TestGenericSetImplementsSet(t *testing.T) {
  var set Set[string] = NewSimpleSet[string]()

  set.Add("lala")
  if (!set.Has("lala")) {
    t.Error("'lala' should be present")
  }
  if (set.Has("no")) {
    t.Error("'no' should NOT be present")
  }
}