package packetprocessors

import (
	"errors"
	"fmt"
	"kmipserver/kmip"
	"kmipserver/server"
)

type AsynchronousCorrelationValue struct{}

func init() {
	server.Kmpiprocessor[4325382] = new(AsynchronousCorrelationValue)
}

func (r *AsynchronousCorrelationValue) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("AsynchronousCorrelationValue", t.Type, t.Length)

	if (len(req)) <= 0 {
		return errors.New("Cannot parse")
	}

	f, s := kmip.ReadTTLV(req)
	p := server.GetProcessor(s.Tag)

	if p != nil {
		p.ProcessPacket(ctx, &s, req[f:])
	}
	return errors.New("Not supported tag")
}
