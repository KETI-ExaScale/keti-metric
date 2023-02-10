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

// GPUSpec defines the desired state of GPU
type GPUSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	UUID        string         `json:"uuid,omitempty"`
	Total       int64          `json:"total,omitempty"`
	Free        int64          `json:"free,omitempty"`
	Used        int64          `json:"used,omitempty"`
	Memory      int64          `json:"memory,omitempty"`
	Name        string         `json:"name,omitempty"`
	Index       int            `json:"index,omitempty"`
	Temperature GPUTemperature `json:"temperature,omitempty"`
	Power       int            `json:"power,omitempty"`
	MPScount    int            `json:"mpsCount,omitempty"`
	Totalpower  int            `json:"totalPower,omitempty"`
	Flops       int            `json:"flops,omitempty"`
	Arch        int            `json:"arch,omitempty"`
	Util        int            `json:"util,omitempty"`
	FanSpeed    int            `json:"fanSpeed,omitempty"`
	RX          int            `json:"rx,omitempty"`
	TX          int            `json:"tx,omitempty"`
	Assingment  int            `json:"assingment,omitempty"`
	Return      int            `json:"return,omitempty"`
}

type GPUTemperature struct {
	Current      int `json:"current,omitempty"`      //현재 온도
	Threshold    int `json:"threshold,omitempty"`    //성능 저하 온도
	Shutdown     int `json:"shutdown,omitempty"`     //강제 종료 온도
	MaxOperating int `json:"maxOperating,omitempty"` //성능 유지 최대 온도
}

// GPUStatus defines the observed state of GPU
type GPUStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	CreateTime string `json:"createTime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GPU is the Schema for the gpus API
type GPU struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GPUSpec   `json:"spec,omitempty"`
	Status GPUStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GPUList contains a list of GPU
type GPUList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GPU `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GPU{}, &GPUList{})
}
