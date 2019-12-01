package message

import ()

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
	ExpirationDate int64;
}


type MessageBody struct {
	Content interface{};
}