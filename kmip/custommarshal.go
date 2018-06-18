package kmip

type Marsheler interface {
	CustomMarshalKMIP() string
}