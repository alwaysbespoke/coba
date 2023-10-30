package v1

import (
	"github.com/google/uuid"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewSIPCall(sipMessage *SIPMessage) *SIPCall {
	callID := sipMessage.Spec.Headers[CallID.String()][0]

	var (
		statusCode   int
		reasonPhrase string
		method       Method
	)

	if sipMessage.Spec.StatusLine != nil {
		statusCode = sipMessage.Spec.StatusLine.StatusCode
		reasonPhrase = sipMessage.Spec.StatusLine.ReasonPhrase
	} else {
		method = sipMessage.Spec.RequestLine.Method
	}

	return &SIPCall{
		ObjectMeta: metav1.ObjectMeta{
			Name:      callID,
			Namespace: sipMessage.Name,
		},
		Spec: SIPCallSpec{
			Messages: []SIPCallMessage{
				{
					ID:           uuid.New().String(),
					Method:       method,
					ReasonPhrase: reasonPhrase,
					StatusCode:   statusCode,
					Created:      metav1.Now(),
				},
			},
		},
	}
}
