package keptnworkloadinstance

import (
	"context"
	"github.com/go-logr/logr"
	klcv1alpha1 "github.com/keptn/lifecycle-toolkit/operator/api/v1alpha1"
	keptncommon "github.com/keptn/lifecycle-toolkit/operator/api/v1alpha1/common"
	controllercommon "github.com/keptn/lifecycle-toolkit/operator/controllers/common"
	utils "github.com/keptn/lifecycle-toolkit/operator/controllers/common"
	"github.com/keptn/lifecycle-toolkit/operator/controllers/common/fake"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	testrequire "github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	k8sfake "sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"testing"
)

func TestKeptnWorkloadInstanceReconciler_reconcileDeployment_FailedReplicaSet(t *testing.T) {

	rep := int32(1)
	replicasetFail := makeReplicaSet("myrep", "default", &rep, 0)

	fakeClient := k8sfake.NewClientBuilder().WithObjects(replicasetFail).Build()

	err := klcv1alpha1.AddToScheme(fakeClient.Scheme())
	testrequire.Nil(t, err)

	workloadInstance := makeWorkloadInstanceWithRef(replicasetFail.ObjectMeta, "ReplicaSet")

	err = fakeClient.Create(context.TODO(), workloadInstance)
	require.Nil(t, err)

	r := &KeptnWorkloadInstanceReconciler{
		Client: fakeClient,
	}

	keptnState, err := r.reconcileDeployment(context.TODO(), workloadInstance)
	testrequire.Nil(t, err)
	testrequire.Equal(t, keptncommon.StateProgressing, keptnState)
}

func makeWorkloadInstanceWithRef(objectMeta metav1.ObjectMeta, refKind string) *klcv1alpha1.KeptnWorkloadInstance {
	workloadInstance := &klcv1alpha1.KeptnWorkloadInstance{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-wli",
			Namespace: "default",
		},
		Spec: klcv1alpha1.KeptnWorkloadInstanceSpec{
			KeptnWorkloadSpec: klcv1alpha1.KeptnWorkloadSpec{
				ResourceReference: klcv1alpha1.ResourceReference{
					UID:  objectMeta.UID,
					Name: objectMeta.Name,
					Kind: refKind,
				},
			},
		},
	}
	return workloadInstance
}

func TestKeptnWorkloadInstanceReconciler_reconcileDeployment_FailedStatefulSet(t *testing.T) {

	rep := int32(1)
	statefulsetFail := makeStatefulSet("mystat", "default", &rep, 0)

	fakeClient := k8sfake.NewClientBuilder().WithObjects(statefulsetFail).Build()

	err := klcv1alpha1.AddToScheme(fakeClient.Scheme())
	testrequire.Nil(t, err)

	workloadInstance := makeWorkloadInstanceWithRef(statefulsetFail.ObjectMeta, "StatefulSet")

	err = fakeClient.Create(context.TODO(), workloadInstance)
	require.Nil(t, err)

	r := &KeptnWorkloadInstanceReconciler{
		Client: fakeClient,
	}

	keptnState, err := r.reconcileDeployment(context.TODO(), workloadInstance)
	testrequire.Nil(t, err)
	testrequire.Equal(t, keptncommon.StateProgressing, keptnState)
}

func TestKeptnWorkloadInstanceReconciler_reconcileDeployment_FailedDaemonSet(t *testing.T) {

	daemonSetFail := makeDaemonSet("mystat", "default", 1, 0)

	fakeClient := k8sfake.NewClientBuilder().WithObjects(daemonSetFail).Build()

	err := klcv1alpha1.AddToScheme(fakeClient.Scheme())
	testrequire.Nil(t, err)

	workloadInstance := makeWorkloadInstanceWithRef(daemonSetFail.ObjectMeta, "DaemonSet")

	err = fakeClient.Create(context.TODO(), workloadInstance)
	require.Nil(t, err)

	r := &KeptnWorkloadInstanceReconciler{
		Client: fakeClient,
	}

	keptnState, err := r.reconcileDeployment(context.TODO(), workloadInstance)
	testrequire.Nil(t, err)
	testrequire.Equal(t, keptncommon.StateProgressing, keptnState)
}

