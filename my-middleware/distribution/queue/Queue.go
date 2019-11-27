package queue

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/message"
)

type Queue struct {
	Length int
	MsgList []message.Message
}
