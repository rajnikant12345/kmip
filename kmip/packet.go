package kmip

import "context"

type Processor interface {
	ProcessPacket(ctx context.Context , t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error)
}
