package marshaller

import (
	"encoding/json"
	"github.com/arma29/mid-rasp/shared"
	"github.com/arma29/mid-rasp/my-middleware/distribution/packet"
)

type Marshaller struct {}

func (Marshaller) Marshall(pct packet.Packet) []byte {

	result, err := json.Marshal(pct)
	shared.CheckError(err)

	return result
}

func (Marshaller) Unmarshall(pct []byte) packet.Packet {
		
	result := packet.Packet{}
	
	err := json.Unmarshal(pct, &result)
	shared.CheckError(err)
	
	return result
}