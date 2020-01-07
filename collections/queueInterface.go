package collections

import (
	"errors"
)

type QueueInt struct {
	queue []int
}

func (q *QueueInt) Push(val int) {
	q.queue = append(q.queue, val)
}

func (q *QueueInt) Pop() (int, error) {
	var pop int
	if len(q.queue) <= 0 {
		err := errors.New("stack is empty")
		return 0, err
	}
	pop, q.queue = q.queue[len(q.queue)-1], q.queue[1:]
	return pop, nil
}
