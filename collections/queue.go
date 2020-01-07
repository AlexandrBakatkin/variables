package collections

import (
	"errors"
)

type Queue struct {
	queue []interface{}
}

func (q *Queue) Push(val interface{}) {
	q.queue = append(q.queue, val)
}

func (q *Queue) Pop() (interface{}, error) {
	length := len(q.queue)

	if length <= 0 {
		err := errors.New("queue is empty")
		return 0, err
	}
	pop := q.queue[0]
	q.queue = q.queue[1:]
	return pop, nil
}
