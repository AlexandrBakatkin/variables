package main

import (
	"fmt"
	"github.com/Bakatkin/learngo/variables/collections"
)

func main() {
	var stack = collections.Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Pop()
	fmt.Println(stack)

	var queue = collections.Queue{}
	queue.Push("one")
	queue.Push("two")
	queue.Push("three")

	names := []string{"1", "2", "3"}
	for i := range names {
		fmt.Println(names[i])
	}
	queue.Pop()
	fmt.Println(queue)
}
