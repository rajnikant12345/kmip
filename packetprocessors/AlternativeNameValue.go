package packetprocessors

import "fmt"

type AlternativeNameValue struct {}



func init() {
	Kmpiprocessor["4200C0"] = new(AlternativeNameValue)
}


func (r * AlternativeNameValue) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("AlternativeNameValue",string(t.Tag), string(t.Type) , t.Length, t.Value)

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
