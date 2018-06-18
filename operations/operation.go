package operations

import "kmipserver/kmip"

var operation = map[int]Operation{}

func GetOperation(i int) Operation {
	return operation[i]
}

type Operation interface {
	Do(b kmip.Batch) *kmip.BatchItem
}
