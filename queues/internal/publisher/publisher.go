package publisher

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"molexar.org/queues/internal/queue"
)

type Publisher struct {
	queue *queue.PriorityQueue
}

func NewPublisher(queue *queue.PriorityQueue) *Publisher {
	return &Publisher{
		queue: queue,
	}
}

func (p *Publisher) StartPublishing(filename string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Ошибка при создании watcher:", err)
	}
	defer watcher.Close()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error while opening file: %s\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	err = watcher.Add("example.txt")
	if err != nil {
		log.Fatal("Ошибка при добавлении файла в watcher:", err)
	}

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {

				for scanner.Scan() {
					line := scanner.Text()
					parts := strings.Fields(line)

					if len(parts) != 2 {
						fmt.Printf("incorrect task from file: %s\n", line)
						continue
					}

					queueName, item := parts[0], parts[1]

					var priority queue.Priority
					switch queueName {
					case "LOW":
						priority = queue.Low
					case "MIDDLE":
						priority = queue.Middle
					case "HIGH":
						priority = queue.High
					default:
						fmt.Printf("incorrect priority name: %s\n", queueName)
						continue
					}

					err := p.queue.Enqueue(item, priority)
					if err != nil {
						fmt.Printf("error while publishing to queue: %s\n", err)
					}
				}
			}
		case err := <-watcher.Errors:
			log.Println("Ошибка watcher:", err)
		}
	}
}
