package queue

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/message"
)

type Queue struct {
	Length int
	MsgList []message.Message
}

type QueueManager struct { 
	queueMap map[string]*Queue
}


func (manager *QueueManager) GetQueue(queueName string) *Queue {

	// Allocate for HashMap
	if (manager.queueMap == nil) {
		manager.queueMap = make(map[string]*Queue)
	}

	queue, exists := manager.queueMap[queueName]

	if (!exists) {
		queue = &Queue{Length:0, MsgList: make([]message.Message, 0)}
		manager.queueMap[queueName] = queue
	}

	return queue
}


func (manager *QueueManager) EnqueueMsg(msg message.Message) {

	dest := msg.Header.Destination

	queue := manager.GetQueue(dest)

	queue.MsgList = append(queue.MsgList, msg)
	queue.Length = len(queue.MsgList)

}


