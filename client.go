package soap

import (
	"github.com/deb-ict/go-soap/soap11"
	"github.com/deb-ict/go-soap/soap12"
	"github.com/deb-ict/go-xml"
)

type Client interface {
	GetNametable() map[string]string
	RegisterNamespace(namespace string, prefix string)
	RegisterType(space string, local string, constructor xml.TypeConstructor)
}

func NewClient(version SoapVersion) Client {
	switch version {
	case SoapVersion11:
		return newSoap11Client()
	case SoapVersion12:
		return newSoap12Client()
	}
	return nil
}

func newSoap11Client() Client {
	return soap11.NewClient()
}

func newSoap12Client() Client {
	return soap12.NewClient()
}
