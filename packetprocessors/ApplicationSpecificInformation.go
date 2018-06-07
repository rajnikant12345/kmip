package packetprocessors

import "fmt"

type ApplicationSpecificInformation struct {}



func init() {
	Kmpiprocessor["420004"] = new(ApplicationSpecificInformation)
}


func (r * ApplicationSpecificInformation) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("ApplicationSpecificInformation",string(t.Tag), string(t.Type) , t.Length, t.Value)

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
