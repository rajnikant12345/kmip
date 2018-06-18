package kmip


func SetAttribue(key string, val []byte, b *Message) {
	switch key {
	case "Unique Identifier":
		b.BatchList[len(b.BatchList)-1].Attr.Uid = BinToString(val)
	case "Object Type":
		b.BatchList[len(b.BatchList)-1].Attr.ObjectType = StringToInt(string(val))
	case "Cryptographic Algorithm":
		b.BatchList[len(b.BatchList)-1].Attr.CryptographicAlgo = StringToInt(string(val))
	case "Cryptographic Length":
		b.BatchList[len(b.BatchList)-1].Attr.CryptoGraphicLength = StringToInt(string(val))
	case "Certificate Type":
		b.BatchList[len(b.BatchList)-1].Attr.CertificateType = StringToInt(string(val))
	case "Certificate Length":
		b.BatchList[len(b.BatchList)-1].Attr.CertificacteLength = StringToInt(string(val))
	case "Digital Signature Algorithm":
		b.BatchList[len(b.BatchList)-1].Attr.DigitalSigAlgo = StringToInt(string(val))
	case "Operation Policy Name":
		b.BatchList[len(b.BatchList)-1].Attr.OperationPolicyName = BinToString(val)
	case "Cryptographic Usage Mask":
		b.BatchList[len(b.BatchList)-1].Attr.CryptoUsageMask = StringToInt(string(val))
	case "Lease Time":
		b.BatchList[len(b.BatchList)-1].Attr.LeaseTime = StringToInt(string(val))
	case "State":
		b.BatchList[len(b.BatchList)-1].Attr.State = StringToInt(string(val))
	case "Initial Date":
		b.BatchList[len(b.BatchList)-1].Attr.Initialdate = StringToInt64(string(val))
	case "Activation Date":
		b.BatchList[len(b.BatchList)-1].Attr.Activationdate = StringToInt64(string(val))
	case "Process Start Date":
		b.BatchList[len(b.BatchList)-1].Attr.ProcessStartDate = StringToInt64(string(val))
	case "Protect Stop Date":
		b.BatchList[len(b.BatchList)-1].Attr.ProcessStopDate = StringToInt64(string(val))
	case "Deactivation Date":
		b.BatchList[len(b.BatchList)-1].Attr.Deactivationdate = StringToInt64(string(val))
	case "Destroy Date":
		b.BatchList[len(b.BatchList)-1].Attr.DestroyDate = StringToInt64(string(val))
	case "Compromise Occurrence Date":
		b.BatchList[len(b.BatchList)-1].Attr.CompromiseOccuranceDate = StringToInt64(string(val))
	case "Compromise Date":
		b.BatchList[len(b.BatchList)-1].Attr.CompromiseDate = StringToInt64(string(val))
	case "Archive Date":
		b.BatchList[len(b.BatchList)-1].Attr.ArchiveDate = StringToInt64(string(val))
	case "Object Group":
		b.BatchList[len(b.BatchList)-1].Attr.ObjectGroup = BinToString(val)
	case "Fresh":
		b.BatchList[len(b.BatchList)-1].Attr.Fresh = StringToInt64(string(val))
	case "Contact Information":
		b.BatchList[len(b.BatchList)-1].Attr.ContactInformation = BinToString(val)
	case "Last Change Date":
		b.BatchList[len(b.BatchList)-1].Attr.LastChangeDate = StringToInt64(string(val))
	case "Key Value Present":
		b.BatchList[len(b.BatchList)-1].Attr.KeyValuePresent = StringToInt64(string(val))
	case "Original Creation Date":
		b.BatchList[len(b.BatchList)-1].Attr.OriginalCreationDate = StringToInt64(string(val))
	}
}
