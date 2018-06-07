package packetprocessors

import "fmt"

type AttributeName struct {}



func init() {
	Kmpiprocessor["42000A"] = new(AttributeName)
}


func (r * AttributeName) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("AttributeName",string(t.Tag), string(t.Type) , t.Length, string(t.Value))

	f,s := ReadTTLV(req)
	p := GetProcessor(string(s.Tag))

	if(len(req[f:])) <= 0 {
		return nil, nil
	}

	if p!= nil {
		p.ProcessPacket(&s,req[f:], nil, nil)
	}
	return nil,nil
}
