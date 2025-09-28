package structures

import "math/rand"

type Stack[T any] struct { // last in, first out
	contents []T
}

func (s *Stack[T]) ToSlice() []T { // copy of the internal slice
	return append(make([]T, 0), s.contents...)
}

func (s *Stack[T]) Clear() {
	s.contents = make([]T, 0)
}

func (s *Stack[T]) Size() int {
	return len(s.contents)
}

/*
Returns nil if q is empty
Otherwise returns the last element of q and removes it
*/
func (s *Stack[T]) Pop() *T {
	front := s.CheckBack()
	if front == nil {
		return nil
	}
	s.contents = s.contents[:len(s.contents)-1]
	return front
}

func (s *Stack[T]) PushBack(item T) {
	s.contents = append(s.contents, item)
}

func (s *Stack[T]) PushListBack(items []T) {
	s.contents = append(s.contents, items...)
}

// nil if empty
func (s *Stack[T]) CheckBack() *T {
	if len(s.contents) == 0 {
		return nil
	}
	return &s.contents[len(s.contents)-1]
}

func (s *Stack[T]) Shuffle() {
	rand.Shuffle(len(s.contents), func(i, j int) {
		s.contents[i], s.contents[j] = s.contents[j], s.contents[i]
	}) // may need to change to be deterministic, idk
}
