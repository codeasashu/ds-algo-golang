package queues

import (
	"errors"
)

// Enqueue
// Dequeue
// isEmpty
// isFull
// peek

// Queue is basic Queue
type Queue struct {
	items []int;
	front int;
	rear int;
	totalItems int;
}

// Enqueue adds items to queue
func (queue *Queue) Enqueue(item int) error {
	if queuefull := queue.IsFull(); queuefull == true {
		return errors.New("can't work with full queue")
	}

	queue.items[queue.rear] = item;
	// For circular arrays
	queue.rear = (queue.rear + 1) % len(queue.items);

	if (queue.totalItems + 1) > len(queue.items) {
		queue.totalItems = len(queue.items);
	} else {
		queue.totalItems++;
	}
	return nil;
}

// Dequeue removes items from queues
func (queue *Queue) Dequeue() int {
	var firstItem int = queue.items[queue.front];
	queue.items[queue.front] = 0;
	if (queue.front + 1) < len(queue.items) {
		// For circular arrays
		queue.front = (queue.front + 1) % len(queue.items);
	}
	if queue.totalItems < 0 {
		queue.totalItems = 0;
	} else {
		queue.totalItems--;
	}
	return firstItem;
}

// IsEmpty checks if queue is empty
func (queue *Queue) IsEmpty() bool {
	return len(queue.items) == 0;
}

// IsFull checks is queue is full
func (queue *Queue) IsFull() bool {
	return len(queue.items) == queue.totalItems;
}

// Peek checks first item in queue
func (queue *Queue) Peek() int {
	return queue.items[queue.front];
}

// Init inits a queue with items
func (queue *Queue) Init(length int) {
	queue.items = make([]int, length);
}