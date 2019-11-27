package invoker

import (
	"fmt"
	"github.com/arma29/mid-rasp/my-middleware/infrastructure/srh"
	"github.com/arma29/mid-rasp/my-middleware/distribution/marshaller"
	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
)

type QueueInvoker struct {
	Host string
	Port int
}

func (invoker QueueInvoker) Invoke() {
	srhImpl := srh.SRH{ServerHost:invoker.Host, ServerPort:invoker.Port}
	marshallerImpl := marshaller.Marshaller{}
	queueServer := queue.QueueServer{}

	// control loop
	for {
		// receive request packet
		rcvMsgBytes := srhImpl.Receive()

		// unmarshall request packet
		packetReceived := marshallerImpl.Unmarshal(rcvMsgBytes)

		// extract operation name
		operation := packetReceived.Header.Operation

		// demux request
		switch operation {
			case "subscribe" :
				
				// Get Message Data
				msgReceived := packetReceived.Body.Message
				host := msgReceived.Header.Host
				port := msgReceived.Header.Port
				dest := msgReceived.Header.Destination

				// Subscribe to Queue
				subManager := &queueServer.SubManager
				subManager.SubscribeRequest(host, port, dest)
				
				// Logging Info
				fmt.Printf("Invoker Op -> Subscribe -> %s:%d subscribed to \"%s\"\n", host, port, dest)
				fmt.Printf("Invoker Op -> Subscribe -> Lista de Subscribers (%d)\n", len(subManager.SubList))
				for _, subscriber := range subManager.SubList {
					fmt.Printf("%s:%d\n", subscriber.Host, subscriber.Port)

					for k, v := range subscriber.QueuesSubscribed {
						if v {
							fmt.Printf("\t%s\n", k)
						}
					}
				}

			case "Lookup":
				fmt.Printf("Lookup not implemented")
			}

	}
}