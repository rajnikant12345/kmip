package packetprocessors

import (
	"errors"
	"fmt"
	
	"kmipserver/kmip"
	"kmipserver/server"
)

type HashingAlgorithm struct{}

func init() {
	server.Kmpiprocessor[4325432] = new(HashingAlgorithm)
}

func (r *HashingAlgorithm) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("HashingAlgorithm", t.Type, t.Length)

	if (len(req)) <= 0 {
		return errors.New("Cannot parse")
	}

	f, s := kmip.ReadTTLV(req)
	p := server.GetProcessor(s.Tag)

	if p != nil {
		if ctx.AttrName == "Cryptographic Parameters" {
			ctx.BatchList[len(ctx.BatchList)-1].Attr.CryptoParams.HashASlgo = kmip.StringToInt(string(t.Value))
		}

		if ctx.AttrName == "Digest" {
			ctx.BatchList[len(ctx.BatchList)-1].Attr.Digest.HashAlgo = kmip.StringToInt(string(t.Value))
		}


		p.ProcessPacket(ctx, &s, req[f:])
	}
	return errors.New("Not supported tag")
}
