package packetprocessors

import "fmt"

type SignatureKeyInformation struct {}



func init() {
	Kmpiprocessor["42004E"] = new(SignatureKeyInformation)
}


func (r * SignatureKeyInformation) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("SignatureKeyInformation",string(t.Tag), string(t.Type) , t.Length, t.Value)

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
