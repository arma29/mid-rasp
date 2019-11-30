package queue

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/message"
)

type Queue struct {
	MsgList []message.Message
}

type QueueManager struct {
	queueMap map[string]*Queue
}

func (manager *QueueManager) GetQueue(queueName string) *Queue {

	// Allocate for HashMap
	if manager.queueMap == nil {
		manager.queueMap = make(map[string]*Queue)
	}

	queue, exists := manager.queueMap[queueName]

	if !exists {
		queue = &Queue{MsgList: make([]message.Message, 0)}
		manager.queueMap[queueName] = queue
	}

	return queue
}

func (manager *QueueManager) EnqueueMsg(msg message.Message) {

	dest := msg.Header.Destination

	queue := manager.GetQueue(dest)
	queue.MsgList = append(queue.MsgList, msg)
}

func (manager *QueueManager) DequeueMsg(queueName string) message.Message {

	queue := manager.GetQueue(queueName)
	msg := queue.MsgList[0]
	queue.MsgList[0] = message.Message{}
	queue.MsgList = queue.MsgList[1:]

	return msg
}

func (manager *QueueManager) IsEmpty(queueName string) bool {
	queue := manager.GetQueue(queueName)
	return len(queue.MsgList) == 0
}

// Clean Old Messages
// func (manager *QueueManager) CleanMessagesFromQueue(Queue) {

// 	for _, msg := range manager.
// }
