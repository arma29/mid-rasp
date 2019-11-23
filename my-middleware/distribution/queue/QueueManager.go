package queue

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/message"
)

type QueueManager struct {
	host string
	port int
	queueMap map[string]Queue
}


func (manager QueueManager) GetQueue(queueName string) Queue {

	// Allocate for HashMap
	if (manager.queueMap == nil) {
		manager.queueMap = make(map[string]Queue)
	}

	queue, exists := manager.queueMap[queueName]

	if (!exists) {
		queue = Queue{}
		manager.queueMap[queueName] = queue
	}

	return queue
}


func (manager QueueManager) QueueMsg(msg message.Message) {

	msgHeader := msg.Header

	queue := manager.queueMap[msgHeader.Destination]
	queue.msgList.Append(msg)
	queue.length = len(queue.msgList)
}


