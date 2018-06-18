package packetprocessors

import (
	"errors"
	"fmt"
	"kmipserver/kmip"
	"kmipserver/server"
)

type ActivationDate struct{}

func init() {
	server.Kmpiprocessor[4325377] = new(ActivationDate)
}

func (r *ActivationDate) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("ActivationDate", t.Type, t.Length)

	if (len(req)) <= 0 {
		return errors.New("Cannot parse")
	}

	f, s := kmip.ReadTTLV(req)
	p := server.GetProcessor(s.Tag)

	if p != nil {
		ctx.BatchList[len(ctx.BatchList)-1].Attr.Activationdate = kmip.StringToInt64(string(t.Value))
		p.ProcessPacket(ctx, &s, req[f:])
	}
	return errors.New("Not supported tag")
}
