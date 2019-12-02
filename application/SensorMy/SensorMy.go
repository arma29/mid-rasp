package main

import (
	"fmt"

	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
	rad "github.com/arma29/mid-rasp/radiation"
	"github.com/arma29/mid-rasp/shared"
	"github.com/mitchellh/mapstructure"

	"github.com/arma29/mid-rasp/sensorqueue"
	"github.com/stianeikeland/go-rpio"

	"time"
)

var (
	pin rpio.Pin
)

func main() {

	// Sensor Info
	SENSOR_HOST := "localhost"
	SENSOR_PORT := 9015
	RADIATION_QUEUE := "radiation"
	ALERT_QUEUE := "alert"

	// Prepara o GPIO
	rpioErr := rpio.Open()
	if rpioErr != nil {
		panic(fmt.Sprint("Unable to open gpio", rpioErr.Error()))
	}

	defer rpio.Close()

	pin = rpio.Pin(18)
	pin.Output()

	// Alert Subscribe Queue
	alertQueueProxy := queue.QueueManagerProxy{Host: SENSOR_HOST, Port: SENSOR_PORT, QueueName: ALERT_QUEUE}
	alertChannel := alertQueueProxy.Subscribe()
	go OnAlert(alertChannel)

	// Radiation Publish Queue
	radQueueProxy := queue.QueueManagerProxy{Host: SENSOR_HOST, Port: SENSOR_PORT, QueueName: RADIATION_QUEUE}
	radQueueProxy.Send("publishRequest", nil)
	PublishRadiation(radQueueProxy)

	// Keep process running
	// wait := make(chan int)
	// <-wait

}

func PublishRadiation(proxy queue.QueueManagerProxy) {

	sensorqueue.InitQueue()
	for i := 0; i < shared.SAMPLE_SIZE; i++ {
		// Radiation data
		value := rad.GenerateRadiationValue()
		timetamp := time.Now().UnixNano()
		radValue := rad.Radiation{Value: value, Timestamp: timetamp}

		sensorqueue.Enqueue(radValue)
		if hasConnection(proxy) {
			parseQueue(proxy)
		} else {
			fmt.Println("Sem conexão")
		}
		// Publish Radiation data
		// proxy.Send("publish", radValue)

		fmt.Printf("Estrutura Enviada:")
		fmt.Println(radValue)

		// Já deixa o LED apagado
		pin.Low()

		// Garantir taxa máxima
		time.Sleep(shared.REAL_TIME)
	}
}

func OnAlert(alertChannel chan interface{}) {

	for res := range alertChannel {
		validator := rad.Validator{}
		mapstructure.Decode(res, &validator)

		if validator.IsDangerous {
			fmt.Printf("Falha registrada em: ")
			fmt.Println(time.Unix(0, validator.Timestamp))

			// Acende o LED
			pin.High()
		}
	}
}

func hasConnection(proxy queue.QueueManagerProxy) bool {
	return len(proxy.Host) > 0
}

func parseQueue(proxy queue.QueueManagerProxy) {
	for !sensorqueue.Empty() {
		radValue := sensorqueue.Peek()
		// Publish Radiation data
		proxy.Send("publish", radValue)
		sensorqueue.Dequeue()
	}
}
