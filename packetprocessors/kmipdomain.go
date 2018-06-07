package packetprocessors

import (
	"encoding/hex"
	"fmt"
)

type TTLV struct {
	Tag    []byte
	Type   []byte
	Length int
	Value  []byte
}

const TagLength = 6
const TypeLength = 2
const Len = 8

func ReadTTLV(b []byte) (int, TTLV) {
	k := TTLV{}

	k.Tag = b[:TagLength]
	k.Type = b[TagLength : TypeLength+TagLength]
	lstr := b[TypeLength+TagLength : TypeLength+TagLength+Len]

	l, e := hex.DecodeString(string(lstr))
	if e != nil {
		fmt.Println(e.Error())
	}
	k.Length = int(l[3]&0x000000ff) | int((l[2]&0x000000ff))<<8 | (int(l[1]) & 0x000000ff << 16) | (int(l[0]) & 0x000000ff << 24)

	if len(b[TypeLength+TagLength+Len:]) < k.Length*2 {
		panic("Bhag Yahan Se")
	}
	switch string(k.Type) {
	case "01":
		k.Length = k.Length*2
		return TypeLength + TagLength + Len, k
	case "04", "07", "08":
		l := 16 * (k.Length / 8)
		if k.Length%8 != 0 {
			l += 16
		}
		k.Value = b[TypeLength+TagLength+Len : TypeLength+TagLength+Len+l]
		return TypeLength + TagLength + Len + l, k
	default:
		k.Value = b[TypeLength+TagLength+Len : TypeLength+TagLength+Len+16]
		return TypeLength + TagLength + Len + 16, k
	}

	return 0, TTLV{}
}


func GetType(b []byte) (string , int) {
	return "",0
}
