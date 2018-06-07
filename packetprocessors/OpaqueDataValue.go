package packetprocessors

import "fmt"

type OpaqueDataValue struct {}



func init() {
	Kmpiprocessor["42005A"] = new(OpaqueDataValue)
}


func (r * OpaqueDataValue) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("OpaqueDataValue",string(t.Tag), string(t.Type) , t.Length, t.Value)

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
