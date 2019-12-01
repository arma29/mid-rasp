package main

import (
	"fmt"
	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
	"github.com/mitchellh/mapstructure"
	// "github.com/arma29/mid-rasp/my-middleware/distribution/message"
	rad "github.com/arma29/mid-rasp/radiation"
)

func main() {

	// Operator Info
	OPERATOR_HOST := "localhost"
	OPERATOR_PORT := 9004
	RADIATION_QUEUE := "radiation"
	// ALERT_QUEUE := "alert"

	// Object responsable for delievering message to queue
	radQueueProxy := queue.QueueManagerProxy{Host: OPERATOR_HOST, Port: OPERATOR_PORT, QueueName: RADIATION_QUEUE}
	radChannel := radQueueProxy.Subscribe()
	
	go GetRadiation(radChannel)

	// Alert Queue
	// alertQueueProxy := queue.QueueManagerProxy{Host: OPERATOR_HOST, Port: OPERATOR_PORT, QueueName: ALERT_QUEUE}
	// alertQueueProxy.

	// Stop main thread
	messages := make(chan string)
	<-messages
}

func GetRadiation(ch chan interface{}) {

	for res := range ch {
		// msg := res.(message.Message)
		// rad := (msg.Body.Content).(rad.Radiation)
		radiation := rad.Radiation{}
		mapstructure.Decode(res, &radiation)
		fmt.Printf("valor: ")
		fmt.Println(radiation.Value)
	}

}