package main

import(
	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
)

func main() {
	
	// Operator Info
	OPERATOR_HOST := "localhost"
	OPERATOR_PORT := 9004

	// Object responsable for delievering message to queue
	radQueueProxy := queue.QueueManagerProxy{Host: OPERATOR_HOST, Port: OPERATOR_PORT, QueueName: "radiation"}
	radQueueProxy.Send("subscribe", nil)

	// Stop main thread
	messages := make(chan string)
	<-messages
}
