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
	"fmt"
	"time"

	"github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha1/common"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KeptnWorkloadVersionSpec defines the desired state of KeptnWorkloadVersion
type KeptnWorkloadVersionSpec struct {
	KeptnWorkloadSpec `json:",inline"`
	WorkloadName      string            `json:"workloadName"`
	PreviousVersion   string            `json:"previousVersion,omitempty"`
	TraceId           map[string]string `json:"traceId,omitempty"`
}

// KeptnWorkloadVersionStatus defines the observed state of KeptnWorkloadVersion
type KeptnWorkloadVersionStatus struct {
	// +kubebuilder:default:=Pending
	PreDeploymentStatus common.KeptnState `json:"preDeploymentStatus,omitempty"`
	// +kubebuilder:default:=Pending
	DeploymentStatus common.KeptnState `json:"deploymentStatus,omitempty"`
	// +kubebuilder:default:=Pending
	PreDeploymentEvaluationStatus common.KeptnState `json:"preDeploymentEvaluationStatus,omitempty"`
	// +kubebuilder:default:=Pending
	PostDeploymentEvaluationStatus common.KeptnState `json:"postDeploymentEvaluationStatus,omitempty"`
	// +kubebuilder:default:=Pending
	PostDeploymentStatus               common.KeptnState   `json:"postDeploymentStatus,omitempty"`
	PreDeploymentTaskStatus            []TaskStatus        `json:"preDeploymentTaskStatus,omitempty"`
	PostDeploymentTaskStatus           []TaskStatus        `json:"postDeploymentTaskStatus,omitempty"`
	PreDeploymentEvaluationTaskStatus  []EvaluationStatus  `json:"preDeploymentEvaluationTaskStatus,omitempty"`
	PostDeploymentEvaluationTaskStatus []EvaluationStatus  `json:"postDeploymentEvaluationTaskStatus,omitempty"`
	StartTime                          metav1.Time         `json:"startTime,omitempty"`
	EndTime                            metav1.Time         `json:"endTime,omitempty"`
	CurrentPhase                       string              `json:"currentPhase,omitempty"`
	PhaseTraceIDs                      common.PhaseTraceID `json:"phaseTraceIDs,omitempty"`
	// +kubebuilder:default:=Pending
	Status common.KeptnState `json:"status,omitempty"`
}

type TaskStatus struct {
	TaskDefinitionName string `json:"taskDefinitionName,omitempty"`
	// +kubebuilder:default:=Pending
	Status    common.KeptnState `json:"status,omitempty"`
	TaskName  string            `json:"taskName,omitempty"`
	StartTime metav1.Time       `json:"startTime,omitempty"`
	EndTime   metav1.Time       `json:"endTime,omitempty"`
}

type EvaluationStatus struct {
	EvaluationDefinitionName string `json:"evaluationDefinitionName,omitempty"`
	// +kubebuilder:default:=Pending
	Status         common.KeptnState `json:"status,omitempty"`
	EvaluationName string            `json:"evaluationName,omitempty"`
	StartTime      metav1.Time       `json:"startTime,omitempty"`
	EndTime        metav1.Time       `json:"endTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=keptnworkloadversions,shortName=kwi
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AppName",type=string,JSONPath=`.spec.app`
// +kubebuilder:printcolumn:name="WorkloadName",type=string,JSONPath=`.spec.workloadName`
// +kubebuilder:printcolumn:name="WorkloadVersion",type=string,JSONPath=`.spec.version`
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.currentPhase`
// +kubebuilder:printcolumn:name="PreDeploymentStatus",priority=1,type=string,JSONPath=`.status.preDeploymentStatus`
// +kubebuilder:printcolumn:name="PreDeploymentEvaluationStatus",priority=1,type=string,JSONPath=`.status.preDeploymentEvaluationStatus`
// +kubebuilder:printcolumn:name="DeploymentStatus",type=string,priority=1,JSONPath=`.status.deploymentStatus`
// +kubebuilder:printcolumn:name="PostDeploymentStatus",type=string,priority=1,JSONPath=`.status.postDeploymentStatus`
// +kubebuilder:printcolumn:name="PostDeploymentEvaluationStatus",priority=1,type=string,JSONPath=`.status.postDeploymentEvaluationStatus`

// KeptnWorkloadVersion is the Schema for the keptnworkloadversions API
type KeptnWorkloadVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeptnWorkloadVersionSpec   `json:"spec,omitempty"`
	Status KeptnWorkloadVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeptnWorkloadVersionList contains a list of KeptnWorkloadVersion
type KeptnWorkloadVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeptnWorkloadVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeptnWorkloadVersion{}, &KeptnWorkloadVersionList{})
}

