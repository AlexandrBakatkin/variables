package main

import (
	"flag"
	"fmt"
	"github.com/Bakatkin/learngo/variables/collections"
	"strings"
)

var (
	m = flag.Bool("n", true, "пропуск символа новой строки")
	sep = flag.String("s", " ", "разделитель")
)

func main (){
	n := 10
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	*m = false
	if !*m {
		fmt.Println()
		fmt.Println(n)
	}

	fmt.Println("-----------------stack")
	stack:= new (collections.Stack)
	fmt.Println(stack.Pop())
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack)

	stack2:= new (collections.Stack)
	stack2.Push(4)
	stack2.Push(5)
	stack2.Push(6)
	fmt.Println(stack2)
}