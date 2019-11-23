package queueServer

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
)

type QueueServer struct {

	queueManager queue.queueManager
	pubManager queue.PublisherMananger
	subMananger queue.SubscriberMananger
}


func main() {

	Server := QueueServer{}

	
}