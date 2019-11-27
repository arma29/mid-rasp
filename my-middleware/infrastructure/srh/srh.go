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

	// Remenber to close connection
	defer listener.Close()

	for {
		conn, err = listener.Accept()
		if err == nil {
			break
		}
	}

	// Receive Message
	pktLengthBytes := make([]byte, 4)
	_, err = conn.Read(pktLengthBytes)
	shared.CheckError(err)

	pktLength := binary.LittleEndian.Uint32(pktLengthBytes)

	// receive message
	pkt := make([]byte, pktLength)
	_, err = conn.Read(pkt)
	shared.CheckError(err)

	return pkt
	
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