func (w KeptnWorkloadVersionList) GetItems() []client.Object {
	var b []client.Object
	for _, i := range w.Items {
		b = append(b, &i)
	}
	return b
}

func (w KeptnWorkloadVersion) IsPreDeploymentCompleted() bool {
	return w.Status.PreDeploymentStatus.IsCompleted()
}

func (w KeptnWorkloadVersion) IsPreDeploymentEvaluationCompleted() bool {
	return w.Status.PreDeploymentEvaluationStatus.IsCompleted()
}

func (w KeptnWorkloadVersion) IsPreDeploymentSucceeded() bool {
	return w.Status.PreDeploymentStatus.IsSucceeded()
}

func (w KeptnWorkloadVersion) IsPreDeploymentFailed() bool {
	return w.Status.PreDeploymentStatus.IsFailed()
}

func (w KeptnWorkloadVersion) IsPreDeploymentEvaluationSucceeded() bool {
	return w.Status.PreDeploymentEvaluationStatus.IsSucceeded()
}

func (w KeptnWorkloadVersion) IsPreDeploymentEvaluationFailed() bool {
	return w.Status.PreDeploymentEvaluationStatus.IsFailed()
}

func (w KeptnWorkloadVersion) IsPostDeploymentCompleted() bool {
	return w.Status.PostDeploymentStatus.IsCompleted()
}

func (w KeptnWorkloadVersion) IsPostDeploymentEvaluationCompleted() bool {
	return w.Status.PostDeploymentEvaluationStatus.IsCompleted()
}

func (w KeptnWorkloadVersion) IsPostDeploymentSucceeded() bool {
	return w.Status.PostDeploymentStatus.IsSucceeded()
}

func (w KeptnWorkloadVersion) IsPostDeploymentFailed() bool {
	return w.Status.PostDeploymentStatus.IsFailed()
}

func (w KeptnWorkloadVersion) IsPostDeploymentEvaluationSucceeded() bool {
	return w.Status.PostDeploymentEvaluationStatus.IsSucceeded()
}

func (w KeptnWorkloadVersion) IsPostDeploymentEvaluationFailed() bool {
	return w.Status.PostDeploymentEvaluationStatus.IsFailed()
}

func (w KeptnWorkloadVersion) IsDeploymentCompleted() bool {
	return w.Status.DeploymentStatus.IsCompleted()
}

func (w KeptnWorkloadVersion) IsDeploymentSucceeded() bool {
	return w.Status.DeploymentStatus.IsSucceeded()
}

func (w KeptnWorkloadVersion) IsDeploymentFailed() bool {
	return w.Status.DeploymentStatus.IsFailed()
}

func (w *KeptnWorkloadVersion) SetStartTime() {
	if w.Status.StartTime.IsZero() {
		w.Status.StartTime = metav1.NewTime(time.Now().UTC())
	}
}

func (w *KeptnWorkloadVersion) SetEndTime() {
	if w.Status.EndTime.IsZero() {
		w.Status.EndTime = metav1.NewTime(time.Now().UTC())
	}
}

func (w *KeptnWorkloadVersion) IsStartTimeSet() bool {
	return !w.Status.StartTime.IsZero()
}

