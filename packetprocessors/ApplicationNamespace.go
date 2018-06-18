package packetprocessors

import (
	"errors"
	"fmt"
	
	"kmipserver/kmip"
	"kmipserver/server"
)

type ApplicationNamespace struct{}

func init() {
	server.Kmpiprocessor[4325379] = new(ApplicationNamespace)
}

func (r *ApplicationNamespace) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("ApplicationNamespace", t.Type, t.Length)

	if (len(req)) <= 0 {
		return errors.New("Cannot parse")
	}

	f, s := kmip.ReadTTLV(req)
	p := server.GetProcessor(s.Tag)

	if p != nil {
		ctx.BatchList[len(ctx.BatchList)-1].Attr.ApplicationSpecificInformation.ApplicationNamespace = kmip.BinToString(t.Value)[:t.Length]
		p.ProcessPacket(ctx, &s, req[f:])
	}
	return errors.New("Not supported tag")
}
