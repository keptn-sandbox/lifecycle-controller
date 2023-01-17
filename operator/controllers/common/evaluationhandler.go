package common

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	klcv1alpha2 "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2"
	apicommon "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2/common"
	controllererrors "github.com/keptn/lifecycle-toolkit/operator/controllers/errors"
	"github.com/keptn/lifecycle-toolkit/operator/controllers/lifecycle/interfaces"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type EvaluationHandler struct {
	client.Client
	Recorder    record.EventRecorder
	Log         logr.Logger
	Tracer      trace.Tracer
	Scheme      *runtime.Scheme
	SpanHandler ISpanHandler
}

//nolint:gocognit,gocyclo
func (r EvaluationHandler) ReconcileEvaluations(ctx context.Context, phaseCtx context.Context, reconcileObject client.Object, evaluationCreateAttributes CreateAttributes) ([]klcv1alpha2.ItemStatus, apicommon.StatusSummary, error) {
	piWrapper, err := interfaces.NewPhaseItemWrapperFromClientObject(reconcileObject)
	if err != nil {
		return nil, apicommon.StatusSummary{}, err
	}

	evaluations, statuses := r.setupEvaluations(evaluationCreateAttributes, piWrapper)

	var summary apicommon.StatusSummary
	summary.Total = len(evaluations)
	// Check current state of the PrePostEvaluationTasks
	var newStatus []klcv1alpha2.ItemStatus
	for _, evaluationName := range evaluations {
		oldstatus := GetOldStatus(evaluationName, statuses)

		evaluationStatus := GetItemStatus(evaluationName, statuses)
		evaluation := &klcv1alpha2.KeptnEvaluation{}
		evaluationExists := false

		if oldstatus != evaluationStatus.Status {
			RecordEvent(r.Recorder, apicommon.PhaseReconcileEvaluation, "Normal", reconcileObject, "EvaluationStatusChanged", fmt.Sprintf("evaluation status changed from %s to %s", oldstatus, evaluationStatus.Status), piWrapper.GetVersion())
		}

		// Check if evaluation has already succeeded or failed
		if evaluationStatus.Status.IsCompleted() {
			newStatus = append(newStatus, evaluationStatus)
			continue
		}

		// Check if Evaluation is already created
		if evaluationStatus.Name != "" {
			err := r.Client.Get(ctx, types.NamespacedName{Name: evaluationStatus.Name, Namespace: piWrapper.GetNamespace()}, evaluation)
			if err != nil && errors.IsNotFound(err) {
				evaluationStatus.Name = ""
			} else if err != nil {
				return nil, summary, err
			}
			evaluationExists = true
		}

		// Create new Evaluation if it does not exist
		if !evaluationExists {
			err := r.handleEvaluationNotExists(
				ctx,
				phaseCtx,
				evaluationCreateAttributes,
				evaluationName,
				piWrapper,
				reconcileObject,
				evaluation,
				&evaluationStatus,
			)
			if err != nil {
				return nil, summary, err
			}
		} else {
			r.handleEvaluationExists(
				phaseCtx,
				piWrapper,
				evaluation,
				&evaluationStatus,
			)
		}
		// Update state of the Check
		newStatus = append(newStatus, evaluationStatus)
	}

	for _, ns := range newStatus {
		summary = apicommon.UpdateStatusSummary(ns.Status, summary)
	}
	if apicommon.GetOverallState(summary) != apicommon.StateSucceeded {
		RecordEvent(r.Recorder, apicommon.PhaseReconcileEvaluation, "Warning", reconcileObject, "NotFinished", "has not finished", piWrapper.GetVersion())
	}
	return newStatus, summary, nil
}

//nolint:dupl
func (r EvaluationHandler) CreateKeptnEvaluation(ctx context.Context, namespace string, reconcileObject client.Object, evaluationCreateAttributes CreateAttributes) (string, error) {
	piWrapper, err := interfaces.NewPhaseItemWrapperFromClientObject(reconcileObject)
	if err != nil {
		return "", err
	}

	phase := apicommon.PhaseCreateEvaluation

	newEvaluation := piWrapper.GenerateEvaluation(evaluationCreateAttributes.Definition, evaluationCreateAttributes.CheckType)
	err = controllerutil.SetControllerReference(reconcileObject, &newEvaluation, r.Scheme)
	if err != nil {
		r.Log.Error(err, "could not set controller reference:")
	}
	err = r.Client.Create(ctx, &newEvaluation)
	if err != nil {
		r.Log.Error(err, "could not create KeptnEvaluation")
		RecordEvent(r.Recorder, phase, "Warning", reconcileObject, "CreateFailed", "could not create KeptnEvaluation", piWrapper.GetVersion())
		return "", err
	}
	RecordEvent(r.Recorder, phase, "Normal", reconcileObject, "Created", "created", piWrapper.GetVersion())

	return newEvaluation.Name, nil
}

