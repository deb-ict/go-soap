package soap11

import (
	"github.com/deb-ict/go-xml"
)

const (
	Version     = "1.1"
	ContentType = "text/xml; charset=\"utf-8\""
	Namespace   = "http://schemas.xmlsoap.org/soap/envelope/"
)

type Envelope struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Attrs   []*xml.Attr `xml:",any,attr"`
	Header  *Header     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Body    *Body       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type Header struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Attrs   []*xml.Attr `xml:",any,attr"`
	Content []any       `xml:",omitempty"`
}

type Body struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Attrs   []*xml.Attr `xml:",any,attr"`
	Content []any       `xml:",omitempty"`
}

type Fault struct {
	XMLName xml.Name     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	Attrs   []*xml.Attr  `xml:",any,attr"`
	Code    string       `xml:"http://schemas.xmlsoap.org/soap/envelope/ faultcode"`
	String  string       `xml:"http://schemas.xmlsoap.org/soap/envelope/ faultstring"`
	Actor   string       `xml:"http://schemas.xmlsoap.org/soap/envelope/ faultactor"`
	Detail  *FaultDetail `xml:"http://schemas.xmlsoap.org/soap/envelope/ detail"`
}

type FaultDetail struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ detail"`
	Attrs   []*xml.Attr
	Content []any `xml:",any,omitempty"`
}

func NewEnvelope() *Envelope {
	return &Envelope{
		Body: &Body{},
	}
}

func (env *Envelope) AddHeaderContent(value ...any) {
	if env.Header == nil {
		env.Header = &Header{}
	}
	if env.Header.Content == nil {
		env.Header.Content = make([]any, 0)
	}
	env.Header.Content = append(env.Header.Content, value...)
}

func (env *Envelope) AddBodyContent(value ...any) {
	if env.Body == nil {
		env.Body = &Body{}
	}
	if env.Body.Content == nil {
		env.Body.Content = make([]any, 0)
	}
	env.Body.Content = append(env.Body.Content, value...)
}

func (env *Envelope) GetHeaderContent() []any {
	if env.Header != nil && env.Header.Content != nil {
		return env.Header.Content
	}
	return []any{}
}

func (env *Envelope) GetBodyContent() []any {
	if env.Body != nil && env.Body.Content != nil {
		return env.Body.Content
	}
	return []any{}
}
