package packetprocessors

import (
	"errors"
	"fmt"

	"kmipserver/kmip"
	"kmipserver/server"
)

type Qlength struct{}

func init() {
	server.Kmpiprocessor[4325491] = new(Qlength)
}

func (r *Qlength) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("Qlength", t.Type, t.Length)

	if (len(req)) <= 0 {
		return errors.New("Cannot parse")
	}

	f, s := kmip.ReadTTLV(req)
	p := server.GetProcessor(s.Tag)

	if p != nil {
		ctx.BatchList[len(ctx.BatchList)-1].Attr.CryptoDomainParams.Qlen = kmip.StringToInt(string(t.Value))
		p.ProcessPacket(ctx, &s, req[f:])
	}
	return errors.New("Not supported tag")
}
