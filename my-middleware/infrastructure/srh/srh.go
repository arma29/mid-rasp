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


func (srh SRH) Send(msg []byte) error {

	var conn net.Conn
	var err error

	for {
		conn, err = net.Dial("tcp", srh.ServerHost + ":" + strconv.Itoa(srh.ServerPort))

		if err == nil && conn != nil {
			break
		}
	}

	defer conn.Close()

	// Send Message
	msgLengthBytes := make([]byte, 4)
	msgLength := uint32(len(msg))


	binary.LittleEndian.PutUint32(msgLengthBytes, msgLength)
	_ , err = conn.Write(msgLengthBytes)
	if err != nil {
		return err
	}
		
	_, err = conn.Write(msg)
	if err != nil {
		return err
	}

	return err
}

