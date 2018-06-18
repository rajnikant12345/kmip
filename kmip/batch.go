package kmip

import "kmipserver/attributes"

type VendorExtension struct {
	Type int
	Value interface{}
}

type MessageExtension struct {
	VendorIdentification string
	CriticalityIndicator int64
	VendorExtension VendorExtension
}


type Batch struct {
	Batchid   string `xml:",omitempty"`
	Operation int `xml:",omitempty"`
	AttrList	[]string
	Attr      attributes.Attribute `xml:",omitempty"`
	MessageExtension MessageExtension
}

type Message struct {
	AttrName          string `xml:",omitempty"`
	ProtoVersionMajor int `xml:",omitempty"`
	ProtoVersionMinor int
	BatchCount        int `xml:",omitempty"`
	BatchOrderOperation int64 `xml:",omitempty"`
	MaxResponSize     int `xml:",omitempty"`
	BatchList         []Batch `xml:",omitempty"`
}
