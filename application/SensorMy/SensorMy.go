package main

import(
	// "fmt"
	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
	rad "github.com/arma29/mid-rasp/radiation"
	
)

func main() {

	SENSOR_HOST := "localhost"
	SENSOR_PORT := 9015

	// Object responsable for delievering message to queue
	radQueueProxy := queue.QueueManagerProxy{Host: SENSOR_HOST, Port: SENSOR_PORT, QueueName: "radiation"}
	radQueueProxy.Send("publishRequest", nil)


	radQueueProxy.Send("publish", rad.Radiation{Value:5})

}
