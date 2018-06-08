package kmip

import "context"

var Kmpiprocessor  = map[int]Processor{}

func GetProcessor(packettype int) Processor {
	p,ok := Kmpiprocessor[packettype]
	if !ok {
		return nil
	}
	return p
}



func ProessPacket(b []byte) ([]byte,error) {
	f,s := ReadTTLV(b)
	p := GetProcessor(s.Tag)
	if p!= nil {
		return p.ProcessPacket(context.Background(), &s  , b[f:], nil, nil)
	}
	return nil, nil
}
