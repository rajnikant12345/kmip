
package packetprocessors

import (
	"fmt"
	"kmipserver/kmip"
	"errors"
	"context"
)

type AttributeValue struct {}



func init() {
	kmip.Kmpiprocessor[4325387] = new(AttributeValue)
}


func (r * AttributeValue) ProcessPacket(ctx context.Context, t *kmip.TTLV, req []byte, response []byte , processor kmip.Processor) ([]byte,error) {

	fmt.Println("AttributeValue",t.Tag, t.Type , t.Length, t.Value)

	if(len(req)) <= 0 {
		return nil,errors.New("Incomplete Packet")
	}

	f,s := kmip.ReadTTLV(req)
	p := kmip.GetProcessor(s.Tag)

	if p!= nil {
		p.ProcessPacket(ctx , &s,req[f:], nil, nil)
	}
	return nil,errors.New("Invalid Packet")
}

