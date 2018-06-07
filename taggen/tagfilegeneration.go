package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
)

func GenerateTagFile(tagname string, tagvalue string) {
	f,err := os.Create(tagname+".go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Fprintf(f,"package packetprocessors\n\n")
	fmt.Fprintf(f,"import \"fmt\"\n\n")

	fmt.Fprintf(f,"type %s struct {}\n\n",tagname)
	str := `

func init() {
	Kmpiprocessor["%s"] = new(%s)
}


func (r * %s) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("%s",string(t.Tag), string(t.Type) , t.Length, t.Value)

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
`
	fmt.Fprintf(f,str,tagvalue, tagname,tagname,tagname)
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
	LoadTagsType(os.Args[1])
	for i,j := range TagMap {
		GenerateTagFile(j,i)
	}

}