func TestKeptnWorkloadInstanceReconciler_reconcileDeployment_ReadyReplicaSet(t *testing.T) {

	rep := int32(1)
	replicaSet := makeReplicaSet("myrep", "default", &rep, 1)

	fakeClient := k8sfake.NewClientBuilder().WithObjects(replicaSet).Build()

	err := klcv1alpha1.AddToScheme(fakeClient.Scheme())
	testrequire.Nil(t, err)

	workloadInstance := makeWorkloadInstanceWithRef(replicaSet.ObjectMeta, "ReplicaSet")

	err = fakeClient.Create(context.TODO(), workloadInstance)
	require.Nil(t, err)

	r := &KeptnWorkloadInstanceReconciler{
		Client: fakeClient,
	}

	keptnState, err := r.reconcileDeployment(context.TODO(), workloadInstance)
	testrequire.Nil(t, err)
	testrequire.Equal(t, keptncommon.StateSucceeded, keptnState)
}

func TestKeptnWorkloadInstanceReconciler_reconcileDeployment_ReadyStatefulSet(t *testing.T) {

	rep := int32(1)
	statefulSet := makeStatefulSet("mystat", "default", &rep, 1)

	fakeClient := k8sfake.NewClientBuilder().WithObjects(statefulSet).Build()

	err := klcv1alpha1.AddToScheme(fakeClient.Scheme())
	testrequire.Nil(t, err)

	workloadInstance := makeWorkloadInstanceWithRef(statefulSet.ObjectMeta, "StatefulSet")

	err = fakeClient.Create(context.TODO(), workloadInstance)
	require.Nil(t, err)

	r := &KeptnWorkloadInstanceReconciler{
		Client: fakeClient,
	}

	keptnState, err := r.reconcileDeployment(context.TODO(), workloadInstance)
	testrequire.Nil(t, err)
	testrequire.Equal(t, keptncommon.StateSucceeded, keptnState)
}

func TestKeptnWorkloadInstanceReconciler_reconcileDeployment_ReadyDaemonSet(t *testing.T) {

	daemonSet := makeDaemonSet("mystat", "default", 1, 1)

	fakeClient := k8sfake.NewClientBuilder().WithObjects(daemonSet).Build()

	err := klcv1alpha1.AddToScheme(fakeClient.Scheme())
	testrequire.Nil(t, err)

	workloadInstance := makeWorkloadInstanceWithRef(daemonSet.ObjectMeta, "DaemonSet")

	err = fakeClient.Create(context.TODO(), workloadInstance)
	require.Nil(t, err)

	r := &KeptnWorkloadInstanceReconciler{
		Client: fakeClient,
	}

	keptnState, err := r.reconcileDeployment(context.TODO(), workloadInstance)
	testrequire.Nil(t, err)
	testrequire.Equal(t, keptncommon.StateSucceeded, keptnState)
}

func TestKeptnWorkloadInstanceReconciler_IsPodRunning(t *testing.T) {
	p1 := makeNominatedPod("pod1", "node1", v1.PodRunning)
	p2 := makeNominatedPod("pod2", "node1", v1.PodPending)
	podList := &v1.PodList{Items: []v1.Pod{p1, p2}}
	podList2 := &v1.PodList{Items: []v1.Pod{p2}}
	r := &KeptnWorkloadInstanceReconciler{
		Client: k8sfake.NewClientBuilder().WithLists(podList).Build(),
	}
	isPodRunning, err := r.isPodRunning(context.TODO(), klcv1alpha1.ResourceReference{UID: "pod1"}, "node1")
	testrequire.Nil(t, err)
	if !isPodRunning {
		t.Errorf("Wrong!")
	}

	r2 := &KeptnWorkloadInstanceReconciler{
		Client: k8sfake.NewClientBuilder().WithLists(podList2).Build(),
	}
	isPodRunning, err = r2.isPodRunning(context.TODO(), klcv1alpha1.ResourceReference{UID: "pod1"}, "node1")
	testrequire.Nil(t, err)
	if isPodRunning {
		t.Errorf("Wrong!")
	}

}

