package crh

import (
	"net"
	"strconv"
	"encoding/binary"
	"github.com/arma29/mid-rasp/shared"
)

type CRH struct {
	ServerHost string
	ServerPort int
}

func (crh CRH) Send(msg []byte) {

	var conn net.Conn
	var err error

	for {
		conn, err = net.Dial("tcp", crh.ServerHost + ":" + strconv.Itoa(crh.ServerPort))

		if err == nil && conn != nil {
			break
		}
	}

	// Send message to Server
	msgLengthBytes := make([]byte, 4)
	length := uint32(len(msg))

	binary.LittleEndian.PutUint32(msgLengthBytes, length)
	_, err = conn.Write(msgLengthBytes)
	shared.CheckError(err)
	
	_, err = conn.Write(msg)
	shared.CheckError(err)

	conn.Close()
}

var listener net.Listener
var conn net.Conn
var err error


func (crh CRH) Receive() []byte {


	listener, err = net.Listen("tcp", crh.ServerHost + ":" + strconv.Itoa(crh.ServerPort))
	shared.CheckError(err)

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