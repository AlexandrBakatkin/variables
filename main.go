package main

import (
	"fmt"
	"github.com/Bakatkin/learngo/variables/collections"
)

func main() {
	var stack = collections.StackInt{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	fmt.Println(stack.Pop())
	fmt.Println(stack)

	var queue = collections.QueueInt{}
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Pop()

	fmt.Println(queue)
}
