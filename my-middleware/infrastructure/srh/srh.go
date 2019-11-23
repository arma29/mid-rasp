package srh

import (
	"net"
	"strconv"
	"encoding/binary"
	"github.com/arma29/mid-rasp/shared"
)

type SRH struct {
	ServerHost string
	ServerPort int
}

var listener net.Listener
var conn net.Conn
var err error

func (srh SRH) Receive() []byte {

	listener, err = net.Listen("tcp", srh.ServerHost + ":" + strconv.Itoa(srh.ServerPort))
	shared.CheckError(err)

	for {
		conn, err = listener.Accept()
		if err == nil {
			break
		}
	}

	// Receive Message
	msgLengthBytes := make([]byte, 4)
	_, err = conn.Read(msgLengthBytes)
	shared.CheckError(err)

	msgLength := binary.LittleEndian.Uint32(msgLengthBytes)

	// receive message
	msg := make([]byte, msgLength)
	_, err = conn.Read(msg)
	shared.CheckError(err)

	return msg
}


func (SRH) Send(msg []byte) {

	// Send Message
	msgLengthBytes := make([]byte, 4)
	msgLength := uint32(len(msg))

	binary.LittleEndian.PutUint32(msgLengthBytes, msgLength)
	_ , err = conn.Write(msgLengthBytes)
	shared.CheckError(err)

	_, err = conn.Write(msg)
	shared.CheckError(err)

	conn.Close()
	listener.Close()
}
