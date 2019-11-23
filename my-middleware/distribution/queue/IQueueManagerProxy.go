package queue

import (
	"../message"
)

type IQueueManager interface {

	send(msg message.Message)
	receive() message.Message
}