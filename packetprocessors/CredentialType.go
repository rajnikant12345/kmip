package packetprocessors

import (
	"errors"
	"fmt"

	"kmipserver/kmip"
	"kmipserver/server"
)

type CredentialType struct{}

func init() {
	server.Kmpiprocessor[4325412] = new(CredentialType)
}

func (r *CredentialType) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("CredentialType", t.Type, t.Length)

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
