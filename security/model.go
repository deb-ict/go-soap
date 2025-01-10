package security

import (
	"github.com/deb-ict/go-xml"
)

const (
	WsseNamespace = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WsuNamespace  = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
)

type Security struct {
	XMLName xml.Name    `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security"`
	Attrs   []*xml.Attr `xml:",any,attr"`
	Content []any       `xml:",any,omitempty"`
}

type BinarySecurityToken struct {
	XMLName      xml.Name    `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd BinarySecurityToken"`
	Attrs        []*xml.Attr `xml:",any,attr"`
	WsuId        string      `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd Id,attr,omitempty"`
	EncodingType string      `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd EncodingType,attr,omitempty"`
	ValueType    string      `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd ValueType,attr,omitempty"`
	Value        string      `xml:",chardata"`
}
