package collections

import "errors"

type Queue []int

var lengthQueue = 0

func (queue Queue) Push (val int){
	queue = append(queue, val)

}

func (queue *Queue) Count() int {
	return lengthQueue
}

func (queue *Queue) Pop() (int, error) {
	if lengthQueue <= 0 {
		return 0, errors.New("Stack is empty")
	}
	temp := *queue
	firstVal := temp[0]
	*queue = temp[1:]
	return firstVal, nil
}