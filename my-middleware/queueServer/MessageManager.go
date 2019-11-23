package queueServer

import (
	"github.com/arma29/mid-rpc/my-middleware/distribution/message"
	"github.com/arma29/mid-rpc/my-middleware/distribution/queue"
)

type MessageManager interface{

	queueManager queue.QueueManagerProxy

}

func (msgManager MessageManager) 