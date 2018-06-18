package packetprocessors

import (
	"errors"
	"fmt"
	
	"kmipserver/kmip"
	"kmipserver/server"
)

type X509CertificateSubject struct{}

func init() {
	server.Kmpiprocessor[4325559] = new(X509CertificateSubject)
}

func (r *X509CertificateSubject) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("X509CertificateSubject", t.Type, t.Length)

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
