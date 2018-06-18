package kmip


type Processor interface {
	ProcessPacket(ctx *Message, t *TTLV, req []byte) error
}
