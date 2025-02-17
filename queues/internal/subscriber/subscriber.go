package subscriber

import (
	"errors"
	"fmt"

	"molexar.org/queues/internal/queue"
)

type Subscriber struct {
	queue *queue.PriorityQueue
}

func NewSubscriber(queue *queue.PriorityQueue) *Subscriber {
	return &Subscriber{
		queue: queue,
	}
}

func (s *Subscriber) Subscribe(handler func(item any) error) {
	go func() {
		for {
			item, err := s.queue.Dequeue()
			if err != nil {
				if errors.Is(err, queue.ErrQueueEmpty) {
					continue
				}
				fmt.Printf("error while reading message from qeueue: %s\n", err)
			}
			err = handler(item)
			if err != nil {
				fmt.Printf("error while handle message from queue: %s\n", err)
			}
			fmt.Printf("message: %s - was procceed\n", item)
		}
	}()
}
