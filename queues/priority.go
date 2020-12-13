package queues

import (
	"errors"
)

// Enqueue
// Dequeue
// isEmpty
// isFull
// peek

type PQueue struct {
	items []int;
	front int;
	rear int;
	totalItems int;
}

func (queue *PQueue) Enqueue(item int) error {
	if queuefull := queue.isFull(); queuefull == true {
		return errors.New("can't work with full queue")
	}

	for i := len(queue.items)-1; i >= 0; i-- {

	}

	queue.items[queue.rear] = item;
	queue.rear = (queue.rear + 1) % len(queue.items);

	if (queue.totalItems + 1) > len(queue.items) {
		queue.totalItems = len(queue.items);
	} else {
		queue.totalItems++;
	}
	return nil;
}

func (queue *PQueue) isFull() bool {
	return len(queue.items) == queue.totalItems;
}

// func main() {
// 	var q PQueue = PQueue{
// 		items: make([]int, 4),
// 	};

// 	enqueue(&q, 20);
// 	enqueue(&q, 10);
// 	enqueue(&q, 30);
// 	enqueue(&q, 40);

// 	fmt.Println(q);
// }