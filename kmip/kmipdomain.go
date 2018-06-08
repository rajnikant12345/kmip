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

func IntToString(in int) string {
	return fmt.Sprintf("%08X",in)
}


func ReadTTLV(b []byte) (int, TTLV) {
	k := TTLV{}

	t,_ := strconv.ParseInt(string(b[:blocklen-2]),16,32)
	k.Tag = int(t)
	t,_ = strconv.ParseInt(string(b[blocklen-2: blocklen]),16,32)
	k.Type = int(t)


	lstr := b[blocklen : blocklen*2]

	k.Length = StringToInt(string(lstr))

	if len(b[blocklen*2:]) < k.Length*2 {
		return -1, TTLV{}
	}
	switch k.Type {
	case 1:
		k.Length = k.Length*2
		return blocklen*2, k
	case 4, 7, 8:
		l := 16 * (k.Length / 8)
		if k.Length%8 != 0 {
			l += 16
		}
		k.Value = b[blocklen*2 : blocklen*2+l]
		return blocklen*2+l, k
	default:
		k.Value = b[blocklen*2 : blocklen*4]
		return blocklen*4, k
	}

	return -1, TTLV{}
}

