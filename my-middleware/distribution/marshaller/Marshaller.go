package marshaller

import (
	"encoding/json"
	"../packet"
)

type Marshaller struct {}

func (Marshaller) Marshal(pct packet.Packet) []byte {

	result, err := json.Marshal(pct)
	shared.CheckError(err)

	return result
}

func (Marshaller) Unmarshal(pct []byte) pct.Packet {
		
	err := json.Unmarshal(pct, &result)
	shared.CheckError(err)
	
	return result
}