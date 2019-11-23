package message

import (
	"time"
)

type Message struct {

	Header MessageHeader;
	Body MessageBody;
}


type MessageHeader struct {

	Destination string;
	Enterprise string;
	SensorID string;
	Priority int;
	ExpirationDate time.Duration;
}


type MessageBody struct {

	content interface{};
}