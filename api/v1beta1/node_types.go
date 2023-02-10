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

// NodeSpec defines the desired state of Node
type NodeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Name         string   `json:"name,omitempty"`
	TotalCPU     int64    `json:"totalCPU,omitempty"`
	CPU          int64    `json:"cpu,omitempty"`
	Count        int      `json:"count,omitempty"`
	TotalMemory  int64    `json:"totalMemory,omitempty"`
	Memory       int64    `json:"memory,omitempty"`
	TotalStorage int64    `json:"totalStorage,omitempty"`
	Storage      int64    `json:"storage,omitempty"`
	NetworkRX    int64    `json:"networkRX,omitempty"`
	NetworkTX    int64    `json:"networkTX,omitempty"`
	TotalPodnum  int      `json:"totalPodnum,omitempty"`
	UUID         []string `json:"uuid,omitempty"`
	NvLinkInfo   []NvLink `json:"nvlinkInfo,omitempty"`
}

type NvLink struct {
	GPUUUID   map[string]string `json:"gpuUUID,omitempty"`
	CountLink int               `json:"countLink,omitempty"`
}

// NodeStatus defines the observed state of Node
type NodeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	CreateTime string `json:"createTime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Node is the Schema for the nodes API
type Node struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeSpec   `json:"spec,omitempty"`
	Status NodeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NodeList contains a list of Node
type NodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Node `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Node{}, &NodeList{})
}
