package packetprocessors

import "fmt"

type X_509CertificateSubject struct {}



func init() {
	Kmpiprocessor["4200B7"] = new(X_509CertificateSubject)
}


func (r * X_509CertificateSubject) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("X.509CertificateSubject",string(t.Tag), string(t.Type) , t.Length, t.Value)

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
