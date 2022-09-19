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

package controllers

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/google/uuid"
	"github.com/keptn-sandbox/lifecycle-controller/operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceReconciler reconciles a Service object
type ServiceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=services/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=services/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Service object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *ServiceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("Searching for service")

	service := &v1alpha1.Service{}
	err := r.Get(ctx, req.NamespacedName, service)
	if errors.IsNotFound(err) {
		return reconcile.Result{}, nil
	}

	if err != nil {
		return reconcile.Result{}, fmt.Errorf("could not fetch Service: %+v", err)
	}

	if service.IsCompleted() {
		return reconcile.Result{}, nil
	}

	logger.Info("Reconciling Service", "service", service.Name)

	if service.IsServiceRunNotCreated() {
		logger.Info("Service Run does not exist, creating")

		serviceRunName, err := r.createServiceRun(ctx, service)
		if err != nil {
			logger.Error(err, "Could not create ServiceRun")
			return reconcile.Result{}, err
		}

		service.Status.Phase = v1alpha1.ServiceRunRunning
		service.Status.ServiceRunName = serviceRunName

		if err := r.Status().Update(ctx, service); err != nil {
			logger.Error(err, "Could not update Service")
			return reconcile.Result{}, err
		}
		return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Second}, nil
	}

	logger.Info("Checking status")

	serviceRun := &v1alpha1.ServiceRun{}
	err = r.Get(ctx, types.NamespacedName{Name: service.Status.ServiceRunName, Namespace: service.Namespace}, serviceRun)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("could not fetch ServiceRun: %+v", err)
	}

	if serviceRun.IsCompleted() {
		service.Status.Phase = serviceRun.Status.Phase

		if err := r.Status().Update(ctx, service); err != nil {
			logger.Error(err, "Could not update Service")
			return reconcile.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Second}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ServiceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Service{}).
		Complete(r)
}

func (r *ServiceReconciler) createServiceRun(ctx context.Context, service *v1alpha1.Service) (string, error) {
	serviceRun := &v1alpha1.ServiceRun{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"keptn.sh/application": service.Spec.ApplicationName,
				"keptn.sh/service":     service.Name,
			},
			Name:      service.Name + "-" + r.generateSuffix(),
			Namespace: service.Namespace,
		},
		Spec: v1alpha1.ServiceRunSpec{
			ServiceName: service.Name,
		},
	}
	for i := 0; i < 5; i++ {
		if err := r.Create(ctx, serviceRun); err != nil {
			if errors.IsAlreadyExists(err) {
				serviceRun.Name = service.Name + "-" + r.generateSuffix()
				continue
			}
			return "", err
		}
		break
	}
	return serviceRun.Name, nil
}

func (r *ServiceReconciler) generateSuffix() string {
	uid := uuid.New().String()
	return uid[:10]
}
