package collections

import "errors"

type Dequeue []int

var lengthDequeue = 0

func (dequeue *Dequeue) PushFirst(val int){
	temp := []int{val}
	temp2 := *dequeue
	temp = append(temp, temp2[:]...)
	*dequeue = temp
	lengthDequeue++
}

func (dequeue *Dequeue) PushLast(val int){
	*dequeue = append(*dequeue, val)
	lengthDequeue++
}

func (dequeue *Dequeue) PopFirst() (int, error) {
	if lengthQueue <= 0 {
		return 0, errors.New("Stack is empty")
	}
	temp := *dequeue
	firstVal := temp[0]
	*dequeue = temp[1:]
	return firstVal, nil
}

func (dequeue *Dequeue) PopLast() (int, error) {
	if lengthQueue <= 0 {
		return 0, errors.New("Stack is empty")
	}
	temp := *dequeue
	pos := len(*dequeue) - 1
	last := temp[pos]
	*dequeue = temp[:pos]
	return last, nil
}

func (dequeue *Dequeue) Count() int {
	return lengthDequeue
}