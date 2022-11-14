package api_test

import (
	"testing"
	"time"

	"github.com/keptn/lifecycle-toolkit/operator/api/v1alpha1"
	"github.com/keptn/lifecycle-toolkit/operator/api/v1alpha1/common"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestKeptnWorkloadInstance(t *testing.T) {
	workload := &v1alpha1.KeptnWorkloadInstance{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "workload",
			Namespace: "namespace",
		},
		Status: v1alpha1.KeptnWorkloadInstanceStatus{
			PreDeploymentStatus:            common.StateFailed,
			PreDeploymentEvaluationStatus:  common.StateFailed,
			PostDeploymentStatus:           common.StateFailed,
			PostDeploymentEvaluationStatus: common.StateFailed,
			DeploymentStatus:               common.StateFailed,
			Status:                         common.StateFailed,
			PreDeploymentTaskStatus: []v1alpha1.TaskStatus{
				{
					TaskDefinitionName: "defname",
					Status:             common.StateFailed,
					TaskName:           "taskname",
				},
			},
			PostDeploymentTaskStatus: []v1alpha1.TaskStatus{
				{
					TaskDefinitionName: "defname2",
					Status:             common.StateFailed,
					TaskName:           "taskname2",
				},
			},
			PreDeploymentEvaluationTaskStatus: []v1alpha1.EvaluationStatus{
				{
					EvaluationDefinitionName: "defname3",
					Status:                   common.StateFailed,
					EvaluationName:           "taskname3",
				},
			},
			PostDeploymentEvaluationTaskStatus: []v1alpha1.EvaluationStatus{
				{
					EvaluationDefinitionName: "defname4",
					Status:                   common.StateFailed,
					EvaluationName:           "taskname4",
				},
			},
			CurrentPhase: common.PhaseAppDeployment.ShortName,
		},
		Spec: v1alpha1.KeptnWorkloadInstanceSpec{
			KeptnWorkloadSpec: v1alpha1.KeptnWorkloadSpec{
				PreDeploymentTasks:        []string{"task1", "task2"},
				PostDeploymentTasks:       []string{"task3", "task4"},
				PreDeploymentEvaluations:  []string{"task5", "task6"},
				PostDeploymentEvaluations: []string{"task7", "task8"},
				Version:                   "version",
				AppName:                   "appname",
			},
			PreviousVersion: "prev",
			WorkloadName:    "workloadname",
			TraceId:         map[string]string{"traceparent": "trace1"},
		},
	}

	require.True(t, workload.IsPreDeploymentCompleted())
	require.False(t, workload.IsPreDeploymentSucceeded())
	require.True(t, workload.IsPreDeploymentFailed())

	require.True(t, workload.IsPreDeploymentEvaluationCompleted())
	require.False(t, workload.IsPreDeploymentEvaluationSucceeded())
	require.True(t, workload.IsPreDeploymentEvaluationFailed())

	require.True(t, workload.IsPostDeploymentCompleted())
	require.False(t, workload.IsPostDeploymentSucceeded())
	require.True(t, workload.IsPostDeploymentFailed())

	require.True(t, workload.IsPostDeploymentEvaluationCompleted())
	require.False(t, workload.IsPostDeploymentEvaluationSucceeded())
	require.True(t, workload.IsPostDeploymentEvaluationFailed())

	require.True(t, workload.IsDeploymentCompleted())
	require.False(t, workload.IsDeploymentSucceeded())
	require.True(t, workload.IsDeploymentFailed())

	require.False(t, workload.IsEndTimeSet())
	require.False(t, workload.IsStartTimeSet())

	workload.SetStartTime()
	workload.SetEndTime()

	require.True(t, workload.IsEndTimeSet())
	require.True(t, workload.IsStartTimeSet())

	require.Equal(t, []attribute.KeyValue{
		common.AppName.String("appname"),
		common.WorkloadName.String("workloadname"),
		common.WorkloadVersion.String("version"),
		common.WorkloadNamespace.String("namespace"),
	}, workload.GetActiveMetricsAttributes())

	require.Equal(t, []attribute.KeyValue{
		common.AppName.String("appname"),
		common.WorkloadName.String("workloadname"),
		common.WorkloadVersion.String("version"),
		common.WorkloadNamespace.String("namespace"),
		common.WorkloadStatus.String(string(common.StateFailed)),
	}, workload.GetMetricsAttributes())

	require.Equal(t, []attribute.KeyValue{
		common.AppName.String("appname"),
		common.WorkloadName.String("workloadname"),
		common.WorkloadVersion.String("version"),
		common.WorkloadPreviousVersion.String("prev"),
	}, workload.GetDurationMetricsAttributes())

	require.Equal(t, common.StateFailed, workload.GetState())

	require.Equal(t, []string{"task1", "task2"}, workload.GetPreDeploymentTasks())
	require.Equal(t, []string{"task3", "task4"}, workload.GetPostDeploymentTasks())
	require.Equal(t, []string{"task5", "task6"}, workload.GetPreDeploymentEvaluations())
	require.Equal(t, []string{"task7", "task8"}, workload.GetPostDeploymentEvaluations())

	require.Equal(t, []v1alpha1.TaskStatus{
		{
			TaskDefinitionName: "defname",
			Status:             common.StateFailed,
			TaskName:           "taskname",
		},
	}, workload.GetPreDeploymentTaskStatus())

	require.Equal(t, []v1alpha1.TaskStatus{
		{
			TaskDefinitionName: "defname2",
			Status:             common.StateFailed,
			TaskName:           "taskname2",
		},
	}, workload.GetPostDeploymentTaskStatus())

	require.Equal(t, []v1alpha1.EvaluationStatus{
		{
			EvaluationDefinitionName: "defname3",
			Status:                   common.StateFailed,
			EvaluationName:           "taskname3",
		},
	}, workload.GetPreDeploymentEvaluationTaskStatus())

	require.Equal(t, []v1alpha1.EvaluationStatus{
		{
			EvaluationDefinitionName: "defname4",
			Status:                   common.StateFailed,
			EvaluationName:           "taskname4",
		},
	}, workload.GetPostDeploymentEvaluationTaskStatus())

	require.Equal(t, "appname", workload.GetAppName())
	require.Equal(t, "prev", workload.GetPreviousVersion())
	require.Equal(t, "workloadname", workload.GetParentName())
	require.Equal(t, "namespace", workload.GetNamespace())

	workload.SetState(common.StatePending)
	require.Equal(t, common.StatePending, workload.GetState())

	require.True(t, !workload.GetStartTime().IsZero())
	require.True(t, !workload.GetEndTime().IsZero())

	workload.SetCurrentPhase(common.PhaseAppDeployment.LongName)
	require.Equal(t, common.PhaseAppDeployment.LongName, workload.GetCurrentPhase())

	workload.Status.EndTime = v1.Time{Time: time.Time{}}
	workload.Complete()
	require.True(t, !workload.GetEndTime().IsZero())

	require.Equal(t, "version", workload.GetVersion())

	require.Equal(t, "trace1.workloadname.version.phase", workload.GetSpanKey("phase"))

	task := workload.GenerateTask(map[string]string{}, "taskdef", common.PostDeploymentCheckType)
	require.Equal(t, v1alpha1.KeptnTaskSpec{
		AppName:          workload.GetAppName(),
		WorkloadVersion:  workload.GetVersion(),
		Workload:         workload.GetParentName(),
		TaskDefinition:   "taskdef",
		Parameters:       v1alpha1.TaskParameters{},
		SecureParameters: v1alpha1.SecureParameters{},
		Type:             common.PostDeploymentCheckType,
	}, task.Spec)

	evaluation := workload.GenerateEvaluation(map[string]string{}, "taskdef", common.PostDeploymentCheckType)
	require.Equal(t, v1alpha1.KeptnEvaluationSpec{
		AppName:              workload.GetAppName(),
		WorkloadVersion:      workload.GetVersion(),
		Workload:             workload.GetParentName(),
		EvaluationDefinition: "taskdef",
		Type:                 common.PostDeploymentCheckType,
		RetryInterval: metav1.Duration{
			Duration: 5 * time.Second,
		},
	}, evaluation.Spec)

	require.Equal(t, "workloadname/phase", workload.GetSpanName("phase"))

	require.Equal(t, []attribute.KeyValue{
		common.AppName.String("appname"),
		common.WorkloadName.String("workloadname"),
		common.WorkloadVersion.String("version"),
		common.WorkloadNamespace.String("namespace"),
	}, workload.GetSpanAttributes())
}