func makeNominatedPod(podName string, nodeName string, phase v1.PodPhase) v1.Pod {
	return v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: nodeName,
			Name:      podName,
			UID:       types.UID(podName),
		},
		Status: v1.PodStatus{
			Phase:             phase,
			NominatedNodeName: nodeName,
		},
	}
}

func makeReplicaSet(name string, namespace string, wanted *int32, available int32) *appsv1.ReplicaSet {

	return &appsv1.ReplicaSet{
		TypeMeta: metav1.TypeMeta{
			Kind: "ReplicaSet",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			UID:       types.UID(name),
		},
		Spec: appsv1.ReplicaSetSpec{
			Replicas: wanted,
		},
		Status: appsv1.ReplicaSetStatus{
			AvailableReplicas: available,
		},
	}

}

func makeStatefulSet(name string, namespace string, wanted *int32, available int32) *appsv1.StatefulSet {

	return &appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind: "StatefulSet",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			UID:       types.UID(name),
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: wanted,
		},
		Status: appsv1.StatefulSetStatus{
			AvailableReplicas: available,
		},
	}

}

func makeDaemonSet(name string, namespace string, wanted int32, available int32) *appsv1.DaemonSet {

	return &appsv1.DaemonSet{
		TypeMeta: metav1.TypeMeta{
			Kind: "StatefulSet",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			UID:       types.UID(name),
		},
		Spec: appsv1.DaemonSetSpec{},
		Status: appsv1.DaemonSetStatus{
			DesiredNumberScheduled: wanted,
			NumberReady:            available,
		},
	}

}

