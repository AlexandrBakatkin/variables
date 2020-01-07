package collections

import "errors"

type Dequeue struct {
	dequeue []interface{}
}

func (dq *Dequeue) PushFirst(val interface{}) {
	arr := []interface{}{val}
	dq.dequeue = append(arr, dq.dequeue[:]...)
}

func (dq *Dequeue) PushLast(val interface{}) {
	dq.dequeue = append(dq.dequeue, val)
}

func (dq *Dequeue) PopFirst() (interface{}, error) {
	if len(dq.dequeue) == 0 {
		return 0, errors.New("Stack is empty")
	}
	firstVal := dq.dequeue[0]
	dq.dequeue = dq.dequeue[1:]
	return firstVal, nil
}

func (dq *Dequeue) PopLast() (interface{}, error) {
	if len(dq.dequeue) == 0 {
		return 0, errors.New("Stack is empty")
	}
	pop := dq.dequeue[len(dq.dequeue)-1]
	dq.dequeue = dq.dequeue[:len(dq.dequeue)-1]
	return pop, nil
}

func (dq *Dequeue) Length() int {
	return len(dq.dequeue)
}
