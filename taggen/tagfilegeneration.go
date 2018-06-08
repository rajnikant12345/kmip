package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
	"strconv"
)

func GenerateTagFile(tagname string, tagvalue int) {
	f,err := os.Create(tagname+".go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	str := `
package packetprocessors

import (
	"fmt"
	"kmipserver/kmip"
	"errors"
)

type %s struct {}



func init() {
	kmip.Kmpiprocessor[%d] = new(%s)
}


func (r * %s) ProcessPacket(t *kmip.TTLV, req []byte, response []byte , processor kmip.Processor) ([]byte,error) {

	fmt.Println("%s",t.Tag, t.Type , t.Length, t.Value)

	f,s := kmip.ReadTTLV(req)
	p := kmip.GetProcessor(s.Tag)

	if(len(req[f:])) <= 0 {
		return nil,errors.New("Incomplete Packet")
	}
	if p!= nil {
		p.ProcessPacket(&s,req[f:], nil, nil)
	}
	return nil,errors.New("Invalid Packet")
}

`
fmt.Fprintf(f,str,tagname,tagvalue, tagname,tagname,tagname)
}

var TagMap map[string]string

func ParseTag(b []byte) {
	s := strings.Split(string(b),",")
	TagMap[s[1]] = s[0]
}

func LoadTagsType(tagfile string) {
	TagMap = make(map[string]string)
	f,_ := os.Open(tagfile)
	b := bufio.NewReader(f)
	d,_,err := b.ReadLine()
	for err != io.EOF{
		ParseTag(d)
		d,_,err = b.ReadLine()
	}
}


func main() {
	//LoadTagsType(os.Args[1])
	//for i,j := range TagMap {
	i,_ := strconv.ParseInt(os.Args[2],16,32)

	GenerateTagFile(os.Args[1],int(i))
	//}

}
