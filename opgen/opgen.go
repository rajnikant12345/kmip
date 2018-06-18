package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func GenerateTagFile(tagname string, tagvalue int) {
	f, err := os.Create(tagname + ".go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	str := `
package operations

import (
	"context"
	"kmipserver/kmip"
)

type %s struct {

}

func init() {
	operation[%d] = new(%s)
}


func (c * %s) Do(ctx context.Context , t *kmip.TTLV, req []byte, response []byte) ([]byte,error) {
	fmt.Println("operation : %s")
	return nil, nil
}

`
	fmt.Fprintf(f, str, tagname, tagvalue, tagname, tagname, tagname)
}

var TagMap map[string]string

func ParseTag(b []byte) {
	s := strings.Split(string(b), ",")
	TagMap[s[1]] = s[0]
}

func LoadTagsType(tagfile string) {
	TagMap = make(map[string]string)
	f, _ := os.Open(tagfile)
	b := bufio.NewReader(f)
	d, _, err := b.ReadLine()
	for err != io.EOF {
		ParseTag(d)
		d, _, err = b.ReadLine()
	}
}

func main() {
	//LoadTagsType(os.Args[1])
	//for i,j := range TagMap {

	for {
		var t, d string
		fmt.Scanf("%s %s\n", &t, &d)
		fmt.Println(t, d)
		if t == "---END---" {
			break
		}
		i, _ := strconv.ParseInt(d, 16, 32)
		GenerateTagFile(t, int(i))
	}

	//	i,_ := strconv.ParseInt(os.Args[2],16,32)

	//GenerateTagFile(os.Args[1],int(i))
	//}

}
