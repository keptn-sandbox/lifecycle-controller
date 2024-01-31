package keptnworkload

import (
	"context"
	"github.com/keptn/lifecycle-toolkit/lifecycle-operator/controllers/common/eventsender"
	"github.com/keptn/lifecycle-toolkit/lifecycle-operator/controllers/common/testcommon"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"testing"

	klcv1beta1 "github.com/keptn/lifecycle-toolkit/lifecycle-operator/apis/lifecycle/v1beta1"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestKeptnWorkloadReconciler_CreateWorkloadVersion(t *testing.T) {

	workload := &klcv1beta1.KeptnWorkload{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-workload",
			Namespace: "my-namespace",
		},
		Spec: klcv1beta1.KeptnWorkloadSpec{
			AppName: "my-app",
			Version: "v1",
			ResourceReference: klcv1beta1.ResourceReference{
				UID:  "id1",
				Kind: "ReplicaSet",
				Name: "my-replica-set",
			},
		},
	}

	expectedWorkloadVersion := &klcv1beta1.KeptnWorkloadVersion{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-workload-v1",
			Namespace: "my-namespace",
		},
		Spec: klcv1beta1.KeptnWorkloadVersionSpec{
			KeptnWorkloadSpec: workload.Spec,
			WorkloadName:      "my-workload",
		},
	}

	r, _ := setupReconciler(workload)

	res, err := r.Reconcile(context.TODO(), ctrl.Request{
		NamespacedName: types.NamespacedName{
			Namespace: workload.Namespace,
			Name:      workload.Name,
		},
	})
	require.Nil(t, err)
	require.False(t, res.Requeue)

	createdWorkloadVersion := &klcv1beta1.KeptnWorkloadVersion{}
	err = r.Client.Get(context.TODO(),
		types.NamespacedName{
			Namespace: expectedWorkloadVersion.Namespace,
			Name:      expectedWorkloadVersion.Name,
		}, createdWorkloadVersion)

	require.Nil(t, err)

	require.Equal(t, expectedWorkloadVersion.Spec, createdWorkloadVersion.Spec)
}

func TestKeptnWorkloadReconciler_UpdateExistingWorkloadVersion(t *testing.T) {

	workload := &klcv1beta1.KeptnWorkload{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-workload",
			Namespace: "my-namespace",
		},
		Spec: klcv1beta1.KeptnWorkloadSpec{
			AppName: "my-app",
			Version: "v1",
			ResourceReference: klcv1beta1.ResourceReference{
				UID:  "id1",
				Kind: "ReplicaSet",
				Name: "my-replica-set",
			},
		},
	}

	workloadVersion := &klcv1beta1.KeptnWorkloadVersion{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-workload-v1",
			Namespace: "my-namespace",
		},
		Spec: klcv1beta1.KeptnWorkloadVersionSpec{
			KeptnWorkloadSpec: klcv1beta1.KeptnWorkloadSpec{
				AppName: "my-app",
				Version: "v1",
				ResourceReference: klcv1beta1.ResourceReference{
					UID:  "id2",
					Kind: "ReplicaSet",
					Name: "another-replica-set",
				},
			},
		},
	}

	r, _ := setupReconciler(workload, workloadVersion)

	res, err := r.Reconcile(context.TODO(), ctrl.Request{
		NamespacedName: types.NamespacedName{
			Namespace: workload.Namespace,
			Name:      workload.Name,
		},
	})
	require.Nil(t, err)
	require.False(t, res.Requeue)

	updatedWorkloadVersion := &klcv1beta1.KeptnWorkloadVersion{}
	err = r.Client.Get(context.TODO(),
		types.NamespacedName{
			Namespace: workloadVersion.Namespace,
			Name:      workloadVersion.Name,
		}, updatedWorkloadVersion)

	require.Nil(t, err)

	require.Equal(t, workload.Spec, updatedWorkloadVersion.Spec.KeptnWorkloadSpec)
}

func TestKeptnWorkload(t *testing.T) {
	workload := &klcv1beta1.KeptnWorkload{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "workload",
			Namespace: "namespace",
		},
		Spec: klcv1beta1.KeptnWorkloadSpec{
			Version: "version",
			AppName: "app",
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	}

	workloadVersion := generateWorkloadVersion("prev", map[string]string{}, workload)
	require.Equal(t, klcv1beta1.KeptnWorkloadVersion{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{},
			Name:        "workload-version",
			Namespace:   "namespace",
		},
		Spec: klcv1beta1.KeptnWorkloadVersionSpec{
			KeptnWorkloadSpec: klcv1beta1.KeptnWorkloadSpec{
				Version: "version",
				AppName: "app",
				Metadata: map[string]string{
					"foo": "bar",
				},
			},
			WorkloadName:    "workload",
			PreviousVersion: "prev",
		},
	}, workloadVersion)
}

func setupReconciler(objs ...client.Object) (*KeptnWorkloadReconciler, chan string) {
	// setup logger
	opts := zap.Options{
		Development: true,
	}
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	fakeClient := testcommon.NewTestClient(objs...)

	recorder := record.NewFakeRecorder(100)
	r := &KeptnWorkloadReconciler{
		Client:      fakeClient,
		Scheme:      scheme.Scheme,
		EventSender: eventsender.NewK8sSender(recorder),
		Log:         ctrl.Log.WithName("test-appController"),
	}
	return r, recorder.Events
}
