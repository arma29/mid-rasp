package message

import (
	"time"
)

type Message struct {
	Header MessageHeader;
	Body MessageBody;
}


type MessageHeader struct {
	Host string;
	Port int;
	Destination string;
	Enterprise string;
	SensorID string;
	Priority int;
	ExpirationDate time.Duration;
}


type MessageBody struct {
	Content interface{};
}