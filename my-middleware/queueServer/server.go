package main

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/invoker"
	"github.com/arma29/mid-rasp/shared"
	"fmt"
)

func main() {

	fmt.Println("Queue Server running...")
	
	// Listen for publish/subscribe requests
	go func() {
		queueInvoker := invoker.QueueInvoker{Host:shared.QUEUE_SERVER_HOST, Port:shared.QUEUE_SERVER_PORT}
		queueInvoker.Invoke()
	}()

	// Maintain process running
	wait := make(chan string)
	<-wait
	
}