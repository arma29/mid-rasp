package main

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
	rad "github.com/arma29/mid-rasp/radiation"
	"time"
	"fmt"
)

func main() {

	// Sensor Info
	SENSOR_HOST := "localhost"
	SENSOR_PORT := 9015
	RADIATION_QUEUE := "radiation"
	ALERT_QUEUE := "alert"

	// Object responsable for delievering message to queue
	radQueueProxy := queue.QueueManagerProxy{Host: SENSOR_HOST, Port: SENSOR_PORT, QueueName: RADIATION_QUEUE}
	radQueueProxy.Send("publishRequest", nil)

	// Published Data
	radQueueProxy.Send("publish", rad.Radiation{Value: 5, Timestamp: 0})
	// radQueueProxy.Send("publish", rad.Radiation{Value: 9})

	// Alert Queue
	alertQueueProxy := queue.QueueManagerProxy{Host: SENSOR_HOST, Port: SENSOR_PORT, QueueName: ALERT_QUEUE}
	go OnAlert(alertQueueProxy.Subscribe())

}



func sendRadiation(ch <-chan interface{} ) {
	
	fmt.Println("Alerta de Radiação Disparado!")
}


func OnAlert(ch <-chan interface{}) {
	for res := range ch {
		validator := res.(rad.Validator)
		if validator.IsDangerous {
			fmt.Printf("Falha registrada em: ")
			fmt.Println(time.Unix(0, validator.Timestamp))
		}
	}
}