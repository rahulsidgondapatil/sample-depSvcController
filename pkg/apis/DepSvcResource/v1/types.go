package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DepSvcResource is a top-level type
type DepSvcResource struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Status DepSvcResourceStatus `json:"status,omitempty"`
	// This is where you can define
	// your own custom spec
	Spec DepSvcResourceSpec `json:"spec,omitempty"`
}

// custom spec
type DepSvcResourceSpec struct {
	Message string `json:"message,omitempty"`
}

// custom status
type DepSvcResourceStatus struct {
	Name string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// no client needed for list as it's been created in above
type DepSvcResourceList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []DepSvcResource `json:"items"`
}
