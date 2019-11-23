package message

import (
	"time"
)

type Message struct {

	header MessageHeader;
	body MessageBody;
}


type MessageHeader struct {

	destination string;
	subject string;
	author string;
	priority int;
	expirationDate time.Duration;
}


type MessageBody struct {

	content interface{};
}