func Test_getLatestAppVersion(t *testing.T) {
	type args struct {
		apps *klcv1alpha1.KeptnAppVersionList
		wli  *klcv1alpha1.KeptnWorkloadInstance
	}
	tests := []struct {
		name           string
		args           args
		wantFound      bool
		wantAppVersion klcv1alpha1.KeptnAppVersion
		wantErr        bool
	}{
		{
			name: "app version found",
			args: args{
				apps: &klcv1alpha1.KeptnAppVersionList{
					Items: []klcv1alpha1.KeptnAppVersion{
						{
							ObjectMeta: metav1.ObjectMeta{
								Name:      "my-app",
								Namespace: "default",
							},
							Spec: klcv1alpha1.KeptnAppVersionSpec{
								KeptnAppSpec: klcv1alpha1.KeptnAppSpec{
									Version: "1.0",
									Workloads: []klcv1alpha1.KeptnWorkloadRef{
										{
											Name:    "my-workload",
											Version: "1.0",
										},
									},
								},
								AppName: "my-app",
							},
						},
						{
							ObjectMeta: metav1.ObjectMeta{
								Name:      "my-app",
								Namespace: "default",
							},
							Spec: klcv1alpha1.KeptnAppVersionSpec{
								KeptnAppSpec: klcv1alpha1.KeptnAppSpec{
									Version: "2.0",
									Workloads: []klcv1alpha1.KeptnWorkloadRef{
										{
											Name:    "my-workload",
											Version: "1.0",
										},
									},
								},
								AppName: "my-app",
							},
						},
					},
				},
				wli: &klcv1alpha1.KeptnWorkloadInstance{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-workloadinstance",
						Namespace: "default",
					},
					Spec: klcv1alpha1.KeptnWorkloadInstanceSpec{
						KeptnWorkloadSpec: klcv1alpha1.KeptnWorkloadSpec{
							AppName: "my-app",
							Version: "1.0",
						},
						WorkloadName: "my-app-my-workload",
					},
				},
			},
			wantFound: true,
			wantAppVersion: klcv1alpha1.KeptnAppVersion{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-app",
					Namespace: "default",
				},
				Spec: klcv1alpha1.KeptnAppVersionSpec{
					KeptnAppSpec: klcv1alpha1.KeptnAppSpec{
						Version: "2.0",
						Workloads: []klcv1alpha1.KeptnWorkloadRef{
							{
								Name:    "my-workload",
								Version: "1.0",
							},
						},
					},
					AppName: "my-app",
				},
			},
			wantErr: false,
		},
		{
			name: "app version not found",
			args: args{
				apps: &klcv1alpha1.KeptnAppVersionList{
					Items: []klcv1alpha1.KeptnAppVersion{
						{
							ObjectMeta: metav1.ObjectMeta{
								Name:      "my-app",
								Namespace: "default",
							},
							Spec: klcv1alpha1.KeptnAppVersionSpec{
								KeptnAppSpec: klcv1alpha1.KeptnAppSpec{
									Version: "1.0",
									Workloads: []klcv1alpha1.KeptnWorkloadRef{
										{
											Name:    "my-other-workload",
											Version: "1.0",
										},
									},
								},
								AppName: "my-app",
							},
						},
					},
				},
				wli: &klcv1alpha1.KeptnWorkloadInstance{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-workloadinstance",
						Namespace: "default",
					},
					Spec: klcv1alpha1.KeptnWorkloadInstanceSpec{
						KeptnWorkloadSpec: klcv1alpha1.KeptnWorkloadSpec{
							AppName: "my-app",
							Version: "1.0",
						},
						WorkloadName: "my-app-my-workload",
					},
				},
			},
			wantFound:      false,
			wantAppVersion: klcv1alpha1.KeptnAppVersion{},
			wantErr:        false,
		},
		{
			name: "app version with invalid version",
			args: args{
				apps: &klcv1alpha1.KeptnAppVersionList{
					Items: []klcv1alpha1.KeptnAppVersion{
						{
							ObjectMeta: metav1.ObjectMeta{
								Name:      "my-app",
								Namespace: "default",
							},
							Spec: klcv1alpha1.KeptnAppVersionSpec{
								KeptnAppSpec: klcv1alpha1.KeptnAppSpec{
									Version: "",
									Workloads: []klcv1alpha1.KeptnWorkloadRef{
										{
											Name:    "my-workload",
											Version: "1.0",
										},
									},
								},
								AppName: "my-app",
							},
						},
					},
				},
				wli: &klcv1alpha1.KeptnWorkloadInstance{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-workloadinstance",
						Namespace: "default",
					},
					Spec: klcv1alpha1.KeptnWorkloadInstanceSpec{
						KeptnWorkloadSpec: klcv1alpha1.KeptnWorkloadSpec{
							AppName: "my-app",
							Version: "1.0",
						},
						WorkloadName: "my-app-my-workload",
					},
				},
			},
			wantFound:      false,
			wantAppVersion: klcv1alpha1.KeptnAppVersion{},
			wantErr:        true,
		},
		{
			name: "app version list empty",
			args: args{
				apps: &klcv1alpha1.KeptnAppVersionList{
					Items: []klcv1alpha1.KeptnAppVersion{},
				},
				wli: &klcv1alpha1.KeptnWorkloadInstance{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-workloadinstance",
						Namespace: "default",
					},
					Spec: klcv1alpha1.KeptnWorkloadInstanceSpec{
						KeptnWorkloadSpec: klcv1alpha1.KeptnWorkloadSpec{
							AppName: "my-app",
							Version: "1.0",
						},
						WorkloadName: "my-app-my-workload",
					},
				},
			},
			wantFound:      false,
			wantAppVersion: klcv1alpha1.KeptnAppVersion{},
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, gotAppVersion, err := getLatestAppVersion(tt.args.apps, tt.args.wli)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.wantFound, found)
			require.Equal(t, tt.wantAppVersion, gotAppVersion)
		})
	}
}

