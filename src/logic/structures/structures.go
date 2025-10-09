package structures

import "math/rand"

type Stack[T any] struct { // last in, first out
	Contents []T
}

func (s *Stack[T]) ToSlice() []T { // copy of the internal slice
	return append(make([]T, 0), s.Contents...)
}

func (s *Stack[T]) Clear() {
	s.Contents = make([]T, 0)
}

func (s *Stack[T]) Size() int {
	return len(s.Contents)
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
	s.Contents = s.Contents[:len(s.Contents)-1]
	return front
}

func (s *Stack[T]) PushBack(item T) {
	s.Contents = append(s.Contents, item)
}

func (s *Stack[T]) PushListBack(items []T) {
	s.Contents = append(s.Contents, items...)
}

// nil if empty
func (s *Stack[T]) CheckBack() *T {
	if len(s.Contents) == 0 {
		return nil
	}
	return &s.Contents[len(s.Contents)-1]
}

func (s *Stack[T]) Shuffle() {
	rand.Shuffle(len(s.Contents), func(i, j int) {
		s.Contents[i], s.Contents[j] = s.Contents[j], s.Contents[i]
	}) // may need to change to be deterministic, idk
}