func (w *KeptnWorkloadVersion) IsEndTimeSet() bool {
	return !w.Status.EndTime.IsZero()
}

func (w KeptnWorkloadVersion) GetStartTime() time.Time {
	return w.Status.StartTime.Time
}

func (w KeptnWorkloadVersion) GetEndTime() time.Time {
	return w.Status.EndTime.Time
}

func (w *KeptnWorkloadVersion) Complete() {
	w.SetEndTime()
}

func (t *TaskStatus) SetStartTime() {
	if t.StartTime.IsZero() {
		t.StartTime = metav1.NewTime(time.Now().UTC())
	}
}

func (t *TaskStatus) SetEndTime() {
	if t.EndTime.IsZero() {
		t.EndTime = metav1.NewTime(time.Now().UTC())
	}
}

func (e *EvaluationStatus) SetStartTime() {
	if e.StartTime.IsZero() {
		e.StartTime = metav1.NewTime(time.Now().UTC())
	}
}

func (e *EvaluationStatus) SetEndTime() {
	if e.EndTime.IsZero() {
		e.EndTime = metav1.NewTime(time.Now().UTC())
	}
}

func (w KeptnWorkloadVersion) GetActiveMetricsAttributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		common.AppName.String(w.Spec.AppName),
		common.WorkloadName.String(w.Spec.WorkloadName),
		common.WorkloadVersion.String(w.Spec.Version),
		common.WorkloadNamespace.String(w.Namespace),
	}
}

func (w KeptnWorkloadVersion) GetMetricsAttributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		common.AppName.String(w.Spec.AppName),
		common.WorkloadName.String(w.Spec.WorkloadName),
		common.WorkloadVersion.String(w.Spec.Version),
		common.WorkloadNamespace.String(w.Namespace),
		common.WorkloadStatus.String(string(w.Status.Status)),
	}
}

func (w KeptnWorkloadVersion) GetDurationMetricsAttributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		common.AppName.String(w.Spec.AppName),
		common.WorkloadName.String(w.Spec.WorkloadName),
		common.WorkloadVersion.String(w.Spec.Version),
		common.WorkloadPreviousVersion.String(w.Spec.PreviousVersion),
	}
}

func (w KeptnWorkloadVersion) GetState() common.KeptnState {
	return w.Status.Status
}

func (w KeptnWorkloadVersion) GetPreDeploymentTasks() []string {
	return w.Spec.PreDeploymentTasks
}

func (w KeptnWorkloadVersion) GetPostDeploymentTasks() []string {
	return w.Spec.PostDeploymentTasks
}

func (w KeptnWorkloadVersion) GetPreDeploymentTaskStatus() []TaskStatus {
	return w.Status.PreDeploymentTaskStatus
}

func (w KeptnWorkloadVersion) GetPostDeploymentTaskStatus() []TaskStatus {
	return w.Status.PostDeploymentTaskStatus
}

func (w KeptnWorkloadVersion) GetPreDeploymentEvaluations() []string {
	return w.Spec.PreDeploymentEvaluations
}

func (w KeptnWorkloadVersion) GetPostDeploymentEvaluations() []string {
	return w.Spec.PostDeploymentEvaluations
}

func (w KeptnWorkloadVersion) GetPreDeploymentEvaluationTaskStatus() []EvaluationStatus {
	return w.Status.PreDeploymentEvaluationTaskStatus
}

func (w KeptnWorkloadVersion) GetPostDeploymentEvaluationTaskStatus() []EvaluationStatus {
	return w.Status.PostDeploymentEvaluationTaskStatus
}

func (w KeptnWorkloadVersion) GetAppName() string {
	return w.Spec.AppName
}

func (w KeptnWorkloadVersion) GetPreviousVersion() string {
	return w.Spec.PreviousVersion
}

func (w KeptnWorkloadVersion) GetParentName() string {
	return w.Spec.WorkloadName
}

func (w KeptnWorkloadVersion) GetNamespace() string {
	return w.Namespace
}

