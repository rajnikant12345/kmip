package packetprocessors

import (
	"errors"
	"fmt"

	"kmipserver/kmip"
	"kmipserver/server"
	"strings"
	"kmipserver/attributes"
)

type AttributeValue struct{}

func init() {
	server.Kmpiprocessor[4325387] = new(AttributeValue)
}

func (r *AttributeValue) ProcessPacket(ctx *kmip.Message, t *kmip.TTLV, req []byte) error {

	fmt.Println("AttributeValue", t.Type, t.Length)

	if (len(req)) <= 0 {
		return errors.New("Cannot parse")
	}

	f, s := kmip.ReadTTLV(req)
	p := server.GetProcessor(s.Tag)

	if p != nil {
		if strings.HasPrefix(strings.ToLower(ctx.AttrName) , "x-") {
			c :=attributes.CustomAttributeValue{t.Type , t.Value}
			ctx.BatchList[len(ctx.BatchList)-1].Attr.CustomAttribute[ctx.AttrName] = c
		} else if t.Type != 1 {
			kmip.SetAttribue(ctx.AttrName, t.Value[:t.Length*2], ctx)
		}
		p.ProcessPacket(ctx, &s, req[f:])
	}
	return errors.New("Not supported tag")
}
