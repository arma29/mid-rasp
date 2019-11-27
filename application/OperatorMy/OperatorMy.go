package main

import(
	// "fmt"
	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
	// rad "github.com/arma29/mid-rasp/radiation"
	
)

func main() {

	OPERATOR_HOST := "localhost"
	OPERATOR_PORT := 9004

	// Object responsable for delievering message to queue
	radQueueProxy := queue.QueueManagerProxy{Host: OPERATOR_HOST, Port: OPERATOR_PORT, QueueName: "radiation"}
	radQueueProxy.Send("subscribe", nil)

	// fmt.Scanln()
}
