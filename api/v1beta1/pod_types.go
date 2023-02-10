/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PodSpec defines the desired state of Pod
type PodSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	UID         string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	ProcessName string `json:"processName,omitempty"`
	PID         uint32 `json:"pid,omitempty"`
	Memory      int64  `json:"memory,omitempty"`
	CPU         string `json:"cpu,omitempty"`
	NetworkRX   int64  `json:"networkRX,omitempty"`
	NetworkTX   int64  `json:"networkTX,omitempty"`
	Storage     int64  `json:"storage,omitempty"`
	ContainerID string `json:"containerID,omitempty"`
}

// PodStatus defines the observed state of Pod
type PodStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	CreateTime string `json:"createTime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Pod is the Schema for the pods API
type Pod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodSpec   `json:"spec,omitempty"`
	Status PodStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PodList contains a list of Pod
type PodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pod `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Pod{}, &PodList{})
}
