package requestor

import (
	"github.com/arma29/mid-rpc/my-middleware/distribution/packet"

)

type Requestor struct{}

func (requestor Requestor) Invoke(inv aux.Invocation) interface{} {
	
	marshallerInstance := marshaller.Marshaller{}
	crhInstance := crh.CRH{ServerHost: inv.Host, ServerPort: inv.Port}

	reqHeader := miop.RequestHeader{Operation: inv.Request.Op, ObjectID: inv.ObjectID}
	reqBody := miop.RequestBody{Body:inv.Request.Params}
	header := miop.Header{ByteOrder: true, Size: 4 }
	body := miop.Body{RequestHeader: reqHeader, RequestBody: reqBody}
	packetRequest := miop.Packet{Header: header, Body: body}



	msgRequestBytes := marshallerInstance.Marshal(packetRequest)
	msgResponseBytes := crhInstance.SendReceive(msgRequestBytes)

	result := msgResponsePacket.Body.ResponseBody.Body


	return result
}