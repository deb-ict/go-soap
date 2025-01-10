package soap

import (
	"github.com/deb-ict/go-soap/soap11"
	"github.com/deb-ict/go-soap/soap12"
)

type Envelope interface {
	AddHeaderContent(value ...any)
	AddBodyContent(value ...any)
	GetHeaderContent() []any
	GetBodyContent() []any
}

func NewEnvelope(version SoapVersion) Envelope {
	switch version {
	case SoapVersion11:
		return newSoap11Envelope()
	case SoapVersion12:
		return newSoap12Envelope()
	}
	return nil
}

func newSoap11Envelope() Envelope {
	return soap11.NewEnvelope()
}

func newSoap12Envelope() Envelope {
	return soap12.NewEnvelope()
}
