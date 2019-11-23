package packet

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/message"
)

type Packet struct {

	Header PacketHeader;
	Body PacketBody
}

type PacketHeader struct {

	Operation string;
}

type PacketBody struct {

	Message message.Message;
}