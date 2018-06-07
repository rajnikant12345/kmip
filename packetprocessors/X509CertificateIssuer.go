package packetprocessors

import "fmt"

type X509CertificateIssuer struct {}



func init() {
	Kmpiprocessor["4200B6"] = new(X509CertificateIssuer)
}


func (r * X509CertificateIssuer) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("X509CertificateIssuer",string(t.Tag), string(t.Type) , t.Length, t.Value)

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
