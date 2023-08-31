//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Analysis) DeepCopyInto(out *Analysis) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Analysis.
func (in *Analysis) DeepCopy() *Analysis {
	if in == nil {
		return nil
	}
	out := new(Analysis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Analysis) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalysisDefinition) DeepCopyInto(out *AnalysisDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalysisDefinition.
func (in *AnalysisDefinition) DeepCopy() *AnalysisDefinition {
	if in == nil {
		return nil
	}
	out := new(AnalysisDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalysisDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalysisDefinitionList) DeepCopyInto(out *AnalysisDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnalysisDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalysisDefinitionList.
func (in *AnalysisDefinitionList) DeepCopy() *AnalysisDefinitionList {
	if in == nil {
		return nil
	}
	out := new(AnalysisDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalysisDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalysisDefinitionSpec) DeepCopyInto(out *AnalysisDefinitionSpec) {
	*out = *in
	if in.Objectives != nil {
		in, out := &in.Objectives, &out.Objectives
		*out = make([]Objective, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.TotalScore = in.TotalScore
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalysisDefinitionSpec.
func (in *AnalysisDefinitionSpec) DeepCopy() *AnalysisDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(AnalysisDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalysisList) DeepCopyInto(out *AnalysisList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Analysis, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalysisList.
func (in *AnalysisList) DeepCopy() *AnalysisList {
	if in == nil {
		return nil
	}
	out := new(AnalysisList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalysisList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalysisSpec) DeepCopyInto(out *AnalysisSpec) {
	*out = *in
	in.Timeframe.DeepCopyInto(&out.Timeframe)
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.AnalysisDefinition = in.AnalysisDefinition
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalysisSpec.
func (in *AnalysisSpec) DeepCopy() *AnalysisSpec {
	if in == nil {
		return nil
	}
	out := new(AnalysisSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalysisValueTemplate) DeepCopyInto(out *AnalysisValueTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalysisValueTemplate.
func (in *AnalysisValueTemplate) DeepCopy() *AnalysisValueTemplate {
	if in == nil {
		return nil
	}
	out := new(AnalysisValueTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalysisValueTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalysisValueTemplateList) DeepCopyInto(out *AnalysisValueTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnalysisValueTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalysisValueTemplateList.
func (in *AnalysisValueTemplateList) DeepCopy() *AnalysisValueTemplateList {
	if in == nil {
		return nil
	}
	out := new(AnalysisValueTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalysisValueTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalysisValueTemplateSpec) DeepCopyInto(out *AnalysisValueTemplateSpec) {
	*out = *in
	out.Provider = in.Provider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalysisValueTemplateSpec.
func (in *AnalysisValueTemplateSpec) DeepCopy() *AnalysisValueTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(AnalysisValueTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetric) DeepCopyInto(out *KeptnMetric) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetric.
func (in *KeptnMetric) DeepCopy() *KeptnMetric {
	if in == nil {
		return nil
	}
	out := new(KeptnMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnMetric) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetricList) DeepCopyInto(out *KeptnMetricList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetricList.
func (in *KeptnMetricList) DeepCopy() *KeptnMetricList {
	if in == nil {
		return nil
	}
	out := new(KeptnMetricList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnMetricList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetricSpec) DeepCopyInto(out *KeptnMetricSpec) {
	*out = *in
	out.Provider = in.Provider
	if in.Range != nil {
		in, out := &in.Range, &out.Range
		*out = new(RangeSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetricSpec.
func (in *KeptnMetricSpec) DeepCopy() *KeptnMetricSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnMetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetricStatus) DeepCopyInto(out *KeptnMetricStatus) {
	*out = *in
	if in.RawValue != nil {
		in, out := &in.RawValue, &out.RawValue
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	in.LastUpdated.DeepCopyInto(&out.LastUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetricStatus.
func (in *KeptnMetricStatus) DeepCopy() *KeptnMetricStatus {
	if in == nil {
		return nil
	}
	out := new(KeptnMetricStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetricsProvider) DeepCopyInto(out *KeptnMetricsProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetricsProvider.
func (in *KeptnMetricsProvider) DeepCopy() *KeptnMetricsProvider {
	if in == nil {
		return nil
	}
	out := new(KeptnMetricsProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnMetricsProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetricsProviderList) DeepCopyInto(out *KeptnMetricsProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnMetricsProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetricsProviderList.
func (in *KeptnMetricsProviderList) DeepCopy() *KeptnMetricsProviderList {
	if in == nil {
		return nil
	}
	out := new(KeptnMetricsProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnMetricsProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetricsProviderSpec) DeepCopyInto(out *KeptnMetricsProviderSpec) {
	*out = *in
	in.SecretKeyRef.DeepCopyInto(&out.SecretKeyRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetricsProviderSpec.
func (in *KeptnMetricsProviderSpec) DeepCopy() *KeptnMetricsProviderSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnMetricsProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Objective) DeepCopyInto(out *Objective) {
	*out = *in
	out.AnalysisValueTemplateRef = in.AnalysisValueTemplateRef
	in.Target.DeepCopyInto(&out.Target)
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
func (in *Operator) DeepCopyInto(out *Operator) {
	*out = *in
	if in.LessThanOrEqual != nil {
		in, out := &in.LessThanOrEqual, &out.LessThanOrEqual
		*out = new(OperatorValue)
		(*in).DeepCopyInto(*out)
	}
	if in.LessThan != nil {
		in, out := &in.LessThan, &out.LessThan
		*out = new(OperatorValue)
		(*in).DeepCopyInto(*out)
	}
	if in.GreaterThan != nil {
		in, out := &in.GreaterThan, &out.GreaterThan
		*out = new(OperatorValue)
		(*in).DeepCopyInto(*out)
	}
	if in.GreaterThanOrEqual != nil {
		in, out := &in.GreaterThanOrEqual, &out.GreaterThanOrEqual
		*out = new(OperatorValue)
		(*in).DeepCopyInto(*out)
	}
	if in.EqualTo != nil {
		in, out := &in.EqualTo, &out.EqualTo
		*out = new(OperatorValue)
		(*in).DeepCopyInto(*out)
	}
	if in.InRange != nil {
		in, out := &in.InRange, &out.InRange
		*out = new(RangeValue)
		(*in).DeepCopyInto(*out)
	}
	if in.NotInRange != nil {
		in, out := &in.NotInRange, &out.NotInRange
		*out = new(RangeValue)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Operator.
func (in *Operator) DeepCopy() *Operator {
	if in == nil {
		return nil
	}
	out := new(Operator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorValue) DeepCopyInto(out *OperatorValue) {
	*out = *in
	out.FixedValue = in.FixedValue.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorValue.
func (in *OperatorValue) DeepCopy() *OperatorValue {
	if in == nil {
		return nil
	}
	out := new(OperatorValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderRef) DeepCopyInto(out *ProviderRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderRef.
func (in *ProviderRef) DeepCopy() *ProviderRef {
	if in == nil {
		return nil
	}
	out := new(ProviderRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RangeSpec) DeepCopyInto(out *RangeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RangeSpec.
func (in *RangeSpec) DeepCopy() *RangeSpec {
	if in == nil {
		return nil
	}
	out := new(RangeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RangeValue) DeepCopyInto(out *RangeValue) {
	*out = *in
	out.LowBound = in.LowBound.DeepCopy()
	out.HighBound = in.HighBound.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RangeValue.
func (in *RangeValue) DeepCopy() *RangeValue {
	if in == nil {
		return nil
	}
	out := new(RangeValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Target) DeepCopyInto(out *Target) {
	*out = *in
	if in.Failure != nil {
		in, out := &in.Failure, &out.Failure
		*out = new(Operator)
		(*in).DeepCopyInto(*out)
	}
	if in.Warning != nil {
		in, out := &in.Warning, &out.Warning
		*out = new(Operator)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Target.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timeframe) DeepCopyInto(out *Timeframe) {
	*out = *in
	in.From.DeepCopyInto(&out.From)
	in.To.DeepCopyInto(&out.To)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timeframe.
func (in *Timeframe) DeepCopy() *Timeframe {
	if in == nil {
		return nil
	}
	out := new(Timeframe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TotalScore) DeepCopyInto(out *TotalScore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TotalScore.
func (in *TotalScore) DeepCopy() *TotalScore {
	if in == nil {
		return nil
	}
	out := new(TotalScore)
	in.DeepCopyInto(out)
	return out
}
