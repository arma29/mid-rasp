package queueServer

import (
	"github.com/arma29/mid-rpc/my-middleware/distribution/queue"
	"github.com/arma29/mid-rpc/my-middleware/distribution/message"
)

type QueueServer interface {}

func CreateQueue(name string) {
	return queue.Queue{length: 0, items: message.Message[]}
}



func main() {

	
}