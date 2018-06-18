package operations

import (
	"kmipserver/kmip"
	"client"
)

type Create struct {
}

func init() {
	operation[1] = new(Create)
}

func (c *Create) Do(b kmip.Batch) *kmip.BatchItem {
	k := kmip.BatchItem{}
	k.Operation = b.Operation
	k.ResultStatus = 0
	k.ResponsePayload.ObjectType = b.Attr.ObjectType
	ss := client.CreateOp()
	k.ResponsePayload.UniqueIdentifier = ss
	return &k
}
