package queue

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/message"
	"github.com/arma29/mid-rasp/my-middleware/distribution/packet"
	"github.com/arma29/mid-rasp/my-middleware/infrastructure/crh"
	"github.com/arma29/mid-rasp/my-middleware/distribution/marshaller"
	"github.com/arma29/mid-rasp/shared"
	"fmt"
)

type QueueManagerProxy struct {
	Host string
	Port int
	QueueName string
}


func (proxy QueueManagerProxy) Send(op string, content interface{}) {

	crh := crh.CRH{ServerHost: shared.QUEUE_SERVER_HOST, ServerPort: shared.QUEUE_SERVER_PORT }
	marshaller := marshaller.Marshaller{}

	msgHeader := message.MessageHeader{Host: proxy.Host, Port: proxy.Port, Destination: proxy.QueueName}
	msgBody := message.MessageBody{Content: content}
	msg := message.Message{Header: msgHeader, Body: msgBody}

	pkt := packet.Packet{}
	pkt.Header = packet.PacketHeader{Operation: op}
	pkt.Body = packet.PacketBody{Message: msg}

	crh.Send(marshaller.Marshal(pkt))

	//TODO: Check if subscription was successful

	// Listen for notifications
	if (op == "subscribe") {
		proxy.Receive()
	}
	
}




func (proxy QueueManagerProxy) Receive() message.Message {

	crh := crh.CRH{ServerHost: proxy.Host, ServerPort: proxy.Port }
	marshaller := marshaller.Marshaller{}

	pktBytes := crh.Receive()

	pkt := marshaller.Unmarshal(pktBytes)
	msg := pkt.Body.Message

	fmt.Printf("Mensagem Recebida:\n")
	fmt.Printf("\t%v", msg)

	return msg
}






