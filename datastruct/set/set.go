package set 

type SimpleSet[T comparable] struct {
  m map[T]struct{}
}

func NewSimpleSet[T comparable]() *SimpleSet[T] {
  return &SimpleSet[T]{
    m: make(map[T]struct{}),
  }
}

func (s *SimpleSet[T]) Add(element T) {
  s.m[element] = struct{}{}
}


func (s *SimpleSet[T]) Has(element T) bool {
  _, ok := s.m[element]
  return ok
}