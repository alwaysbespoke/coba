package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SIPCall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              SIPCallSpec   `json:"spec"`
	Status            SIPCallStatus `json:"status"`
}

// +k8s:deepcopy-gen=true
type SIPCallSpec struct {
	SbcID    string           `json:"sbc-id"`
	Messages []SIPCallMessage `json:"messages"`
}

// +k8s:deepcopy-gen=true
type SIPCallStatus struct {
	State   string `json:"state"`
	Message string `json:"message"`
}

// +k8s:deepcopy-gen=true
type SIPCallMessage struct {
	ID           string      `json:"id"`
	Method       Method      `json:"method"`
	StatusCode   int         `json:"status-code"`
	ReasonPhrase string      `json:"reason-phrase"`
	Created      metav1.Time `json:"created"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SIPCallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []SIPCall `json:"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SIPMessage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              SIPMessageSpec   `json:"spec"`
	Status            SIPMessageStatus `json:"status"`
}

// +k8s:deepcopy-gen=true
type SIPMessageSpec struct {
	RequestLine *RequestLine        `json:"request-line"`
	StatusLine  *StatusLine         `json:"status-line"`
	Headers     map[string][]string `json:"headers"`
	Body        []byte              `json:"body"`
}

// +k8s:deepcopy-gen=true
type SIPMessageStatus struct {
	State   string `json:"state"`
	Message string `json:"message"`
}

// +k8s:deepcopy-gen=true
type RequestLine struct {
	Method     Method     `json:"method"`
	RequestURI string     `json:"request-uri"`
	SIPVersion SIPVersion `json:"sip-version"`
}

// +k8s:deepcopy-gen=true
type StatusLine struct {
	SIPVersion   SIPVersion `json:"sip-version"`
	StatusCode   int        `json:"status-code"`
	ReasonPhrase string     `json:"reason-phrase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SIPMessageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []SIPMessage `json:"items"`
}
