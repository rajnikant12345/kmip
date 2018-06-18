package packetprocessors

import (
	"errors"
	"fmt"
	
	"kmipserver/kmip"
	"kmipserver/server"
)

type LastChangeDate struct{}

func init() {
	server.Kmpiprocessor[4325448] = new(LastChangeDate)
}

func (r *LastChangeDate) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("LastChangeDate", t.Type, t.Length)

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
