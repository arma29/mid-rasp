package main

import (
	"fmt"
	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
	"github.com/mitchellh/mapstructure"
	rad "github.com/arma29/mid-rasp/radiation"
	"time"
)

func main() {

	// Operator Info
	OPERATOR_HOST := "localhost"
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
	<- wait
}

func GetRadiation(radChannel chan interface{}, dangerRadChannel chan bool) {

	for res := range radChannel {
		// Get RadiationStruct
		radiation := rad.Radiation{}
		mapstructure.Decode(res, &radiation)

		value := radiation.Value
		timestamp := radiation.Timestamp

		// Check if Radiation is Dangerous
		dangerRadChannel <- rad.IsRadiationDangerous(value)

		timeSTR := time.Unix(0, timestamp)
		fmt.Printf("\n\nValor: %f\nTimeStamp:%s", value, timeSTR)
	}

}


func SentAlert(alertQueueProxy queue.QueueManagerProxy, dangerRadChannel chan bool) {

	for res := range dangerRadChannel {
		validator := rad.Validator{IsDangerous:res, Timestamp: time.Now().UnixNano()}
		alertQueueProxy.Send("publish", validator)

		fmt.Printf("\nIsDangerous: ")
		fmt.Println(res)
	}
}