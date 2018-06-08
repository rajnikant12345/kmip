package packetprocessors

import (
	"fmt"
	"kmipserver/kmip"
	"context"
	"errors"
)

type RequestMessage struct {}

func init() {
	kmip.Kmpiprocessor[4325496] = new(RequestMessage)
}

func (r * RequestMessage) ProcessPacket(ctx context.Context , t *kmip.TTLV, req []byte, response []byte , processor kmip.Processor) ([]byte,error) {

	fmt.Println("RequestMessage",t.Tag, t.Type , t.Length, t.Value)

	if(len(req)) <= 0 {
		return nil,errors.New("Incomplete Packet")
	}

	f,s := kmip.ReadTTLV(req)
	p := kmip.GetProcessor(s.Tag)

	if p!= nil {
		p.ProcessPacket(ctx, &s,req[f:], nil, nil)
	}
	return nil,nil
}
