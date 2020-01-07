package collections

import "errors"

type Dequeue struct {
	dequeue []interface{}
}

// PushFirst prepends an element to the front of the dequeue.
func (dq *Dequeue) PushFirst(val interface{}) {
	arr := []interface{}{val}
	dq.dequeue = append(arr, dq.dequeue[:]...)
}

// PushLast prepends an element to the back of the dequeue.
func (dq *Dequeue) PushLast(val interface{}) {
	dq.dequeue = append(dq.dequeue, val)
}

// PopFirst removes and returns the element from the front of the dequeue.
func (dq *Dequeue) PopFirst() (interface{}, error) {
	if len(dq.dequeue) == 0 {
		return 0, errors.New("Stack is empty")
	}
	firstVal := dq.dequeue[0]
	dq.dequeue = dq.dequeue[1:]
	return firstVal, nil
}

// PopLast removes and returns the element from the back of the dequeue.
func (dq *Dequeue) PopLast() (interface{}, error) {
	if len(dq.dequeue) == 0 {
		return 0, errors.New("Stack is empty")
	}
	pop := dq.dequeue[len(dq.dequeue)-1]
	dq.dequeue = dq.dequeue[:len(dq.dequeue)-1]
	return pop, nil
}

// Length returns the count of elements in the dequeue
func (dq *Dequeue) Length() int {
	return len(dq.dequeue)
}
