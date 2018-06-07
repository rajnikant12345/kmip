package packetprocessors

var Kmpiprocessor  = map[string]Processor{}

func GetProcessor(packettype string) Processor {
	p,ok := Kmpiprocessor[packettype]
	if !ok {
		return nil
	}
	return p
}



func ProessPacket(b []byte) ([]byte,error) {
	f,s := ReadTTLV(b)
	p := GetProcessor(string(s.Tag))
	if p!= nil {
		return p.ProcessPacket( &s  , b[f:], nil, nil)
	}
	return nil, nil
}
