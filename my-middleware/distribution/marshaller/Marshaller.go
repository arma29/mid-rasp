package marshaller

import (
	"encoding/json"
	"github.com/arma29/mid-rasp/shared"
	"github.com/arma29/mid-rasp/my-middleware/distribution/packet"
)

type Marshaller struct {}

func (Marshaller) Marshal(pkt packet.Packet) []byte {

	result, err := json.Marshal(pkt)
	shared.CheckError(err)

	return result
}

func (Marshaller) Unmarshal(pkt []byte) packet.Packet {
		
	result := packet.Packet{}
	
	err := json.Unmarshal(pkt, &result)
	shared.CheckError(err)
	
	return result
}