package kmip



func checkWriteInt(t TTLV , val interface{} )  string {

	var v int
	var v1 int64
	v,ok := val.(int)
	if  !ok {
		v1,ok = val.(int64)
		if  !ok {
			return ""
		}
	}
	if v1==0 && v ==0  {
		return ""
	}
	return WriteTTLV(TTLV{t.Tag , t.Type, t.Length , nil}, val)

}

func checkWriteString(t TTLV , val string)  string {
	if len(val) == 0 {
		return ""
	}
	//length/2 is necessary for structures
	return WriteTTLV(TTLV{t.Tag , t.Type, t.Length/2 , nil}, val)
}

//===========================================================================

type ProtoCol struct {
	ProtoVersionMajor int `xml:",omitempty"`
	ProtoVersionMinor int
}


func (r *ProtoCol)CustomMarshalKMIP() string {
	//marshal batch list
	s1 := checkWriteInt(TTLV{ProtocolVersionMajor , 2, 4 , nil}, r.ProtoVersionMajor)
	s1 += WriteTTLV(TTLV{ProtocolVersionMinor , 2, 4 , nil}, r.ProtoVersionMinor)
	s := checkWriteString(TTLV{ProtocolVersion , 1, len(s1) , nil}, s1)
	return s
}
//===========================================================================
type ResponseHeader struct {
	P ProtoCol
	TimeStamp int64
	BatchCount int
}

func (r *ResponseHeader)CustomMarshalKMIP() string {
	s1 := r.P.CustomMarshalKMIP()
	s1 += checkWriteInt(TTLV{TimeStampTag , 9,8  , nil}, r.TimeStamp)
	s1 += checkWriteInt(TTLV{BatchCountTag , 2, 4 , nil}, r.BatchCount)
	s :=  checkWriteString(TTLV{ResponseHeaderTag , 1, len(s1) , nil}, s1)
	return s
}
//================================================================================



//======================================================================================

type ResponseMessage struct {
	ResponseHeader ResponseHeader
	BatchResponse BatchResponse
}


func (r *ResponseMessage)CustomMarshalKMIP() string {
	s1 := r.ResponseHeader.CustomMarshalKMIP()
	s1 += r.BatchResponse.CustomMarshalKMIP()
	s :=  checkWriteString(TTLV{ResponseMessageTag , 1, len(s1), nil}, s1)
	return s

}
