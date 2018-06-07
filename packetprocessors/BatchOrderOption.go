package packetprocessors

import "fmt"

type BatchOrderOption struct {}



func init() {
	Kmpiprocessor["420010"] = new(BatchOrderOption)
}


func (r * BatchOrderOption) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("BatchOrderOption",string(t.Tag), string(t.Type) , t.Length, t.Value)

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
