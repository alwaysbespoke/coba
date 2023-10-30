package v1

type Header string

const (
	CallID        Header = "Call-ID"
	Contact       Header = "Contact"
	ContentLength Header = "Content-Length"
	CSeq          Header = "CSeq"
	From          Header = "From"
	To            Header = "To"
	Via           Header = "Via"
)

const (
	requiredHeaderCount = 5
)

func (h Header) String() string {
	return string(h)
}
