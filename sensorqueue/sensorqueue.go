package sensorqueue

import (
	rad "github.com/arma29/mid-rasp/radiation"
)

// maxQueueSize represents the max length of queue
const maxQueueSize = 100

var queue []rad.Radiation

// InitQueue is
func InitQueue() {
	queue = make([]rad.Radiation, 0)
}

// Empty is
func Empty() bool {
	return len(queue) == 0
}

// Enqueue is
func Enqueue(e rad.Radiation) {
	if len(queue) >= maxQueueSize {
		return
	}
	queue = append(queue, e)
}

// Dequeue is
func Dequeue() {
	if Empty() {
		return
	}
	queue[0] = rad.Radiation{}
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
