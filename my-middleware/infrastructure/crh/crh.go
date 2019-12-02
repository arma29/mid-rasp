package crh

import (
	"encoding/binary"
	"net"
	"strconv"

	"github.com/arma29/mid-rasp/shared"
)

type CRH struct {
	ServerHost string
	ServerPort int
}

func (crh CRH) Send(msg []byte) error {

	var connSend net.Conn
	var errSend error

	for {
		connSend, errSend = net.Dial("tcp", crh.ServerHost+":"+strconv.Itoa(crh.ServerPort))

		if errSend == nil && connSend != nil {
			break
		}
	}

	defer connSend.Close()

	// Send message to Server
	msgLengthBytes := make([]byte, 4)
	length := uint32(len(msg))

	binary.LittleEndian.PutUint32(msgLengthBytes, length)
	_, errSend = connSend.Write(msgLengthBytes)

	if errSend != nil {
		return errSend
	}

	_, errSend = connSend.Write(msg)

	return errSend
}

var listener net.Listener
var connReceive net.Conn
var errReceive error

func (crh CRH) Receive() []byte {

	listener, errReceive = net.Listen("tcp", crh.ServerHost+":"+strconv.Itoa(crh.ServerPort))
	shared.CheckError(errReceive)

	defer listener.Close()

	for {
		connReceive, errReceive = listener.Accept()
		if errReceive == nil && connReceive != nil {
			break
		}
	}

	defer connReceive.Close()

	// Receive Message
	pktLengthBytes := make([]byte, 4)
	_, errReceive = connReceive.Read(pktLengthBytes)
	shared.CheckError(errReceive)

	pktLength := binary.LittleEndian.Uint32(pktLengthBytes)

	// receive message
	pkt := make([]byte, pktLength)
	_, errReceive = connReceive.Read(pkt)
	shared.CheckError(errReceive)

	return pkt
}
