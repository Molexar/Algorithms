package queue

import "sync"

type Queue struct {
	mu    sync.Mutex
	queue []any
}

func NewQueue(queue []any) *Queue {
	return &Queue{
		mu:    sync.Mutex{},
		queue: queue,
	}
}

func (q *Queue) Enqueue(item any) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, item)

	return nil
}

func (q *Queue) Dequeue() (any, error) {

	if len(q.queue) < 1 {
		return nil, ErrQueueEmpty
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	item := q.queue[0]
	q.queue = q.queue[1:]
	return item, nil
}
