
package packetprocessors

import (
	"fmt"
	"kmipserver/kmip"
	"errors"
	"context"
)

type Template struct {}



func init() {
	kmip.Kmpiprocessor[4325520] = new(Template)
}


func (r * Template) ProcessPacket(ctx context.Context , t *kmip.TTLV, req []byte, response []byte , processor kmip.Processor) ([]byte,error) {

	fmt.Println("Template",t.Tag, t.Type , t.Length, t.Value)

	if(len(req)) <= 0 {
		return nil,errors.New("Incomplete Packet")
	}

	f,s := kmip.ReadTTLV(req)
	p := kmip.GetProcessor(s.Tag)

	if(len(req[f:])) <= 0 {
		return nil,errors.New("Incomplete Packet")
	}
	if p!= nil {
		p.ProcessPacket(ctx, &s,req[f:], nil, nil)
	}
	return nil,errors.New("Invalid Packet")
}

