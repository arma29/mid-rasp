package queue

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/marshaller"
	"github.com/arma29/mid-rasp/my-middleware/distribution/packet"
	"github.com/arma29/mid-rasp/my-middleware/infrastructure/srh"
)

type EventNotification struct {
	NotifChannel chan Notification
}

func (manager *EventNotification) SetChannel(ch chan Notification) {
	manager.NotifChannel = ch
}

func (manager *EventNotification) SendNotifications() {

	marshallerImpl := marshaller.Marshaller{}

	for {
		notif := <-manager.NotifChannel

		subSRH := srh.SRH{ServerHost: notif.Host, ServerPort: notif.Port}

		// Create Packet
		pktHeader := packet.PacketHeader{Operation: "notify"}
		pktBody := packet.PacketBody{Message: notif.Message}
		pkt := packet.Packet{Header: pktHeader, Body: pktBody}

		pktBytes := marshallerImpl.Marshal(pkt)

		err := subSRH.Send(pktBytes)
		if err != nil {
			manager.NotifChannel <- notif
		}

	}
}
