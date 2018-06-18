package packetprocessors

import (
	"errors"
	"fmt"

	"kmipserver/kmip"
	"kmipserver/server"
)

type UniqueBatchItemID struct{}

func init() {
	server.Kmpiprocessor[4325523] = new(UniqueBatchItemID)
}

func (r *UniqueBatchItemID) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("UniqueBatchItemID", t.Type, t.Length, string(t.Value))

	if (len(req)) <= 0 {
		return errors.New("Cannot parse")
	}

	f, s := kmip.ReadTTLV(req)
	p := server.GetProcessor(s.Tag)

	if p != nil {
		ctx.BatchList[len(ctx.BatchList)-1].Batchid = string(t.Value)
		p.ProcessPacket(ctx, &s, req[f:])
	}
	return errors.New("Not supported tag")
}
