package packetprocessors

import "fmt"

type PrivateKeyTemplateAttribute struct {}



func init() {
	Kmpiprocessor["420065"] = new(PrivateKeyTemplateAttribute)
}


func (r * PrivateKeyTemplateAttribute) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("PrivateKeyTemplateAttribute",string(t.Tag), string(t.Type) , t.Length, t.Value)

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
