package packetprocessors

import (
	"errors"
	"fmt"

	"kmipserver/kmip"
	"kmipserver/server"
)

type Data struct{}

func init() {
	server.Kmpiprocessor[4325570] = new(Data)
}

func (r *Data) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("Data", t.Type, t.Length)

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
