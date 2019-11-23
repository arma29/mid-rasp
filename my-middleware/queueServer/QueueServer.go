package queueServer

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
	"fmt"
)

type QueueServer struct {

	queueManager queue.QueueManager
	pubManager PublisherManager
	subMananger SubscriberManager
}


func main() {

	Server := QueueServer{}
	fmt.Println(Server)

	
}