/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KeptnTaskSpec defines the desired state of KeptnTask
type KeptnTaskSpec struct {
	Workload         string           `json:"workload"`
	WorkloadVersion  string           `json:"workloadVersion"`
	AppName          string           `json:"app"`
	TaskDefinition   string           `json:"taskDefinition"`
	Parameters       TaskParameters   `json:"parameters,omitempty"`
	SecureParameters SecureParameters `json:"secureParameters,omitempty"`
}

type KeptnTaskPhase string

const (
	// TaskPending means the task has been accepted by the system, but the corresponding Job did not start
	TaskPending KeptnTaskPhase = "Pending"
	// TaskRunning means that the Job has been started.
	TaskRunning KeptnTaskPhase = "Running"
	// TaskFailed means that the Job failed
	TaskFailed KeptnTaskPhase = "Failed"
	// TaskSucceeded means that the Job has finished successfully
	TaskSucceeded KeptnTaskPhase = "Succeeded"
)

type TaskParameters struct {
	Inline map[string]string `json:"map,omitempty"`
}

type SecureParameters struct {
	Secret string `json:"secret,omitempty"`
}

// KeptnTaskStatus defines the observed state of KeptnTask
type KeptnTaskStatus struct {
	JobName string         `json:"jobName,omitempty"`
	Status  KeptnTaskPhase `json:"status,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AppName",type=string,JSONPath=`.spec.app`
// +kubebuilder:printcolumn:name="Workload",type=string,JSONPath=`.spec.workload`
// +kubebuilder:printcolumn:name="Version",type=string,JSONPath=`.spec.workloadVersion`
// +kubebuilder:printcolumn:name="Job Name",type=string,JSONPath=`.status.jobName`
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`

// KeptnTask is the Schema for the keptntasks API
type KeptnTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeptnTaskSpec   `json:"spec,omitempty"`
	Status KeptnTaskStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KeptnTaskList contains a list of KeptnTask
type KeptnTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeptnTask `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeptnTask{}, &KeptnTaskList{})
}
