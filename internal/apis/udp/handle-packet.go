package udp

import (
	"context"
	"net"

	"github.com/google/uuid"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/alwaysbespoke/coba/internal/clients/sbcs"
	v1 "github.com/alwaysbespoke/coba/pkg/crds/sip/v1"
)

func (a *API) handlePacket(conn net.PacketConn, addr net.Addr, buf []byte) {
	// validate the SIP message
	sipMessage := &v1.SIPMessage{}
	if err := sipMessage.Unmarshal(buf); err != nil {
		a.Logger.Errorf("failed to unmarshal message: %w", err)

		// todo: write back an error message to the client
		if _, err := conn.WriteTo([]byte(`insert error message`), addr); err != nil {
			a.Logger.Errorf("failed to write error response: %w", err)
		}

		return
	}

	// get the Call-ID header from the SIP message
	callIDHeader, ok := sipMessage.GetHeader(v1.CallID.String())
	if callIDHeader == nil || !ok {
		a.Logger.Error("failed to get Call-ID")

		// todo: write back an error message to the client
		if _, err := conn.WriteTo([]byte(`insert error message`), addr); err != nil {
			a.Logger.Errorf("failed to write error response: %w", err)
		}

		return
	}

	// get the Call-ID
	// since headers are stored as a slice and Call-ID should have only one value,
	// the Call-ID header is the first index
	callID := callIDHeader[0]

	// query the KubeAPI to check if the SIPCall object exists
	// todo: add retry logic
	sipCall, err := a.K8Clients.SipV1Client.SIPCalls(a.Config.Namespace).Get(context.Background(), callID, metav1.GetOptions{})
	if err != nil {
		a.Logger.Errorf("failed to get SIPCall: %w", err)

		// todo: write back an error message to the client
		if _, err := conn.WriteTo([]byte(`insert error message`), addr); err != nil {
			a.Logger.Errorf("failed to write error response: %w", err)
		}

		return
	}

	var sbc *sbcs.SBC

	if sipCall == nil {
		sipCall = v1.NewSIPCall(sipMessage)

		// assign an SBC to the SIPCall object
		sbc = a.SbcsClient.AssignSBC()
		sipCall.Spec.SbcID = sbc.Obj.Name

		// if the SIPCall object does not exist, create a new SIPCall object
		// todo: add retry logic
		if _, err := a.K8Clients.SipV1Client.SIPCalls(a.Config.Namespace).Create(context.Background(), sipCall, metav1.CreateOptions{}); err != nil {
			a.Logger.Errorf("failed to create SIPCall: %w", err)

			// todo: write back an error message to the client
			if _, err := conn.WriteTo([]byte(`insert error message`), addr); err != nil {
				a.Logger.Errorf("failed to write error response: %w", err)
			}

			return
		}
	} else {
		// if the SIPCall object exists, update the SIPCall object

		// create the SIPCallMessage
		sipCallMessage := v1.SIPCallMessage{
			ID: sipMessage.Name,
		}

		if sipMessage.Spec.StatusLine != nil {
			sipCallMessage.StatusCode = sipMessage.Spec.StatusLine.StatusCode
			sipCallMessage.ReasonPhrase = sipMessage.Spec.StatusLine.ReasonPhrase
		} else {
			sipCallMessage.Method = sipMessage.Spec.RequestLine.Method
		}

		sipCall.Spec.Messages = append(sipCall.Spec.Messages, sipCallMessage)

		// todo: add retry logic
		if _, err := a.K8Clients.SipV1Client.SIPCalls(a.Config.Namespace).Update(context.Background(), sipCall, metav1.UpdateOptions{}); err != nil {
			a.Logger.Errorf("failed to update SIPCall: %w", err)

			// todo: write back an error message to the client
			if _, err := conn.WriteTo([]byte(`insert error message`), addr); err != nil {
				a.Logger.Errorf("failed to write error response: %w", err)
			}

			return
		}

		// get the SBC associated with the call
		sbc, ok = a.SbcsClient.GetSBC(sipCall.Spec.SbcID)
		if !ok {
			a.Logger.Errorf("SBC (%s) could not be found for Call-ID (%s): %w", sipCall.Spec.SbcID, callID, err)

			// todo: write back an error message to the client
			if _, err := conn.WriteTo([]byte(`insert error message`), addr); err != nil {
				a.Logger.Errorf("failed to write error response: %w", err)
			}

			return
		}
	}

	// create a UUID for the new SIPMessage and set it as the name
	sipMessage.Name = uuid.New().String()

	// create a new SIPMessage
	// todo: add retry logic
	if _, err := a.K8Clients.SipV1Client.SIPMessages(a.Config.Namespace).Create(context.Background(), sipMessage, metav1.CreateOptions{}); err != nil {
		a.Logger.Errorf("failed to create SIPMessage: %w", err)

		// todo: write back an error message to the client
		if _, err := conn.WriteTo([]byte(`insert error message`), addr); err != nil {
			a.Logger.Errorf("failed to write error response: %w", err)
		}

		return
	}

	// marshal the SIPMessage to SIP
	sipMessageBytes, err := sipMessage.MarshalToSIP()
	if err != nil {
		a.Logger.Errorf("failed to marshal SIPMessage: %w", err)

		// todo: write back an error message to the client
		if _, err := conn.WriteTo([]byte(`insert error message`), addr); err != nil {
			a.Logger.Errorf("failed to write error response: %w", err)
		}

		return
	}

	// todo: handle business logic

	// proxy the SIP message
	if _, err := sbc.Conn.Write(sipMessageBytes); err != nil {
		a.Logger.Errorf("failed to proxy SIP message: %w", err)

		// todo: write back an error message to the client
		if _, err := conn.WriteTo([]byte(`insert error message`), addr); err != nil {
			a.Logger.Errorf("failed to write error response: %w", err)
		}

		return
	}

	// todo: handle the response from the SBC
}
