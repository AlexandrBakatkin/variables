package collections

import (
	"errors"
)

type Collections interface {
	Push(int)
	Pop() int
}

type StackInt struct {
	stack []int
}

func (s *StackInt) Push(val int) {
	s.stack = append(s.stack, val)
}

func (s *StackInt) Pop() (int, error) {
	var pop int
	if len(s.stack) <= 0 {
		err := errors.New("stack is empty")
		return 0, err
	}
	pop, s.stack = s.stack[len(s.stack)-1], s.stack[:len(s.stack)-1]
	return pop, nil
}
