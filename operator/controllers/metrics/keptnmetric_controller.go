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

package metrics

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	klcv1alpha2 "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2"
	metricsv1alpha1 "github.com/keptn/lifecycle-toolkit/operator/apis/metrics/v1alpha1"
	"github.com/keptn/lifecycle-toolkit/operator/controllers/common/providers"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

const MB = 1 << (10 * 2)

// KeptnMetricReconciler reconciles a KeptnMetric object
type KeptnMetricReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

// clusterrole
// +kubebuilder:rbac:groups=metrics.keptn.sh,resources=providers,verbs=get;list;watch
// +kubebuilder:rbac:groups=metrics.keptn.sh,resources=keptnmetrics,verbs=get;list;watch;
// +kubebuilder:rbac:groups=metrics.keptn.sh,resources=keptnmetrics/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=metrics.keptn.sh,resources=keptnmetrics/finalizers,verbs=update

// role
// +kubebuilder:rbac:groups=core,namespace=keptn-lifecycle-toolkit-system,resources=secrets,verbs=get

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KeptnMetric object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *KeptnMetricReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.Info("Reconciling Metric")
	metric := &metricsv1alpha1.KeptnMetric{}

	if err := r.Client.Get(ctx, req.NamespacedName, metric); err != nil {
		if errors.IsNotFound(err) {
			// taking down all associated K8s resources is handled by K8s
			r.Log.Info("Metric resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		r.Log.Error(err, "Failed to get the Metric")
		return ctrl.Result{}, nil
	}

	fetchTime := metric.Status.LastUpdated.Add(time.Second * time.Duration(metric.Spec.FetchIntervalSeconds))
	if time.Now().Before(fetchTime) {
		diff := time.Until(fetchTime)
		r.Log.Info("Metric has not been updated for the configured interval. Skipping")
		return ctrl.Result{Requeue: true, RequeueAfter: diff}, nil
	}

	evaluationProvider, err := r.fetchProvider(ctx, types.NamespacedName{Name: metric.Spec.Provider.Name, Namespace: metric.Namespace})
	if err != nil {
		if errors.IsNotFound(err) {
			r.Log.Info(err.Error() + ", ignoring error since object must be deleted")
			return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, nil
		}
		r.Log.Error(err, "Failed to retrieve the provider")
		return ctrl.Result{RequeueAfter: 10 * time.Second}, nil
	}
	// load the provider
	provider, err2 := providers.NewProvider(metric.Spec.Provider.Name, r.Log, r.Client)
	if err2 != nil {
		r.Log.Error(err2, "Failed to get the correct Metric Provider")
		return ctrl.Result{Requeue: false}, err2
	}

	objective := klcv1alpha2.Objective{
		Name:  metric.Name,
		Query: metric.Spec.Query,
	}
	value, rawValue, err := provider.EvaluateQuery(ctx, objective, *evaluationProvider)
	if err != nil {
		r.Log.Error(err, "Failed to evaluate the query")
		return ctrl.Result{Requeue: false}, err
	}
	metric.Status.Value = value
	metric.Status.RawValue = cupSize(rawValue)
	metric.Status.LastUpdated = metav1.Time{Time: time.Now()}

	if err := r.Client.Status().Update(ctx, metric); err != nil {
		r.Log.Error(err, "Failed to update the Metric status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, nil
}

func cupSize(value []byte) []byte {
	if len(value) > MB {
		return value[:MB]
	}
	return value
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeptnMetricReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&metricsv1alpha1.KeptnMetric{}, builder.WithPredicates(predicate.GenerationChangedPredicate{})).
		Complete(r)
}

func (r *KeptnMetricReconciler) fetchProvider(ctx context.Context, namespacedMetric types.NamespacedName) (*klcv1alpha2.KeptnEvaluationProvider, error) {
	provider := &klcv1alpha2.KeptnEvaluationProvider{}
	if err := r.Client.Get(ctx, namespacedMetric, provider); err != nil {
		return nil, err
	}
	return provider, nil
}
