package queue

import (
	"../packet"
)

type Queue struct {

	length int
	items []packet.Packet
}
