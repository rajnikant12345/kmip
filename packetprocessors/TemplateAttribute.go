package packetprocessors

import "fmt"

type TemplateAttribute struct {}



func init() {
	Kmpiprocessor["420091"] = new(TemplateAttribute)
}


func (r * TemplateAttribute) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("TemplateAttribute",string(t.Tag), string(t.Type) , t.Length, t.Value)

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
