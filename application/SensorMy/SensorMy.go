package main

import(
	// "fmt"
	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
	// rad "github.com/arma29/mid-rasp/radiation"
	
)

func main() {

	SENSOR_HOST := "localhost"
	SENSOR_PORT := 9004

	// Object responsable for delievering message to queue
	radQueueProxy := queue.QueueManagerProxy{Host: SENSOR_HOST, Port: SENSOR_PORT, QueueName: "radiation2"}
	radQueueProxy.Send("subscribe", nil)

}
