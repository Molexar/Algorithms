package main

import (
	"fmt"

	"molexar.org/queues/internal/publisher"
	"molexar.org/queues/internal/queue"
	"molexar.org/queues/internal/subscriber"
)

func main() {
	queue := queue.NewPriorityQueue(make([]any, 2), make([]any, 2), make([]any, 2))

	subscriber := subscriber.NewSubscriber(queue)
	publisher := publisher.NewPublisher(queue)

	subscriber.Subscribe(func(item any) error {
		fmt.Printf("task completed: %s", item)
		return nil
	})

	publisher.StartPublishing("tasks.txt")
}
