package kmip


type AttrValue struct {
	Type int
	Value interface{}
}

func (r *AttrValue) CustomMarshalKMIP() string {
	return ""
}

type ResponseArrtibutes struct {
	Name string
	Value AttrValue
}

func (r *ResponseArrtibutes) CustomMarshalKMIP() string {
	return ""
}


type ResponsePayload struct {
	ObjectType int
	UniqueIdentifier string
	ResponseArrtibutesList []ResponseArrtibutes
}

func (r *ResponsePayload) CustomMarshalKMIP() string {
	s :=  checkWriteInt(TTLV{ObjectTypeTag , 5, 4 , nil}, r.ObjectType)
	s += checkWriteString(TTLV{UniqueIdentifierTag , 7, len(r.UniqueIdentifier) , nil}, r.UniqueIdentifier)
	for _,i := range r.ResponseArrtibutesList {
		s += i.CustomMarshalKMIP()
	}
	//add attributes
	s1 := checkWriteString(TTLV{ResponsePayloadTag , 1, len(s) , nil}, s)
	return s1
}


type BatchItem struct {
	Batchid   string //byte string
	Operation int
	ResultStatus int
	ResultReason int
	ResultMessage string
	ResponsePayload ResponsePayload
}

func (r *BatchItem) CustomMarshalKMIP() string {
	s := checkWriteString(TTLV{BatchIdTag , 8, len(r.Batchid) , nil}, r.Batchid)
	s +=  checkWriteInt(TTLV{OperationTag , 5, 4 , nil}, r.Operation)
	s +=  WriteTTLV(TTLV{ResultStatusTag , 5, 4 , nil}, r.ResultStatus)

	s +=  checkWriteInt(TTLV{ResultReasonTag , 2, 4 , nil}, r.ResultReason)

	s += checkWriteString(TTLV{ResultMessageTag , 7, len(r.ResultMessage) , nil}, r.ResultMessage)

	s +=  r.ResponsePayload.CustomMarshalKMIP()
	s1 := checkWriteString(TTLV{BatchItemTag , 1, len(s) , nil}, s)
	return s1
}


type BatchResponse struct {
	BatchItemList []BatchItem
}


func (r *BatchResponse) CustomMarshalKMIP() string {
	var s string
	for _,i := range r.BatchItemList {
		s += i.CustomMarshalKMIP()
	}
	return s
}