func (w *KeptnWorkloadVersion) SetState(state common.KeptnState) {
	w.Status.Status = state
}

func (w KeptnWorkloadVersion) GetCurrentPhase() string {
	return w.Status.CurrentPhase
}

func (w *KeptnWorkloadVersion) SetCurrentPhase(phase string) {
	w.Status.CurrentPhase = phase
}

func (w KeptnWorkloadVersion) GetVersion() string {
	return w.Spec.Version
}

func (w KeptnWorkloadVersion) GenerateTask(taskDefinition string, checkType common.CheckType) KeptnTask {
	return KeptnTask{
		ObjectMeta: metav1.ObjectMeta{
			Name:      common.GenerateTaskName(checkType, taskDefinition),
			Namespace: w.Namespace,
		},
		Spec: KeptnTaskSpec{
			AppName:          w.GetAppName(),
			WorkloadVersion:  w.GetVersion(),
			Workload:         w.GetParentName(),
			TaskDefinition:   taskDefinition,
			Parameters:       TaskParameters{},
			SecureParameters: SecureParameters{},
			Type:             checkType,
		},
	}
}

func (w KeptnWorkloadVersion) GenerateEvaluation(evaluationDefinition string, checkType common.CheckType) KeptnEvaluation {
	return KeptnEvaluation{
		ObjectMeta: metav1.ObjectMeta{
			Name:      common.GenerateEvaluationName(checkType, evaluationDefinition),
			Namespace: w.Namespace,
		},
		Spec: KeptnEvaluationSpec{
			AppName:              w.GetAppName(),
			WorkloadVersion:      w.GetVersion(),
			Workload:             w.GetParentName(),
			EvaluationDefinition: evaluationDefinition,
			Type:                 checkType,
			RetryInterval: metav1.Duration{
				Duration: 5 * time.Second,
			},
		},
	}
}

func (w KeptnWorkloadVersion) GetSpanAttributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		common.AppName.String(w.Spec.AppName),
		common.WorkloadName.String(w.Spec.WorkloadName),
		common.WorkloadVersion.String(w.Spec.Version),
		common.WorkloadNamespace.String(w.Namespace),
	}
}

func (w KeptnWorkloadVersion) GetSpanKey(phase string) string {
	return fmt.Sprintf("%s.%s.%s.%s", w.Spec.TraceId["traceparent"], w.Spec.WorkloadName, w.Spec.Version, phase)
}

func (w KeptnWorkloadVersion) GetSpanName(phase string) string {
	if phase == "" {
		return w.Name
	}
	return fmt.Sprintf("%s/%s", w.Spec.WorkloadName, phase)
}

func (w KeptnWorkloadVersion) SetSpanAttributes(span trace.Span) {
	span.SetAttributes(w.GetSpanAttributes()...)
}

func (w *KeptnWorkloadVersion) CancelRemainingPhases(phase common.KeptnPhaseType) {
	// no need to cancel anything when post-eval tasks fail
	if phase == common.PhaseWorkloadPostEvaluation {
		return
	}
	// cancel everything if app pre-eval tasks have failed
	if phase == common.PhaseAppPreEvaluation {
		w.Status.PreDeploymentStatus = common.StateCancelled
		w.Status.PreDeploymentEvaluationStatus = common.StateCancelled
	}
	// cancel deployment and post-deployment tasks if workload pre-eval tasks have failed
	w.Status.DeploymentStatus = common.StateCancelled
	w.Status.PostDeploymentStatus = common.StateCancelled
	w.Status.PostDeploymentEvaluationStatus = common.StateCancelled
	w.Status.Status = common.StateFailed
}

func (w *KeptnWorkloadVersion) SetPhaseTraceID(phase string, carrier propagation.MapCarrier) {
	if w.Status.PhaseTraceIDs == nil {
		w.Status.PhaseTraceIDs = common.PhaseTraceID{}
	}
	w.Status.PhaseTraceIDs[common.GetShortPhaseName(phase)] = carrier
}
