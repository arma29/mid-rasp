package queue

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/message"
	"github.com/arma29/mid-rasp/my-middleware/distribution/packet"
	"github.com/arma29/mid-rasp/my-middleware/infrastructure/crh"
	"github.com/arma29/mid-rasp/my-middleware/distribution/marshaller"
	"github.com/arma29/mid-rasp/shared"
)

type QueueManagerProxy struct {
	queueName string
}


func (proxy QueueManagerProxy) Send(op string, msg message.Message) {

	crh := crh.CRH{ServerHost: shared.QUEUE_SERVER_HOST, ServerPort: shared.QUEUE_SERVER_PORT }
	marshaller := marshaller.Marshaller{}


	pkt := packet.Packet{}
	pkt.Header = packet.PacketHeader{Operation: op}
	pkt.Body = packet.PacketBody{Message: msg}

	crh.Send(marshaller.Marshall(pkt))
}