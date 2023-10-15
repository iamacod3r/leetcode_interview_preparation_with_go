package datastructure

import (
	"container/list"
	"fmt"
	"sync"
)

type Stack struct {
	stack *list.List
}

func (s *Stack) Push(item any) {
	s.stack.PushFront(item)
}

func (s *Stack) Pop() (any, error) {
	if s.stack.Len() > 0 {
		element := s.stack.Front()
		defer s.stack.Remove(element)
		return element.Value, nil
	}

	return nil, fmt.Errorf("pop error: stack is empty")
}

func (s *Stack) Peek() (any, error) {

	if s.stack.Len() > 0 {
		return s.stack.Front().Value, nil
	}
	return nil, fmt.Errorf("peek error: stack is empty")
}

func (s *Stack) Size() int {
	return s.stack.Len()
}

func (s *Stack) IsEmpty() bool {
	return s.stack.Len() == 0
}

type StackSlice struct {
	stack []any
	lock  sync.RWMutex
}

func (s *StackSlice) Push(item any) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.stack = append(s.stack, item)
}

func (s *StackSlice) Pop() (any, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	pos := len(s.stack)
	if pos > 0 {
		item := s.stack[pos-1]
		s.stack = s.stack[:pos-1]
		return item, nil
	}

	return nil, fmt.Errorf("pop error: stack is empty")
}

func (s *StackSlice) Peek() (any, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	pos := len(s.stack)
	if pos > 0 {
		return s.stack[pos-1], nil
	}

	return nil, fmt.Errorf("peek error: stack is empty")
}

func (s *StackSlice) Size() int {
	return len(s.stack)
}

func (s *StackSlice) IsEmpty() bool {
	return len(s.stack) == 0
}
