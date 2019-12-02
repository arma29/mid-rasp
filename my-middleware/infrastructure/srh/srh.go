package srh

import (
	"encoding/binary"
	"net"
	"strconv"

	"github.com/arma29/mid-rasp/shared"
)

type SRH struct {
	ServerHost string
	ServerPort int
}

var listener net.Listener
var connReceive net.Conn
var errReceive error

func (srh SRH) Receive() ([]byte, error) {

	listener, errReceive = net.Listen("tcp", srh.ServerHost+":"+strconv.Itoa(srh.ServerPort))
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
	// shared.CheckError(errReceive)
	if errReceive != nil {
		return nil, errReceive
	}

	pktLength := binary.LittleEndian.Uint32(pktLengthBytes)

	// receive message
	pkt := make([]byte, pktLength)
	_, errReceive = connReceive.Read(pkt)
	if errReceive != nil {
		return nil, errReceive
	}
	// shared.CheckError(errReceive)

	return pkt, errReceive
}

func (srh SRH) Send(msg []byte) error {

	var connSend net.Conn
	var errSend error

	for {
		connSend, errSend = net.Dial("tcp", srh.ServerHost+":"+strconv.Itoa(srh.ServerPort))

		if errSend == nil && connSend != nil {
			break
		}
	}

	defer connSend.Close()

	// Send Message
	msgLengthBytes := make([]byte, 4)
	msgLength := uint32(len(msg))

	binary.LittleEndian.PutUint32(msgLengthBytes, msgLength)
	_, errSend = connSend.Write(msgLengthBytes)
	if errSend != nil {
		return errSend
	}

	_, errSend = connSend.Write(msg)
	if errSend != nil {
		return errSend
	}

	return errSend
}
