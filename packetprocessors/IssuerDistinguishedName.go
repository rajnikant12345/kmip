package packetprocessors

import (
	"errors"
	"fmt"
	
	"kmipserver/kmip"
	"kmipserver/server"
)

type IssuerDistinguishedName struct{}

func init() {
	server.Kmpiprocessor[4325554] = new(IssuerDistinguishedName)
}

func (r *IssuerDistinguishedName) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("IssuerDistinguishedName", t.Type, t.Length)

	if (len(req)) <= 0 {
		return errors.New("Cannot parse")
	}

	f, s := kmip.ReadTTLV(req)
	p := server.GetProcessor(s.Tag)

	if p != nil {
		if ctx.AttrName == "X.509 Certificate Identifier" {
			ctx.BatchList[len(ctx.BatchList)-1].Attr.X509CertificateId.IssuerDistinguishedName = t.Value[:t.Length*2]
		}
		if ctx.AttrName == "X.509 Certificate Issuer" {
			ctx.BatchList[len(ctx.BatchList)-1].Attr.X509CertificateIssuer.IssuerDistinguishedName = t.Value[:t.Length*2]
		}

		p.ProcessPacket(ctx, &s, req[f:])
	}
	return errors.New("Not supported tag")
}
