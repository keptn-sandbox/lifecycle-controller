//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha1/common"
	"go.opentelemetry.io/otel/propagation"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapReference) DeepCopyInto(out *ConfigMapReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapReference.
func (in *ConfigMapReference) DeepCopy() *ConfigMapReference {
	if in == nil {
		return nil
	}
	out := new(ConfigMapReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerSpec) DeepCopyInto(out *ContainerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerSpec.
func (in *ContainerSpec) DeepCopy() *ContainerSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvaluationStatus) DeepCopyInto(out *EvaluationStatus) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvaluationStatus.
func (in *EvaluationStatus) DeepCopy() *EvaluationStatus {
	if in == nil {
		return nil
	}
	out := new(EvaluationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvaluationStatusItem) DeepCopyInto(out *EvaluationStatusItem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvaluationStatusItem.
func (in *EvaluationStatusItem) DeepCopy() *EvaluationStatusItem {
	if in == nil {
		return nil
	}
	out := new(EvaluationStatusItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionReference) DeepCopyInto(out *FunctionReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionReference.
func (in *FunctionReference) DeepCopy() *FunctionReference {
	if in == nil {
		return nil
	}
	out := new(FunctionReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	out.FunctionReference = in.FunctionReference
	out.Inline = in.Inline
	out.HttpReference = in.HttpReference
	out.ConfigMapReference = in.ConfigMapReference
	in.Parameters.DeepCopyInto(&out.Parameters)
	out.SecureParameters = in.SecureParameters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpReference) DeepCopyInto(out *HttpReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpReference.
func (in *HttpReference) DeepCopy() *HttpReference {
	if in == nil {
		return nil
	}
	out := new(HttpReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Inline) DeepCopyInto(out *Inline) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Inline.
func (in *Inline) DeepCopy() *Inline {
	if in == nil {
		return nil
	}
	out := new(Inline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnApp) DeepCopyInto(out *KeptnApp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnApp.
func (in *KeptnApp) DeepCopy() *KeptnApp {
	if in == nil {
		return nil
	}
	out := new(KeptnApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnApp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnAppList) DeepCopyInto(out *KeptnAppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnApp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnAppList.
func (in *KeptnAppList) DeepCopy() *KeptnAppList {
	if in == nil {
		return nil
	}
	out := new(KeptnAppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnAppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnAppSpec) DeepCopyInto(out *KeptnAppSpec) {
	*out = *in
	if in.Workloads != nil {
		in, out := &in.Workloads, &out.Workloads
		*out = make([]KeptnWorkloadRef, len(*in))
		copy(*out, *in)
	}
	if in.PreDeploymentTasks != nil {
		in, out := &in.PreDeploymentTasks, &out.PreDeploymentTasks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PostDeploymentTasks != nil {
		in, out := &in.PostDeploymentTasks, &out.PostDeploymentTasks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PreDeploymentEvaluations != nil {
		in, out := &in.PreDeploymentEvaluations, &out.PreDeploymentEvaluations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PostDeploymentEvaluations != nil {
		in, out := &in.PostDeploymentEvaluations, &out.PostDeploymentEvaluations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnAppSpec.
func (in *KeptnAppSpec) DeepCopy() *KeptnAppSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnAppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnAppStatus) DeepCopyInto(out *KeptnAppStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnAppStatus.
func (in *KeptnAppStatus) DeepCopy() *KeptnAppStatus {
	if in == nil {
		return nil
	}
	out := new(KeptnAppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnAppVersion) DeepCopyInto(out *KeptnAppVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnAppVersion.
func (in *KeptnAppVersion) DeepCopy() *KeptnAppVersion {
	if in == nil {
		return nil
	}
	out := new(KeptnAppVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnAppVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnAppVersionList) DeepCopyInto(out *KeptnAppVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnAppVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnAppVersionList.
func (in *KeptnAppVersionList) DeepCopy() *KeptnAppVersionList {
	if in == nil {
		return nil
	}
	out := new(KeptnAppVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnAppVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnAppVersionSpec) DeepCopyInto(out *KeptnAppVersionSpec) {
	*out = *in
	in.KeptnAppSpec.DeepCopyInto(&out.KeptnAppSpec)
	if in.TraceId != nil {
		in, out := &in.TraceId, &out.TraceId
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnAppVersionSpec.
func (in *KeptnAppVersionSpec) DeepCopy() *KeptnAppVersionSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnAppVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnAppVersionStatus) DeepCopyInto(out *KeptnAppVersionStatus) {
	*out = *in
	if in.WorkloadStatus != nil {
		in, out := &in.WorkloadStatus, &out.WorkloadStatus
		*out = make([]WorkloadStatus, len(*in))
		copy(*out, *in)
	}
	if in.PreDeploymentTaskStatus != nil {
		in, out := &in.PreDeploymentTaskStatus, &out.PreDeploymentTaskStatus
		*out = make([]TaskStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PostDeploymentTaskStatus != nil {
		in, out := &in.PostDeploymentTaskStatus, &out.PostDeploymentTaskStatus
		*out = make([]TaskStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PreDeploymentEvaluationTaskStatus != nil {
		in, out := &in.PreDeploymentEvaluationTaskStatus, &out.PreDeploymentEvaluationTaskStatus
		*out = make([]EvaluationStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PostDeploymentEvaluationTaskStatus != nil {
		in, out := &in.PostDeploymentEvaluationTaskStatus, &out.PostDeploymentEvaluationTaskStatus
		*out = make([]EvaluationStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PhaseTraceIDs != nil {
		in, out := &in.PhaseTraceIDs, &out.PhaseTraceIDs
		*out = make(common.PhaseTraceID, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(propagation.MapCarrier, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnAppVersionStatus.
func (in *KeptnAppVersionStatus) DeepCopy() *KeptnAppVersionStatus {
	if in == nil {
		return nil
	}
	out := new(KeptnAppVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnEvaluation) DeepCopyInto(out *KeptnEvaluation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnEvaluation.
func (in *KeptnEvaluation) DeepCopy() *KeptnEvaluation {
	if in == nil {
		return nil
	}
	out := new(KeptnEvaluation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnEvaluation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnEvaluationDefinition) DeepCopyInto(out *KeptnEvaluationDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnEvaluationDefinition.
func (in *KeptnEvaluationDefinition) DeepCopy() *KeptnEvaluationDefinition {
	if in == nil {
		return nil
	}
	out := new(KeptnEvaluationDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnEvaluationDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnEvaluationDefinitionList) DeepCopyInto(out *KeptnEvaluationDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnEvaluationDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnEvaluationDefinitionList.
func (in *KeptnEvaluationDefinitionList) DeepCopy() *KeptnEvaluationDefinitionList {
	if in == nil {
		return nil
	}
	out := new(KeptnEvaluationDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnEvaluationDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnEvaluationDefinitionSpec) DeepCopyInto(out *KeptnEvaluationDefinitionSpec) {
	*out = *in
	if in.Objectives != nil {
		in, out := &in.Objectives, &out.Objectives
		*out = make([]Objective, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnEvaluationDefinitionSpec.
func (in *KeptnEvaluationDefinitionSpec) DeepCopy() *KeptnEvaluationDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnEvaluationDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnEvaluationDefinitionStatus) DeepCopyInto(out *KeptnEvaluationDefinitionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnEvaluationDefinitionStatus.
func (in *KeptnEvaluationDefinitionStatus) DeepCopy() *KeptnEvaluationDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(KeptnEvaluationDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnEvaluationList) DeepCopyInto(out *KeptnEvaluationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnEvaluation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnEvaluationList.
func (in *KeptnEvaluationList) DeepCopy() *KeptnEvaluationList {
	if in == nil {
		return nil
	}
	out := new(KeptnEvaluationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnEvaluationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnEvaluationProvider) DeepCopyInto(out *KeptnEvaluationProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnEvaluationProvider.
func (in *KeptnEvaluationProvider) DeepCopy() *KeptnEvaluationProvider {
	if in == nil {
		return nil
	}
	out := new(KeptnEvaluationProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnEvaluationProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnEvaluationProviderList) DeepCopyInto(out *KeptnEvaluationProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnEvaluationProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnEvaluationProviderList.
func (in *KeptnEvaluationProviderList) DeepCopy() *KeptnEvaluationProviderList {
	if in == nil {
		return nil
	}
	out := new(KeptnEvaluationProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnEvaluationProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnEvaluationProviderSpec) DeepCopyInto(out *KeptnEvaluationProviderSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnEvaluationProviderSpec.
func (in *KeptnEvaluationProviderSpec) DeepCopy() *KeptnEvaluationProviderSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnEvaluationProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnEvaluationProviderStatus) DeepCopyInto(out *KeptnEvaluationProviderStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnEvaluationProviderStatus.
func (in *KeptnEvaluationProviderStatus) DeepCopy() *KeptnEvaluationProviderStatus {
	if in == nil {
		return nil
	}
	out := new(KeptnEvaluationProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnEvaluationSpec) DeepCopyInto(out *KeptnEvaluationSpec) {
	*out = *in
	out.RetryInterval = in.RetryInterval
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnEvaluationSpec.
func (in *KeptnEvaluationSpec) DeepCopy() *KeptnEvaluationSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnEvaluationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnEvaluationStatus) DeepCopyInto(out *KeptnEvaluationStatus) {
	*out = *in
	if in.EvaluationStatus != nil {
		in, out := &in.EvaluationStatus, &out.EvaluationStatus
		*out = make(map[string]EvaluationStatusItem, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnEvaluationStatus.
func (in *KeptnEvaluationStatus) DeepCopy() *KeptnEvaluationStatus {
	if in == nil {
		return nil
	}
	out := new(KeptnEvaluationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnTask) DeepCopyInto(out *KeptnTask) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnTask.
func (in *KeptnTask) DeepCopy() *KeptnTask {
	if in == nil {
		return nil
	}
	out := new(KeptnTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnTask) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnTaskDefinition) DeepCopyInto(out *KeptnTaskDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnTaskDefinition.
func (in *KeptnTaskDefinition) DeepCopy() *KeptnTaskDefinition {
	if in == nil {
		return nil
	}
	out := new(KeptnTaskDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnTaskDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnTaskDefinitionList) DeepCopyInto(out *KeptnTaskDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnTaskDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnTaskDefinitionList.
func (in *KeptnTaskDefinitionList) DeepCopy() *KeptnTaskDefinitionList {
	if in == nil {
		return nil
	}
	out := new(KeptnTaskDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnTaskDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnTaskDefinitionSpec) DeepCopyInto(out *KeptnTaskDefinitionSpec) {
	*out = *in
	in.Function.DeepCopyInto(&out.Function)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnTaskDefinitionSpec.
func (in *KeptnTaskDefinitionSpec) DeepCopy() *KeptnTaskDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnTaskDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnTaskDefinitionStatus) DeepCopyInto(out *KeptnTaskDefinitionStatus) {
	*out = *in
	out.Function = in.Function
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnTaskDefinitionStatus.
func (in *KeptnTaskDefinitionStatus) DeepCopy() *KeptnTaskDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(KeptnTaskDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnTaskList) DeepCopyInto(out *KeptnTaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnTaskList.
func (in *KeptnTaskList) DeepCopy() *KeptnTaskList {
	if in == nil {
		return nil
	}
	out := new(KeptnTaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnTaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnTaskSpec) DeepCopyInto(out *KeptnTaskSpec) {
	*out = *in
	out.Context = in.Context
	in.Parameters.DeepCopyInto(&out.Parameters)
	out.SecureParameters = in.SecureParameters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnTaskSpec.
func (in *KeptnTaskSpec) DeepCopy() *KeptnTaskSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnTaskStatus) DeepCopyInto(out *KeptnTaskStatus) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnTaskStatus.
func (in *KeptnTaskStatus) DeepCopy() *KeptnTaskStatus {
	if in == nil {
		return nil
	}
	out := new(KeptnTaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnWorkload) DeepCopyInto(out *KeptnWorkload) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnWorkload.
func (in *KeptnWorkload) DeepCopy() *KeptnWorkload {
	if in == nil {
		return nil
	}
	out := new(KeptnWorkload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnWorkload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnWorkloadList) DeepCopyInto(out *KeptnWorkloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnWorkload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnWorkloadList.
func (in *KeptnWorkloadList) DeepCopy() *KeptnWorkloadList {
	if in == nil {
		return nil
	}
	out := new(KeptnWorkloadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnWorkloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnWorkloadRef) DeepCopyInto(out *KeptnWorkloadRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnWorkloadRef.
func (in *KeptnWorkloadRef) DeepCopy() *KeptnWorkloadRef {
	if in == nil {
		return nil
	}
	out := new(KeptnWorkloadRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnWorkloadSpec) DeepCopyInto(out *KeptnWorkloadSpec) {
	*out = *in
	if in.PreDeploymentTasks != nil {
		in, out := &in.PreDeploymentTasks, &out.PreDeploymentTasks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PostDeploymentTasks != nil {
		in, out := &in.PostDeploymentTasks, &out.PostDeploymentTasks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PreDeploymentEvaluations != nil {
		in, out := &in.PreDeploymentEvaluations, &out.PreDeploymentEvaluations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PostDeploymentEvaluations != nil {
		in, out := &in.PostDeploymentEvaluations, &out.PostDeploymentEvaluations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.ResourceReference = in.ResourceReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnWorkloadSpec.
func (in *KeptnWorkloadSpec) DeepCopy() *KeptnWorkloadSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnWorkloadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnWorkloadStatus) DeepCopyInto(out *KeptnWorkloadStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnWorkloadStatus.
func (in *KeptnWorkloadStatus) DeepCopy() *KeptnWorkloadStatus {
	if in == nil {
		return nil
	}
	out := new(KeptnWorkloadStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnWorkloadVersion) DeepCopyInto(out *KeptnWorkloadVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnWorkloadVersion.
func (in *KeptnWorkloadVersion) DeepCopy() *KeptnWorkloadVersion {
	if in == nil {
		return nil
	}
	out := new(KeptnWorkloadVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnWorkloadVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnWorkloadVersionList) DeepCopyInto(out *KeptnWorkloadVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnWorkloadVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnWorkloadVersionList.
func (in *KeptnWorkloadVersionList) DeepCopy() *KeptnWorkloadVersionList {
	if in == nil {
		return nil
	}
	out := new(KeptnWorkloadVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnWorkloadVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnWorkloadVersionSpec) DeepCopyInto(out *KeptnWorkloadVersionSpec) {
	*out = *in
	in.KeptnWorkloadSpec.DeepCopyInto(&out.KeptnWorkloadSpec)
	if in.TraceId != nil {
		in, out := &in.TraceId, &out.TraceId
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnWorkloadVersionSpec.
func (in *KeptnWorkloadVersionSpec) DeepCopy() *KeptnWorkloadVersionSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnWorkloadVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnWorkloadVersionStatus) DeepCopyInto(out *KeptnWorkloadVersionStatus) {
	*out = *in
	if in.PreDeploymentTaskStatus != nil {
		in, out := &in.PreDeploymentTaskStatus, &out.PreDeploymentTaskStatus
		*out = make([]TaskStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PostDeploymentTaskStatus != nil {
		in, out := &in.PostDeploymentTaskStatus, &out.PostDeploymentTaskStatus
		*out = make([]TaskStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PreDeploymentEvaluationTaskStatus != nil {
		in, out := &in.PreDeploymentEvaluationTaskStatus, &out.PreDeploymentEvaluationTaskStatus
		*out = make([]EvaluationStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PostDeploymentEvaluationTaskStatus != nil {
		in, out := &in.PostDeploymentEvaluationTaskStatus, &out.PostDeploymentEvaluationTaskStatus
		*out = make([]EvaluationStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	if in.PhaseTraceIDs != nil {
		in, out := &in.PhaseTraceIDs, &out.PhaseTraceIDs
		*out = make(common.PhaseTraceID, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(propagation.MapCarrier, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnWorkloadVersionStatus.
func (in *KeptnWorkloadVersionStatus) DeepCopy() *KeptnWorkloadVersionStatus {
	if in == nil {
		return nil
	}
	out := new(KeptnWorkloadVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Objective) DeepCopyInto(out *Objective) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Objective.
func (in *Objective) DeepCopy() *Objective {
	if in == nil {
		return nil
	}
	out := new(Objective)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceReference) DeepCopyInto(out *ResourceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceReference.
func (in *ResourceReference) DeepCopy() *ResourceReference {
	if in == nil {
		return nil
	}
	out := new(ResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecureParameters) DeepCopyInto(out *SecureParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecureParameters.
func (in *SecureParameters) DeepCopy() *SecureParameters {
	if in == nil {
		return nil
	}
	out := new(SecureParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskContext) DeepCopyInto(out *TaskContext) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskContext.
func (in *TaskContext) DeepCopy() *TaskContext {
	if in == nil {
		return nil
	}
	out := new(TaskContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskParameters) DeepCopyInto(out *TaskParameters) {
	*out = *in
	if in.Inline != nil {
		in, out := &in.Inline, &out.Inline
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskParameters.
func (in *TaskParameters) DeepCopy() *TaskParameters {
	if in == nil {
		return nil
	}
	out := new(TaskParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskStatus) DeepCopyInto(out *TaskStatus) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskStatus.
func (in *TaskStatus) DeepCopy() *TaskStatus {
	if in == nil {
		return nil
	}
	out := new(TaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadStatus) DeepCopyInto(out *WorkloadStatus) {
	*out = *in
	out.Workload = in.Workload
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadStatus.
func (in *WorkloadStatus) DeepCopy() *WorkloadStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadStatus)
	in.DeepCopyInto(out)
	return out
}
