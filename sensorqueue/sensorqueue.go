package sensorqueue

const MAX_QUEUE_SIZE = 100

var queue []float32

// InitQueue is
func InitQueue() {
	queue = make([]float32, 0)
}

// Empty is
func Empty() bool {
	return len(queue) == 0
}

// Enqueue is
func Enqueue(e float32) {
	if len(queue) >= MAX_QUEUE_SIZE {
		return
	}
	queue = append(queue, e)
}

// Dequeue is
func Dequeue() {
	if Empty() {
		return
	}
	queue[0] = 0
	queue = queue[1:]
}

// Peek is
func Peek() interface{} {
	if Empty() {
		return nil
	}
	return queue[0]
}

// Size is
func Size() int {
	return len(queue)
}
