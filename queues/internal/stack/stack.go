package stack

import "sync"

type Stack struct {
	mu    sync.Mutex
	stack []any
}

func NewStack(stack []any) *Stack {
	return &Stack{
		mu:    sync.Mutex{},
		stack: stack,
	}
}

func (s *Stack) Enqueue(item any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.stack = append(s.stack, item)
	return nil
}

func (s *Stack) Dequeue() (any, error) {
	if len(s.stack) < 1 {
		return nil, ErrStackIsEmpty
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	item := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return item, nil
}