func (r EvaluationHandler) emitEvaluationFailureEvents(evaluation *klcv1alpha2.KeptnEvaluation, spanTrace trace.Span, piWrapper *interfaces.PhaseItemWrapper) {
	k8sEventMessage := "evaluation failed"
	for k, v := range evaluation.Status.EvaluationStatus {
		if v.Status == apicommon.StateFailed {
			msg := fmt.Sprintf("evaluation of '%s' failed with value: '%s' and reason: '%s'", k, v.Value, v.Message)
			spanTrace.AddEvent(msg, trace.WithTimestamp(time.Now().UTC()))
			k8sEventMessage = fmt.Sprintf("%s\n%s", k8sEventMessage, msg)
		}
	}
	RecordEvent(r.Recorder, apicommon.PhaseReconcileEvaluation, "Warning", evaluation, "Failed", k8sEventMessage, piWrapper.GetVersion())
}

func (r EvaluationHandler) setupEvaluations(evaluationCreateAttributes CreateAttributes, piWrapper *interfaces.PhaseItemWrapper) ([]string, []klcv1alpha2.ItemStatus) {
	var evaluations []string
	var statuses []klcv1alpha2.ItemStatus

	switch evaluationCreateAttributes.CheckType {
	case apicommon.PreDeploymentEvaluationCheckType:
		evaluations = piWrapper.GetPreDeploymentEvaluations()
		statuses = piWrapper.GetPreDeploymentEvaluationTaskStatus()
	case apicommon.PostDeploymentEvaluationCheckType:
		evaluations = piWrapper.GetPostDeploymentEvaluations()
		statuses = piWrapper.GetPostDeploymentEvaluationTaskStatus()
	}
	return evaluations, statuses
}

func (r EvaluationHandler) handleEvaluationNotExists(ctx context.Context, phaseCtx context.Context, evaluationCreateAttributes CreateAttributes, evaluationName string, piWrapper *interfaces.PhaseItemWrapper, reconcileObject client.Object, evaluation *klcv1alpha2.KeptnEvaluation, evaluationStatus *klcv1alpha2.ItemStatus) error {
	evaluationCreateAttributes.Definition = evaluationName
	evaluationName, err := r.CreateKeptnEvaluation(ctx, piWrapper.GetNamespace(), reconcileObject, evaluationCreateAttributes)
	if err != nil {
		return err
	}
	evaluationStatus.Name = evaluationName
	evaluationStatus.SetStartTime()
	_, _, err = r.SpanHandler.GetSpan(phaseCtx, r.Tracer, evaluation, "")
	if err != nil {
		r.Log.Error(err, "could not get span")
	}

	return nil
}

func (r EvaluationHandler) handleEvaluationExists(phaseCtx context.Context, piWrapper *interfaces.PhaseItemWrapper, evaluation *klcv1alpha2.KeptnEvaluation, evaluationStatus *klcv1alpha2.ItemStatus) {
	_, spanEvaluationTrace, err := r.SpanHandler.GetSpan(phaseCtx, r.Tracer, evaluation, "")
	if err != nil {
		r.Log.Error(err, "could not get span")
	}
	// Update state of Evaluation if it is already created
	evaluationStatus.Status = evaluation.Status.OverallStatus
	if evaluationStatus.Status.IsCompleted() {
		if evaluationStatus.Status.IsSucceeded() {
			spanEvaluationTrace.AddEvent(evaluation.Name + " has finished")
			spanEvaluationTrace.SetStatus(codes.Ok, "Finished")
			RecordEvent(r.Recorder, apicommon.PhaseReconcileEvaluation, "Normal", evaluation, "Succeeded", "evaluation succeeded", piWrapper.GetVersion())
		} else {
			spanEvaluationTrace.AddEvent(evaluation.Name + " has failed")
			r.emitEvaluationFailureEvents(evaluation, spanEvaluationTrace, piWrapper)
			spanEvaluationTrace.SetStatus(codes.Error, "Failed")
		}
		spanEvaluationTrace.End()
		if err := r.SpanHandler.UnbindSpan(evaluation, ""); err != nil {
			r.Log.Error(err, controllererrors.ErrCouldNotUnbindSpan, evaluation.Name)
		}
		evaluationStatus.SetEndTime()
	}
}
