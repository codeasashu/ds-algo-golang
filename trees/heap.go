package trees

type Heap struct {
	items [10]int;
	size int;
}

func getRoot(heap Heap) int {
	if len(heap.items) > 0 {
		return heap.items[0];
	}
	return -1;
}

// Insert items into heap
func (heap *Heap) Insert(value int) {
	heap.items[heap.size] = value;
	heap.size++;
	heap.bubbleUp();
}

// Remove item from top of heap
func (heap *Heap) Remove() {
	heap.items[0] = heap.items[(heap.size -1)]; // Replace root with last item
	heap.size--; // Remove last item
	
	heap.bubbleDown();
}

func (heap Heap) bubbleDown() {
	var index = 0;
	for index <= heap.size && heap.isValidParent(index) == false {
		var largerIndex = heap.largerIndex(index);
		heap.swap(index, largerIndex);
		index = largerIndex;
	}
}

func (heap Heap) largerIndex(index int) int {
	var largerIndex int;
	if heap.items[index] < heap.leftChild(index) {
		largerIndex = leftChildIndex(index);
	} else {
		largerIndex = rightChildIndex(index);
	}
	return largerIndex;
}

func (heap Heap) isValidParent(index int) bool {
	return (heap.items[index] >= heap.leftChild(index) &&
			heap.items[index] >= heap.rightChild(index));
}

func (heap Heap) leftChild(index int) int {
	return heap.items[leftChildIndex(index)];
}

func (heap Heap) rightChild(index int) int {
	return heap.items[rightChildIndex(index)];
}

func leftChildIndex(index int) int {
	return (2 * index) + 1;
}

func rightChildIndex(index int) int {
	return (2 * index) + 2;
}

func getParent(index int) int {
	return int((index - 1) / 2);
}

func (heap *Heap) bubbleUp() {
	var index = heap.size - 1;
	for index > 0 && heap.items[index] > heap.items[getParent(index)] {
		heap.swap(index, getParent(index));
		index = getParent(index);
	}
}

func (heap *Heap) swap(firstIndex int, secondIndex int) {
	var temp = heap.items[firstIndex];
	heap.items[firstIndex] = heap.items[secondIndex];
	heap.items[secondIndex] = temp;
}

// Heapify an array
func Heapify(input []int) []int {
	for key := range input {
		heapify(input, key);
	}
	return input;
}

func heapify(input []int, index int) []int {
	var largerIndex = index;

	var leftChildindex = leftChildIndex(index);
	var rightChildindex = rightChildIndex(index);
	if leftChildindex < len(input) && input[largerIndex] < input[leftChildindex] {
		largerIndex = leftChildindex;
	} else if rightChildindex < len(input) && input[largerIndex] < input[rightChildindex] {
		largerIndex = rightChildindex;
	}

	if index == largerIndex {
		return input;
	}

	input = swap(input, index, largerIndex);
	return heapify(input, largerIndex);
}

func swap(input []int, first int, second int) []int {
	var temp = input[first];
	input[first] = input[second];
	input[second] = temp;
	return input;
}