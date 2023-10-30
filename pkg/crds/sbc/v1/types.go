package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SBC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              SBCSpec   `json:"spec"`
	Status            SBCStatus `json:"status"`
}

// +k8s:deepcopy-gen=true
type SBCSpec struct {
	Address string `json:"address"`
	Region  string `json:"region"`
	AZ      string `json:"az"`
}

// +k8s:deepcopy-gen=true
type SBCStatus struct {
	State   string `json:"state"`
	Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SBCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []SBC `json:"items"`
}
