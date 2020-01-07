package collections

import (
	"errors"
)

type DequeueInt struct {
	queue []int
}

func (d *DequeueInt) Push(val int) {
	d.queue = append(d.queue, val)
}

func (d *DequeueInt) Pop() (int, error) {
	var pop int
	if len(d.queue) <= 0 {
		err := errors.New("stack is empty")
		return 0, err
	}
	pop, d.queue = d.queue[len(d.queue)-1], d.queue[1:]
	return pop, nil
}
