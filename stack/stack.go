package stack

import "errors"

var StackEmpty error = errors.New("Stack is Empty")

type Stack[T any] struct {
	items []T
	len int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
} 

func (s *Stack[T]) IsEmpty() bool {
	return s.len == 0 
}

func (s *Stack[T]) Pop() {
	if ! s.IsEmpty() {
		s.items = s.items[:s.len-1]
		s.len -= 1
	}
}

func (s *Stack[T]) Top() (*T,error)  {
	if s.IsEmpty() {
		return nil,StackEmpty
	}
	return &s.items[s.len-1],nil
}

func (s *Stack[T]) Push(element T) {
	s.items = append(s.items, element)
	s.len += 1
}
