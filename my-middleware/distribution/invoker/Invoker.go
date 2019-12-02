package invoker

import (
	"fmt"

	"github.com/arma29/mid-rasp/my-middleware/distribution/marshaller"
	"github.com/arma29/mid-rasp/my-middleware/distribution/queue"
	"github.com/arma29/mid-rasp/my-middleware/infrastructure/srh"
)

type QueueInvoker struct {
	Host string
	Port int
}

func (invoker QueueInvoker) Invoke() {

	srhImpl := srh.SRH{ServerHost: invoker.Host, ServerPort: invoker.Port}
	marshallerImpl := marshaller.Marshaller{}

	queueServer := queue.QueueServer{}

	// Queue Server Managers
	subManager := &queueServer.SubManager
	pubManager := &queueServer.PubManager
	queueManager := &queueServer.QueueManager
	notifManager := &queueServer.NotifManager
	eventNotifManager := &queueServer.EventNotifManager

	// Execute Nofitication Event Sender
	notifChannel := make(chan queue.Notification, 1000)
	eventNotifManager.SetChannel(notifChannel)
	go eventNotifManager.SendNotifications()

	// control loop
	for {
		// receive request packet
		rcvPktBytes, err := srhImpl.Receive()

		for err != nil {
			rcvPktBytes, err = srhImpl.Receive()
		}

		// unmarshall request packet
		packetReceived := marshallerImpl.Unmarshal(rcvPktBytes)
		msgReceived := packetReceived.Body.Message

		host := msgReceived.Header.Host
		port := msgReceived.Header.Port
		dest := msgReceived.Header.Destination

		// extract operation name
		operation := packetReceived.Header.Operation

		// demux request
		switch operation {
		case "subscribe":
			// Subscribe to Queue
			sub := subManager.SubscribeRequest(host, port, dest)

			// Create Notifications for New Subscriber
			queue := queueManager.GetQueue(dest)
			notifList := notifManager.NewSubscriberToNotify(*sub, *queue)

			for _, notif := range notifList {
				notifChannel <- notif
			}

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

		case "publishRequest":
			// Publish to Queue
			pubManager.PublishRequest(host, port, dest)

			// // Logging Info
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

		case "publish":
			//Check if Publisher is authorized
			publisher, exists := pubManager.GetPublisher(host, port)
			if !exists {
				break
			}

			if publisher.PublishQueues == nil {
				break
			}

			if !publisher.PublishQueues[dest] {
				break
			}

			// Publish to Queue
			queueManager.EnqueueMsg(msgReceived)

			// Create Notification to Subscribers
			if queueManager.IsEmpty(dest) {
				break
			}

			// Logging Info
			// fmt.Printf("\n")
			// fmt.Printf("Invoker Op -> Publish -> %s:%d has published in \"%s\"\n", host, port, dest)
			// queue := queueManager.GetQueue(dest)
			// fmt.Printf("Invoker Op -> Publish -> Lista de Mensagens na Fila \"%s\" (%d)\n", dest, len(queue.MsgList))
			// for _, msg := range queue.MsgList {
			// 	fmt.Printf("%s:%d -> %v\n", msg.Header.Host, msg.Header.Port, msg.Body.Content)
			// }

			subList := subManager.GetSubscribersByQueue(dest)

			notifList := notifManager.NewMessageToNotify(msgReceived, subList)

			for _, notif := range notifList {
				notifChannel <- notif
			}

		}
	}
}
