package v1

import (
	"errors"
	"strconv"
	"strings"
)

type messageParser struct {
	buf                  *strings.Builder
	data                 []byte
	headerName           *string
	lines                int
	message              *SIPMessage
	requiredHeadersCount int // To, From, Call-ID, CSeq, Contact
	requiredHeaders      map[string]int
	spaces               int
}

func newMessageParser(data []byte, message *SIPMessage) *messageParser {
	return &messageParser{
		buf:     &strings.Builder{},
		data:    data,
		message: message,
		requiredHeaders: map[string]int{
			To.String():     0,
			From.String():   0,
			CallID.String(): 0,
			CSeq.String():   0, // todo: confirm if required
			Via.String():    0, // todo: confirm if required
			// Contact.String(): 0, // todo: confirm if required
			// ContentLength.String(): 0, // todo: confirm if required
		},
	}
}

func (p *messageParser) parse() error {
	for i, b := range p.data {
		char := string(b)

		// validate new line format
		switch char {
		case "\r":
			if i < len(p.data)-1 && string(p.data[i+1]) != "\n" {
				return errors.New("invalid new line format: only CRLF allowed")
			}
		case "\n":
			if i > 0 && string(p.data[i-1]) != "\r" {
				return errors.New("invalid new line format: only CRLF allowed")
			}
		}

		// parse the message body
		if i > 3 &&
			string(p.data[i-3]) == "\r" &&
			string(p.data[i-2]) == "\n" &&
			string(p.data[i-1]) == "\r" &&
			string(p.data[i]) == "\n" {

			// validate Start-Line and Headers are present
			switch {
			case p.message.Spec.RequestLine == nil &&
				p.message.Spec.StatusLine == nil:
				return errors.New("invalid message: invalid Start-Line")
			case p.requiredHeadersCount != requiredHeaderCount:
				return errors.New("invalid message: missing required Headers")
			}

			// set the Body
			if len(p.data) == i+1 {
				p.message.Spec.Body = []byte{}
			} else {
				p.message.Spec.Body = p.data[i+1:]
			}

			return nil
		}

		// parse the start line and headers
		switch p.lines {
		case 0:
			if err := p.parseStartLine(char); err != nil {
				return err
			}
		default:
			if err := p.parseHeader(char); err != nil {
				return err
			}
		}
	}

	return errors.New("invalid message format")
}

// RFC 3261 Sections 7.1 and 7.2 define Start-Line syntax
// Section 7.1: Requests
// Request-Line  =  Method SP Request-URI SP SIP-Version CRLF
// Section 7.2: Responses
// Status-Line  =  SIP-Version SP Status-Code SP Reason-Phrase CRLF
// todo: exit early if Method or SIP-Version is greater than 20 characters
// todo: exit early if Status-Code is greater than 3 characters
// todo: exit early if SIP-Version is greater than 7 characters
func (p *messageParser) parseStartLine(char string) error {
	flush := false

	switch char {
	case "\r":
		// do not write CR
	case "\n":
		flush = true
		p.lines++
	case " ":
		flush = true
		p.spaces++
	default:
		p.buf.WriteString(char)
		return nil
	}

	if flush {
		switch {
		case p.lines == 0 && p.spaces == 1:
			return p.parseMethodOrSIPVersion()
		case p.lines == 0 && p.spaces == 2 && p.message.Spec.RequestLine != nil:
			return p.parseRequestURI()
		case p.lines == 1 && p.spaces == 2 && p.message.Spec.RequestLine != nil:
			return p.parseSIPVersion()
		case ((p.lines == 0 && p.spaces == 2) ||
			(p.lines == 1 && p.spaces == 1)) &&
			p.message.Spec.StatusLine != nil:
			return p.parseStatusCode()
		case p.lines == 1 && p.spaces >= 2 && p.message.Spec.StatusLine != nil:
			return p.parseReasonPhrase()
		default:
			return errors.New("invalid Start-Line")
		}

	}

	return nil
}

// RFC 3621 Section 7.3 defines Header syntax
// todo: remove white space from left and right of colon
// todo: handle folding
func (p *messageParser) parseHeader(char string) error {
	if p.message.Spec.Headers == nil {
		p.message.Spec.Headers = map[string][]string{}
	}

	switch char {
	case "\r":
		// do not write CR
	case "\n":
		p.setHeaderValue()
	case ":":
		if p.headerName == nil {
			p.setHeaderName()
		} else {
			p.buf.WriteString(char)
		}
	default:
		p.buf.WriteString(char)
	}

	return nil
}

func (p *messageParser) parseMethodOrSIPVersion() error {
	methodOrSIPVersionStr := p.buf.String()

	// Request-Line: parse the Method
	if method, isMethod := MethodFromString(methodOrSIPVersionStr); isMethod {
		p.message.Spec.RequestLine = &RequestLine{
			Method: method,
		}

		p.buf.Reset()

		return nil
	}

	// Status-Line: parse the SIP-Version
	if version, isVersion := VersionFromString(methodOrSIPVersionStr); isVersion {
		p.message.Spec.StatusLine = &StatusLine{
			SIPVersion: version,
		}

		p.buf.Reset()

		return nil
	}

	return errors.New("invalid Start-Line: invalid Method or SIP-Version")
}

// todo: add validation
func (p *messageParser) parseRequestURI() error {
	reqURIStr := p.buf.String()

	p.message.Spec.RequestLine.RequestURI = reqURIStr

	p.buf.Reset()

	return nil
}

func (p *messageParser) parseStatusCode() error {
	statusCodeStr := p.buf.String()

	statusCode, err := strconv.Atoi(statusCodeStr)
	if err != nil || statusCode < 100 || statusCode > 699 {
		return errors.New("invalid Status-Line: must be an integer between 100 and 699")
	}
	p.message.Spec.StatusLine.StatusCode = statusCode

	p.buf.Reset()

	return nil
}

func (p *messageParser) parseSIPVersion() error {
	versionStr := p.buf.String()

	if version, isVersion := VersionFromString(versionStr); isVersion {
		p.message.Spec.RequestLine.SIPVersion = version

		p.buf.Reset()

		return nil
	}

	return errors.New("invalid Request-Line: invalid SIP-Version")
}

func (p *messageParser) parseReasonPhrase() error {
	phrase := p.buf.String()
	phrase = strings.TrimSpace(phrase) // todo: replace with in-place parsing

	p.message.Spec.StatusLine.ReasonPhrase = phrase

	p.buf.Reset()

	return nil
}

func (p *messageParser) setHeaderName() {
	name := p.buf.String()
	name = strings.TrimSpace(name) // todo: replace with in-place parsing
	p.headerName = &name
	p.buf.Reset()
}

func (p *messageParser) setHeaderValue() error {
	name := *p.headerName

	value := p.buf.String()
	value = strings.TrimSpace(value) // todo: replace with in-place parsing

	v, ok := p.requiredHeaders[name]
	if ok {
		if v > 0 {
			// todo: validate that required headers are not repeatable
			return errors.New("invalid Header: repeated required Header")
		}
		p.requiredHeaders[name] = 1
		p.requiredHeadersCount++
	}

	p.message.Spec.Headers[name] = append(p.message.Spec.Headers[name], value)
	p.headerName = nil
	p.buf.Reset()

	return nil
}
