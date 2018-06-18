package packetprocessors

import (
	"errors"
	"fmt"

	"kmipserver/kmip"
	"kmipserver/server"
)

type SubjectDistinguishedName struct{}

func init() {
	server.Kmpiprocessor[4325556] = new(SubjectDistinguishedName)
}

func (r *SubjectDistinguishedName) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("SubjectDistinguishedName", t.Type, t.Length)

	if (len(req)) <= 0 {
		return errors.New("Cannot parse")
	}

	f, s := kmip.ReadTTLV(req)
	p := server.GetProcessor(s.Tag)

	if p != nil {
		ctx.BatchList[len(ctx.BatchList)-1].Attr.X509CertificateSub.SubjectDistinguishedName = t.Value[:t.Length*2]
		p.ProcessPacket(ctx, &s, req[f:])
	}
	return errors.New("Not supported tag")
}
