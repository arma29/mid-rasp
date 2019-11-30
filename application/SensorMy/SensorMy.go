package main

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
	rad "github.com/arma29/mid-rasp/radiation"
)

func main() {

	// Sensor Info
	SENSOR_HOST := "localhost"
	SENSOR_PORT := 9015
	RADIATION_QUEUE := "radiation"
	// ALERT_QUEUE := "alert"

	// Object responsable for delievering message to queue
	radQueueProxy := queue.QueueManagerProxy{Host: SENSOR_HOST, Port: SENSOR_PORT, QueueName: RADIATION_QUEUE}
	radQueueProxy.Send("publishRequest", nil)

	// // Alert Queue
	// alertQueueProxy := queue.QueueManagerProxy{Host: SENSOR_HOST, Port: SENSOR_PORT, QueueName: ALERT_QUEUE}
	// alertQueueProxy.Send("subscribe", nil)

	// Published Data
	radQueueProxy.Send("publish", rad.Radiation{Value: 5})
	radQueueProxy.Send("publish", rad.Radiation{Value: 9})
}
