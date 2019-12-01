package main

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
	rad "github.com/arma29/mid-rasp/radiation"
	"github.com/mitchellh/mapstructure"
	// "github.com/stianeikeland/go-rpio"
	"time"
	"fmt"
)

func main() {

	// Sensor Info
	SENSOR_HOST := "localhost"
	SENSOR_PORT := 9015
	RADIATION_QUEUE := "radiation"
	ALERT_QUEUE := "alert"

	// Radiation Publish Queue
	radQueueProxy := queue.QueueManagerProxy{Host: SENSOR_HOST, Port: SENSOR_PORT, QueueName: RADIATION_QUEUE}
	radQueueProxy.Send("publishRequest", nil)
	go PublishRadiation(radQueueProxy)

	// Alert Subscribe Queue
	alertQueueProxy := queue.QueueManagerProxy{Host: SENSOR_HOST, Port: SENSOR_PORT, QueueName: ALERT_QUEUE}
	alertChannel := alertQueueProxy.Subscribe()
	go OnAlert(alertChannel)

	// Keep process running
	wait := make(chan int)
	<- wait

}



func PublishRadiation(proxy queue.QueueManagerProxy) {
	
	for {
		// Radiation data
		value := rad.GenerateRadiationValue()
		timetamp := time.Now().UnixNano()
		radValue := rad.Radiation{Value: value, Timestamp: timetamp}

		// Publish Radiation data
		proxy.Send("publish", radValue)

		fmt.Println("\nDado de radiação enviado:")
		fmt.Printf("\t%v\n", radValue)

		sleepDuration, _ := time.ParseDuration("1ms")
		time.Sleep(sleepDuration)
	}
}


func OnAlert(alertChannel chan interface{}) {

	// Prepara o GPIO
	// rpioErr := rpio.Open()
	// if rpioErr != nil {
		// panic(fmt.Sprint("Unable to open gpio", rpioErr.Error()))
	// }

	// defer rpio.Close()

	// pin := rpio.Pin(18)
	// pin.Output()

	for res := range alertChannel {
		validator := rad.Validator{}
		mapstructure.Decode(res, &validator)

		if validator.IsDangerous {
			fmt.Printf("Falha registrada em: ")
			fmt.Println(time.Unix(0, validator.Timestamp))
		}
		
	}
}