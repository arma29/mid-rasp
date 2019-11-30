package invoker

import (
	"fmt"
	"github.com/arma29/mid-rasp/my-middleware/infrastructure/srh"
	"github.com/arma29/mid-rasp/my-middleware/distribution/marshaller"
	"github.com/arma29/mid-rasp/my-middleware/distribution/packet"
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

	// Queue Server Managers
	subManager := &queueServer.SubManager
	pubManager := &queueServer.PubManager
	queueManager := &queueServer.QueueManager
	notifManager := &queueServer.NotifManager


	// Execute Nofitication Event Sender
	go func () {
		for {
			for _, notif := range notifManager.NotificationList {
				subSRH := srh.SRH{ServerHost: notif.Host, ServerPort: notif.Port}

				pktHeader := packet.PacketHeader{Operation: "notify"}
				pktBody := packet.PacketBody{Message: notif.Message}
				pkt := packet.Packet{Header: pktHeader, Body: pktBody}

				pktBytes := marshallerImpl.Marshal(pkt)


				subSRH.Send(pktBytes)

				notif.Status = queue.NOTICATION_SENT

			}
		}
	}()

	// control loop
	for {
		// receive request packet
		rcvMsgBytes := srhImpl.Receive()

		// unmarshall request packet
		packetReceived := marshallerImpl.Unmarshal(rcvMsgBytes)
		msgReceived := packetReceived.Body.Message

		host := msgReceived.Header.Host
		port := msgReceived.Header.Port
		dest := msgReceived.Header.Destination

		// extract operation name
		operation := packetReceived.Header.Operation

		// demux request
		switch operation {
			case "subscribe" :
				// Subscribe to Queue
				sub := subManager.SubscribeRequest(host, port, dest)

				// Create Notifications for New Subscriber
				queue := queueManager.GetQueue(dest)
				notifManager.NewSubscriberToNotify(*sub, *queue)
				
				// Logging Info
				fmt.Printf("\n")
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

			case "publishRequest" :
				// Publish to Queue
				pubManager.PublishRequest(host, port, dest)
				
				// Logging Info
				fmt.Printf("\n")
				fmt.Printf("Invoker Op -> Publish Request -> %s:%d is now a publisher of \"%s\"\n", host, port, dest)
				fmt.Printf("Invoker Op -> Publish Request -> Lista de Publishers (%d)\n", len(pubManager.PubList))
				for _, publisher := range pubManager.PubList {
					fmt.Printf("%s:%d\n", publisher.Host, publisher.Port)

					for k, v := range publisher.PublishQueues {
						if v {
							fmt.Printf("\t%s\n", k)
						}
					}
				}

			case "publish" :
				// TODO: Check if Publisher is authorized

				// Publish to Queue
				queueManager.EnqueueMsg(msgReceived)

				// Create Notification to Subscribers
				subList := subManager.GetSubscribersByQueue(dest)
				notifManager.NewMessageToNotify(msgReceived, subList)

				// Logging Info
				fmt.Printf("\n")
				fmt.Printf("Invoker Op -> Publish -> %s:%d has published in \"%s\"\n", host, port, dest)
				queue := queueManager.GetQueue(dest)
				fmt.Printf("Invoker Op -> Publish -> Lista de Mensagens na Fila \"%s\" (%d)\n", dest, len(queue.MsgList))
				for _, msg := range queue.MsgList {
					fmt.Printf("%s:%d -> %v\n", msg.Header.Host, msg.Header.Port, msg.Body.Content)
				}
		}
	}
}
