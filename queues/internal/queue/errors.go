package queue

import "errors"

var (
	ErrQueueEmpty  = errors.New("queue is empty")
	ErrQueueIsFull = errors.New("queue is full")
)

var ErrUnknownPriority = errors.New("unknown priority")
