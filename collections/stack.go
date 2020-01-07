package collections

import (
	"errors"
)

type Stack struct {
	stack []interface{}
}

// Push prepends an element to the back of the stack.
func (s *Stack) Push(val interface{}) {
	s.stack = append(s.stack, val)
}

// Pop removes and returns the element from the back of the stack.
func (s *Stack) Pop() (interface{}, error) {
	var push interface{}
	length := len(s.stack)

	if length <= 0 {
		err := errors.New("stack is empty")
		return 0, err
	}
	push, s.stack = s.stack[length-1], s.stack[:length-1]
	return push, nil
}