func TestKeptnWorkloadInstanceReconciler_Reconcile(t *testing.T) {
	testNamespace := "some-ns"
	r, eventChannel, _ := setupReconciler(t)

	type fields struct {
		Client      client.Client
		Scheme      *runtime.Scheme
		Recorder    record.EventRecorder
		Log         logr.Logger
		Meters      keptncommon.KeptnMeters
		Tracer      trace.Tracer
		SpanHandler *controllercommon.SpanHandler
	}
	type args struct {
		ctx context.Context
		req ctrl.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		event   string
		wantErr bool
	}{
		{
			name:    "Test that nothing happens when nothing is there to reconcile",
			event:   "",
			wantErr: false,
		},
		{
			name: "Test fetching of data while pre deployment evaluations are not done",
			args: args{
				ctx: context.TODO(),
				req: ctrl.Request{
					NamespacedName: types.NamespacedName{
						Namespace: testNamespace,
						Name:      "some-wi",
					},
				},
			},
			event:   "Normal AppPreDeployEvaluationsNotFinished App Pre-Deployment Evaluations Pre evaluations tasks for app not finished / Namespace: some-ns, Name: some-wi, Version: 1.0.0",
			wantErr: false,
		},
	}

	utils.AddWorkloadInstance(r.Client, "some-wi", testNamespace)
	utils.AddApp(r.Client, "some-app")
	utils.AddAppVersion(
		r.Client,
		testNamespace,
		"some-app",
		"1.0.0",
		[]klcv1alpha1.KeptnWorkloadRef{
			{
				Name:    "some-workload",
				Version: "1.0.0",
			},
		},
		klcv1alpha1.KeptnAppVersionStatus{
			PreDeploymentStatus:                keptncommon.StateSucceeded,
			PostDeploymentStatus:               keptncommon.StateSucceeded,
			PreDeploymentEvaluationStatus:      keptncommon.StateSucceeded,
			PostDeploymentEvaluationStatus:     keptncommon.StateSucceeded,
			WorkloadOverallStatus:              keptncommon.StateSucceeded,
			WorkloadStatus:                     nil,
			CurrentPhase:                       keptncommon.PhaseWorkloadPostEvaluation.ShortName,
			PreDeploymentTaskStatus:            nil,
			PostDeploymentTaskStatus:           nil,
			PreDeploymentEvaluationTaskStatus:  nil,
			PostDeploymentEvaluationTaskStatus: nil,
			Status:                             keptncommon.StateSucceeded,
			StartTime:                          metav1.Time{},
			EndTime:                            metav1.Time{},
		},
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := r.Reconcile(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reconcile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.event != "" {
				event := <-eventChannel
				assert.Matches(t, event, tt.event)
			}
		})
	}
}

func setupReconciler(t *testing.T) (*KeptnWorkloadInstanceReconciler, chan string, *fake.ITracerMock) {
	//setup logger
	opts := zap.Options{
		Development: true,
	}
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	//fake a tracer
	tr := &fake.ITracerMock{StartFunc: func(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
		return ctx, trace.SpanFromContext(ctx)
	}}

	fakeClient, err := fake.NewClient()
	if err != nil {
		t.Errorf("Reconcile() error when setting up fake client %v", err)
	}
	recorder := record.NewFakeRecorder(100)
	r := &KeptnWorkloadInstanceReconciler{
		Client:   fakeClient,
		Scheme:   scheme.Scheme,
		Recorder: recorder,
		Log:      ctrl.Log.WithName("test-appController"),
		Tracer:   tr,
	}
	return r, recorder.Events, tr
}
