package packetprocessors

import (
	"errors"
	"fmt"

	"kmipserver/kmip"
	"kmipserver/server"
)

type RandomIV struct{}

func init() {
	server.Kmpiprocessor[4325573] = new(RandomIV)
}

func (r *RandomIV) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("RandomIV", t.Type, t.Length)

	if (len(req)) <= 0 {
		return errors.New("Cannot parse")
	}

	f, s := kmip.ReadTTLV(req)
	p := server.GetProcessor(s.Tag)

	if p != nil {
		ctx.BatchList[len(ctx.BatchList)-1].Attr.CryptoParams.RandomIv = kmip.StringToInt64(string(t.Value))
		p.ProcessPacket(ctx, &s, req[f:])
	}
	return errors.New("Not supported tag")
}
