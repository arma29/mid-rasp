package queue

import (
	"github.com/arma29/mid-rasp/my-middleware/distribution/message"
)

type Queue struct {
	length int
	msgList []message.Message
}
