package server

import (
	"net"

	"kmipserver/kmip"
	"fmt"
	"kmipserver/operations"
	"errors"
)

var Kmpiprocessor = map[int]kmip.Processor{}

func GetProcessor(packettype int) kmip.Processor {
	p, ok := Kmpiprocessor[packettype]
	if !ok {
		return nil
	}
	return p
}

func ValidateRequestMessage(s kmip.TTLV) error {
	if s.Tag != kmip.RequestMessageTag {
		return errors.New("Invalid tag")
	}
	return nil
}

func ProcessPacket(c net.Conn, b []byte) error {
	f, s := kmip.ReadTTLV(b)

	e := ValidateRequestMessage(s)
	if e != nil {
		return e
	}

	p := GetProcessor(s.Tag)
	if p != nil {
		ctx := new(kmip.Message)
		p.ProcessPacket(ctx, &s, b[f:])
		//if e != nil {
		//	return e
		//}
		r := kmip.ResponseMessage{}
		r.ResponseHeader.P.ProtoVersionMinor = ctx.ProtoVersionMinor
		r.ResponseHeader.P.ProtoVersionMajor = ctx.ProtoVersionMajor
		r.ResponseHeader.BatchCount = ctx.BatchCount

		fmt.Println(r.ResponseHeader.BatchCount)

		for _,i := range ctx.BatchList {
			op := operations.GetOperation(i.Operation)
			item := op.Do(i)
			r.BatchResponse.BatchItemList = append(r.BatchResponse.BatchItemList, *item)
		}

		s := r.CustomMarshalKMIP()
		if len(s) > ctx.MaxResponSize {
			// handle this
		}
		c.Write([]byte(s))
	}
	return nil
}
