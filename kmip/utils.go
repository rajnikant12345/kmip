package kmip

const (
	EDES        = 0x00000001
	EDES3       = 0x00000002
	EAES        = 0x00000003
	ERSA        = 0x00000004
	EDSA        = 0x00000005
	EECDSA      = 0x00000006
	EHMACSHA1   = 0x00000007
	EHMACSHA224 = 0x00000008
	EHMACSHA256 = 0x00000009
	EHMACSHA384 = 0x0000000A
	EHMACSHA512 = 0x0000000B
	EHMACMD5    = 0x0000000C
	EDH         = 0x0000000D
	EECDH       = 0x0000000E
	EECMQV      = 0x0000000F
	EBlowfish   = 0x00000010
	ECamellia   = 0x00000011
	ECAST5      = 0x00000012
	EIDEA       = 0x00000013
	EMARS       = 0x00000014
	ERC2        = 0x00000015
	ERC4        = 0x00000016
	ERC5        = 0x00000017
	ESKIPJACK   = 0x00000018
	ETwofish    = 0x00000019

	EEC = 0x0000001A

	EOneTimePad = 0x0000001B
	EARIA       = 0x80000001
	EFPE        = 0x80000002
	ESEED       = 0x80000003
)

func GetAlgoFromEnum(enu int) string {
	switch enu {
	case EDES:
		return "DES"
	// case EDES3:         return "DES3";
	case EDES3:
		return "TDES"
	case EAES:
		return "AES"
	case ERSA:
		return "RSA"
	case EDSA:
		return "DSA"
	case EECDSA:
		return "ECDSA"
	case EHMACSHA1:
		return "HMAC-SHA1"
	case EHMACSHA224:
		return "HMACSHA224"
	case EHMACSHA256:
		return "HMAC-SHA256"
	case EHMACSHA384:
		return "HMAC-SHA384"
	case EHMACSHA512:
		return "HMAC-SHA512"
	case EHMACMD5:
		return "HMACMD5"
	case EDH:
		return "DH"
	case EECDH:
		return "ECDH"
	case EECMQV:
		return "ECMQV"
	case EBlowfish:
		return "Blowfish"
	case ECamellia:
		return "Camellia"
	case ECAST5:
		return "CAST5"
	case EIDEA:
		return "IDEA"
	case EMARS:
		return "MARS"
	case ERC2:
		return "RC2"
	case ERC4:
		return "RC4"
	case ERC5:
		return "RC5"
	case ESKIPJACK:
		return "SKIPJACK"
	case ETwofish:
		return "Twofish"
	case EEC:
		return "EC"
	case EOneTimePad:
		return "OneTimePad"
	case EARIA:
		return "ARIA"
	case EFPE:
		return "FPE"
	case ESEED:
		return "SEED"
	}
	return ""
}

func GetEnumFromString() int {

}
