package soap

type SoapVersion int

const (
	SoapVersion11 = iota
	SoapVersion12
)

func (v SoapVersion) String() string {
	switch v {
	case SoapVersion11:
		return "1.1"
	case SoapVersion12:
		return "1.2"
	}
	return "unknown"
}
