package packetprocessors

import (
	"errors"
	"fmt"
	"kmipserver/kmip"
	"kmipserver/server"
)

type ApplicationData struct{}

func init() {
	server.Kmpiprocessor[4325378] = new(ApplicationData)
}

func (r *ApplicationData) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("ApplicationData", t.Type, t.Length)

	if (len(req)) <= 0 {
		return errors.New("Cannot parse")
	}

	f, s := kmip.ReadTTLV(req)
	p := server.GetProcessor(s.Tag)

	if p != nil {
		ctx.BatchList[len(ctx.BatchList)-1].Attr.ApplicationSpecificInformation.ApplicationData = kmip.BinToString(t.Value)[:t.Length]
		p.ProcessPacket(ctx, &s, req[f:])
	}
	return errors.New("Not supported tag")
}
