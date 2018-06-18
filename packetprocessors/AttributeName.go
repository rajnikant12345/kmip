package packetprocessors

import (
	"errors"
	"fmt"

	"kmipserver/kmip"
	"kmipserver/server"
)

type AttributeName struct{}

func init() {
	server.Kmpiprocessor[4325386] = new(AttributeName)
}

func (r *AttributeName) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("AttributeName", t.Type, t.Length)

	if (len(req)) <= 0 {
		return errors.New("Cannot parse")
	}

	f, s := kmip.ReadTTLV(req)
	p := server.GetProcessor(s.Tag)

	if p != nil {
		ctx.AttrName = kmip.BinToString(t.Value)[:t.Length]
		ctx.BatchList[len(ctx.BatchList)-1].AttrList = append(ctx.BatchList[len(ctx.BatchList)-1].AttrList,ctx.AttrName)
		p.ProcessPacket(ctx, &s, req[f:])
	}
	return errors.New("Not supported tag")
}
