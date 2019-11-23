package packet

import (
	"../message"
)

type Packet struct {

	header PacketHeader;
	body PacketBody
}

type PacketHeader struct {

	operation string;
}

type PacketBody struct {

	packetContent interface{};
	message message.Message;
}