package soap12

import (
	"github.com/deb-ict/go-xml"
)

const (
	Version     = "1.2"
	ContentType = "application/soap+xml"
	Namespace   = "http://www.w3.org/2003/05/soap-envelope"
)

type Envelope struct {
	XMLName xml.Name    `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Attrs   []*xml.Attr `xml:",any,attr"`
	Header  *Header     `xml:"http://www.w3.org/2003/05/soap-envelope Header"`
	Body    *Body       `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
	RawXml  []byte      `xml:",innerxml"`
}

type Header struct {
	XMLName xml.Name    `xml:"http://www.w3.org/2003/05/soap-envelope Header"`
	Attrs   []*xml.Attr `xml:",any,attr"`
	Content []any       `xml:",any,omitempty"`
	RawXml  []byte      `xml:",innerxml"`
}

type Body struct {
	XMLName xml.Name    `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
	Attrs   []*xml.Attr `xml:",any,attr"`
	Content []any       `xml:",any,omitempty"`
	RawXml  []byte      `xml:",innerxml"`
}

type Fault struct {
	XMLName xml.Name     `xml:"http://www.w3.org/2003/05/soap-envelope Fault"`
	Attrs   []*xml.Attr  `xml:",any,attr"`
	Code    string       `xml:"http://www.w3.org/2003/05/soap-envelope Code"`
	Reason  *FaultReason `xml:"http://www.w3.org/2003/05/soap-envelope Reason"`
	Node    string       `xml:"http://www.w3.org/2003/05/soap-envelope Node"`
	Role    string       `xml:"http://www.w3.org/2003/05/soap-envelope Role"`
	Detail  *FaultDetail `xml:"http://www.w3.org/2003/05/soap-envelope Detail"`
}

type FaultReason struct {
	XMLName xml.Name          `xml:"http://www.w3.org/2003/05/soap-envelope Reason"`
	Attrs   []*xml.Attr       `xml:",any,attr"`
	Text    []FaultReasonText `xml:"http://www.w3.org/2003/05/soap-envelope Text"`
}

type FaultReasonText struct {
	XMLName  xml.Name    `xml:"http://www.w3.org/2003/05/soap-envelope Text"`
	Attrs    []*xml.Attr `xml:",any,attr"`
	Language string      `xml:"http://www.w3.org/2003/05/soap-envelope lang,attr"`
	Text     string      `xml:",chardata"`
}

type FaultDetail struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Detail"`
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
