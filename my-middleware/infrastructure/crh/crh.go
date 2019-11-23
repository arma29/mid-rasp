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

	// Receiver Message
	// msgReceivedLengthBytes := make([]byte, 4)
	// _, err = conn.Read(msgReceivedLengthBytes)
	
	// msgReceivedLengthInt := binary.LittleEndian.Uint32(msgReceivedLengthBytes)

	// msgFromServer := make([]byte, msgReceivedLengthInt)
	// _, err = conn.Read(msgFromServer)
	// shared.CheckError(err)

	conn.Close()

	// return msgFromServer
}