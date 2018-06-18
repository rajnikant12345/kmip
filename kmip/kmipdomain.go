package kmip

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

type TTLV struct {
	Tag    int
	Type   int
	Length int
	Value  []byte
}

const blocklen = 8

func StringToInt(lstr string) int {
	l, e := hex.DecodeString(string(lstr))
	if e != nil {
		fmt.Println(e.Error())
		return -1
	}
	return int(l[3]&0x000000ff) | int((l[2]&0x000000ff))<<8 | (int(l[1]) & 0x000000ff << 16) | (int(l[0]) & 0x000000ff << 24)
}

func StringToInt64(lstr string) int64 {
	l, e := hex.DecodeString(string(lstr))
	if e != nil {
		fmt.Println(e.Error())
		return -1
	}
	return int64(l[7]&0x00000000000000ff) | int64((l[6]&0x00000000000000ff))<<8 | (int64(l[5]) & 0x00000000000000ff << 16) |
		(int64(l[4]) & 0x00000000000000ff << 24) | (int64(l[3]) & 0x00000000000000ff << 32) | (int64(l[2]) & 0x00000000000000ff << 40) |
		(int64(l[1]) & 0x00000000000000ff << 48) | (int64(l[0]) & 0x00000000000000ff << 54)
}

func BinToString(v []byte) string {
	l, e := hex.DecodeString(string(v))
	if e != nil {
		fmt.Println(e.Error())
		return ""
	}
	return string(l)
}


func ReadTTLV(b []byte) (int, TTLV) {
	k := TTLV{}

	t, _ := strconv.ParseInt(string(b[:blocklen-2]), 16, 32)
	k.Tag = int(t)
	t, _ = strconv.ParseInt(string(b[blocklen-2:blocklen]), 16, 32)
	k.Type = int(t)

	lstr := b[blocklen : blocklen*2]

	k.Length = StringToInt(string(lstr))

	if len(b[blocklen*2:]) < k.Length*2 {
		return -1, TTLV{}
	}
	switch k.Type {
	case 1:
		k.Length = k.Length * 2
		return blocklen * 2, k
	case 4, 7, 8:
		l := 16 * (k.Length / 8)
		if k.Length%8 != 0 {
			l += 16
		}
		k.Value = b[blocklen*2 : blocklen*2+l]
		return blocklen*2 + l, k
	default:
		k.Value = b[blocklen*2 : blocklen*4]
		return blocklen * 4, k
	}

	return -1, TTLV{}
}


func WriteStringToByte(s string)string {
	l := hex.EncodeToString([]byte(s))
	return l
}

func WriteByteStringToByte(s string)string {
	return s
}

func IntToString(in int) string {
	t := int64(in << 32)
	return fmt.Sprintf("%016X", t)
}

func Int64ToString(in int64) string {
	return fmt.Sprintf("%016X", in)
}

func WriteTag(tag int) string {
	return fmt.Sprintf("%06X", tag)
}

func WriteType(Type int) string {
	return fmt.Sprintf("%02X", Type)
}

func WriteLength(length int) string {
	return fmt.Sprintf("%08X", length)
}


func WriteTTLV(t TTLV, val interface{})string {

	if val == nil {
		return ""
	}
	s1 := ""


	switch t.Type {
	case 1:
		s1 += val.(string)
	case 2:
		s1 += IntToString(val.(int))
	case 3:
		s1 += Int64ToString(val.(int64))
	//case 4:
	case 5:
		s1 += IntToString(val.(int))
	case 6:
		s1 += Int64ToString(val.(int64))
	case 7:
		s1 += WriteStringToByte(val.(string))
		t.Length = len(s1)
	case 8:
		s1 += WriteByteStringToByte(val.(string))
		t.Length = len(s1)
	case 9:
		s1 += Int64ToString(val.(int64))
	case 10:
		s1 += IntToString(val.(int))
	}

	s := WriteTag(t.Tag)
	s += WriteType(t.Type)
	s += WriteLength(t.Length)

	return s+s1
}

