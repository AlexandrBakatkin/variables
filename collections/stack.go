package collections

import "errors"

type Stack []int

var lengthStack = 0

func (stack *Stack) Push (val int){
	*stack = append(*stack, val)
	lengthStack++
}

func (stack *Stack) Pop() (int, error) {
	if lengthStack <= 0 {
		return 0, errors.New("Stack is empty")
	}
	temp := *stack
	pos := len(*stack) - 1
	pop := temp[pos]
	*stack = temp[:pos]
	lengthStack--
	return pop, nil
}

func (stack *Stack) Count() int {
	return lengthQueue
}