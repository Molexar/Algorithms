package queue

import "errors"

type Priority int

const (
	Low Priority = iota
	Middle
	High
)

type PriorityQueue struct {
	low    *Queue
	middle *Queue
	high   *Queue
}

func NewPriorityQueue(low, middle, high []any) *PriorityQueue {
	return &PriorityQueue{
		low:    NewQueue(low),
		middle: NewQueue(middle),
		high:   NewQueue(high),
	}
}

func (pq *PriorityQueue) Enqueue(item any, priority Priority) error {
	switch priority {
	case Low:
		return pq.low.Enqueue(item)
	case Middle:
		return pq.middle.Enqueue(item)
	case High:
		return pq.high.Enqueue(item)
	}
	return ErrUnknownPriority
}

func (pq *PriorityQueue) Dequeue() (any, error) {
	item, err := pq.high.Dequeue()
	if err != nil && !errors.Is(err, ErrQueueEmpty) {
		return item, err
	}

	item, err = pq.middle.Dequeue()
	if err != nil && !errors.Is(err, ErrQueueEmpty) {
		return item, err
	}

	item, err = pq.low.Dequeue()
	if err != nil {
		return nil, err
	}

	return item, err
}
