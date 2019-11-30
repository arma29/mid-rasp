package main

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
)

func main() {

	// Operator Info
	OPERATOR_HOST := "localhost"
	OPERATOR_PORT := 9004
	RADIATION_QUEUE := "radiation"
	// ALERT_QUEUE := "alert"

	// Object responsable for delievering message to queue
	radQueueProxy := queue.QueueManagerProxy{Host: OPERATOR_HOST, Port: OPERATOR_PORT, QueueName: RADIATION_QUEUE}
	radQueueProxy.Send("subscribe", nil)

	// Alert Queue
	// alertQueueProxy := queue.QueueManagerProxy{Host: OPERATOR_HOST, Port: OPERATOR_PORT, QueueName: ALERT_QUEUE}
	// alertQueueProxy.

	// Stop main thread
	messages := make(chan string)
	<-messages
}
