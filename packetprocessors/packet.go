package packetprocessors


type Processor interface {
	ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error)
}
