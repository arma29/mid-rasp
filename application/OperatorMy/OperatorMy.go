package main

import (
	"fmt"
	"os"
	"time"

	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
	rad "github.com/arma29/mid-rasp/radiation"
	"github.com/mitchellh/mapstructure"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Missing arguments")
		os.Exit(1)
	}
	// Operator Info
	OPERATOR_HOST := os.Args[1]
	OPERATOR_PORT := 9004
	RADIATION_QUEUE := "radiation"
	ALERT_QUEUE := "alert"

	// Channel for Dangerous Radiation monitoring
	dangerRadChannel := make(chan bool, 10)

	// Rad Queue
	radQueueProxy := queue.QueueManagerProxy{Host: OPERATOR_HOST, Port: OPERATOR_PORT, QueueName: RADIATION_QUEUE}
	radChannel := radQueueProxy.Subscribe()
	go GetRadiation(radChannel, dangerRadChannel)

	// Alert Queue
	alertQueueProxy := queue.QueueManagerProxy{Host: OPERATOR_HOST, Port: OPERATOR_PORT, QueueName: ALERT_QUEUE}
	alertQueueProxy.Send("publishRequest", nil)
	go SentAlert(alertQueueProxy, dangerRadChannel)

	// Stop main thread
	wait := make(chan string)
	<-wait
}

func GetRadiation(radChannel chan interface{}, dangerRadChannel chan bool) {

	fmt.Println("Time")
	for res := range radChannel {
		// Get RadiationStruct
		radiation := rad.Radiation{}
		mapstructure.Decode(res, &radiation)

		fmt.Printf("Estrutura Recebida: ")
		fmt.Print(radiation)

		value := radiation.Value

		// Medindo o tempo
		t1 := time.Now().UnixNano()
		t2 := radiation.Timestamp
		s := fmt.Sprintf("%d", t1-t2)
		fmt.Println(s)

		// Check if Radiation is Dangerous
		dangerRadChannel <- rad.IsRadiationDangerous(value)
	}

}

func SentAlert(alertQueueProxy queue.QueueManagerProxy, dangerRadChannel chan bool) {

	for res := range dangerRadChannel {
		validator := rad.Validator{IsDangerous: res, Timestamp: time.Now().UnixNano()}
		alertQueueProxy.Send("publish", validator)

		// fmt.Printf("\nIsDangerous: ")
		// fmt.Println(res)
	}
}
