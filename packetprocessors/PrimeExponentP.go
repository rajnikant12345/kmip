package packetprocessors

import "fmt"

type PrimeExponentP struct {}



func init() {
	Kmpiprocessor["420060"] = new(PrimeExponentP)
}


func (r * PrimeExponentP) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("PrimeExponentP",string(t.Tag), string(t.Type) , t.Length, t.Value)

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
