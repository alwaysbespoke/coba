package v1

import "bytes"

// Unmarshal parses a SIP message stored in a byte slice and stores it in a SIPMessage object
func (m *SIPMessage) Unmarshal(data []byte) error {
	parser := newMessageParser(data, m)

	return parser.parse()
}

// MarshalToSIP parses the SIPMessage object and returns a valid SIP message
func (m *SIPMessage) MarshalToSIP() ([]byte, error) {
	buf := &bytes.Buffer{}
	return buf.Bytes(), nil
}

// GetHeader returns a slice of strings of values for that header
// name and a bool if it exists
func (m *SIPMessage) GetHeader(headerStr string) ([]string, bool) {
	if m.Spec.Headers != nil || len(m.Spec.Headers) == 0 {
		header, ok := m.Spec.Headers[headerStr]

		return header, ok
	}

	return nil, false
}
