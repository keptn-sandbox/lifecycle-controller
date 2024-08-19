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

package options

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	optionsv1alpha1 "github.com/keptn/lifecycle-toolkit/lifecycle-operator/apis/options/v1alpha1"
	"github.com/keptn/lifecycle-toolkit/lifecycle-operator/controllers/common/config"
	"github.com/keptn/lifecycle-toolkit/lifecycle-operator/controllers/common/telemetry"
	controllererrors "github.com/keptn/lifecycle-toolkit/lifecycle-operator/controllers/errors"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KeptnConfigReconciler reconciles a KeptnConfig object
type KeptnConfigReconciler struct {
	client.Client
	Scheme          *runtime.Scheme
	Log             logr.Logger
	LastAppliedSpec *optionsv1alpha1.KeptnConfigSpec
	config          config.IConfig
	Namespace       string
}

func NewReconciler(client client.Client, scheme *runtime.Scheme, log logr.Logger, namespace string) *KeptnConfigReconciler {
	return &KeptnConfigReconciler{
		Client:    client,
		Scheme:    scheme,
		Log:       log,
		config:    config.Instance(),
		Namespace: namespace,
	}
}

// variables needed for Keptn Gateway
const gatewayPort = 8080

func getKeptnGatewayDeployment(namespace string) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Labels:    map[string]string{"app.kubernetes.io/name": "keptn-gateway"},
			Name:      "keptn-gateway",
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "keptn",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": "keptn"},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "keptn-gateway",
							Image: "keptn/keptn-gateway:latest",
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: gatewayPort,
								},
							},
						},
					},
				},
			},
		},
	}
}

func getKeptnGatewayService(namespace string) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels:    map[string]string{"app.kubernetes.io/name": "keptn-gateway"},
			Name:      "keptn-gateway-service",
			Namespace: namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{"app": "keptn"},
			Ports: []corev1.ServicePort{{
				Port:     gatewayPort,
				NodePort: 30091,
				Protocol: "TCP",
			}},
			Type: "NodePort",
		},
	}
}

// +kubebuilder:rbac:groups=options.keptn.sh,resources=keptnconfigs,verbs=get;list;watch
// +kubebuilder:rbac:groups=options.keptn.sh,resources=keptnconfigs/status,verbs=get
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=create;delete
// +kubebuilder:rbac:groups=apps,resources=deployments.apps,verbs=create;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=create;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *KeptnConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.Info("Searching for KeptnConfig")

	cfg := &optionsv1alpha1.KeptnConfig{}
	err := r.Get(ctx, req.NamespacedName, cfg)
	if errors.IsNotFound(err) {
		return reconcile.Result{}, nil
	}
	if err != nil {
		return reconcile.Result{}, fmt.Errorf(controllererrors.ErrCannotRetrieveConfigMsg, err)
	}

	if r.LastAppliedSpec == nil {
		r.Log.Info("initializing KeptnConfig since no config was there before")
		r.initConfig()
	}

	// reconcile config values
	r.config.SetCreationRequestTimeout(time.Duration(cfg.Spec.KeptnAppCreationRequestTimeoutSeconds) * time.Second)
	r.config.SetCloudEventsEndpoint(cfg.Spec.CloudEventsEndpoint)
	r.config.SetBlockDeployment(cfg.Spec.BlockDeployment)
	r.config.SetObservabilityTimeout(cfg.Spec.ObservabilityTimeout)
	r.config.SetKeptnGatewayEnabled(cfg.Spec.KeptnGatewayEnabled)

	result, err := r.reconcileOtelCollectorUrl(cfg)
	if err != nil {
		return result, err
	}

	result, err = r.reconcileKeptnGateway(ctx, cfg)
	if err != nil {
		return result, err
	}

	r.LastAppliedSpec = &cfg.Spec
	return ctrl.Result{}, nil
}

func (r *KeptnConfigReconciler) reconcileOtelCollectorUrl(config *optionsv1alpha1.KeptnConfig) (ctrl.Result, error) {
	r.Log.Info(fmt.Sprintf("reconciling Keptn Config: %s", config.Name))
	otelConfig := telemetry.GetOtelInstance()

	if err := otelConfig.InitOtelCollector(config.Spec.OTelCollectorUrl); err != nil {
		r.Log.Error(err, "unable to initialize OTel tracer options")
		return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, err
	}
	return ctrl.Result{}, nil
}

func (r *KeptnConfigReconciler) reconcileKeptnGateway(ctx context.Context, config *optionsv1alpha1.KeptnConfig) (ctrl.Result, error) {
	if config.Spec.KeptnGatewayEnabled {
		r.Log.Info("Creating Keptn-Gateway resources...")

		err := r.Client.Create(ctx, getKeptnGatewayDeployment(r.Namespace))
		if err != nil {
			r.Log.Error(err, "Unable to Deploy Keptn Gateway Deployment")
		}

		err = r.Client.Create(ctx, getKeptnGatewayService(r.Namespace))
		if err != nil {
			r.Log.Error(err, "Unable to Deploy Keptn Gateway Service")
		}

		r.Log.Info("Created Keptn-Gateway resources")
		return ctrl.Result{}, nil
	}

	r.Log.Info("Deleting Keptn-Gateway resources...")
	err := r.Client.Delete(ctx, getKeptnGatewayDeployment(r.Namespace))
	if err != nil {
		r.Log.Error(err, "Unable to Delete Keptn Gateway Deployment")
		if !errors.IsNotFound(err) {
			return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, err
		}
	}

	err = r.Client.Delete(ctx, getKeptnGatewayService(r.Namespace))
	if err != nil {
		r.Log.Error(err, "Unable to Delete Keptn Gateway Service")
		if !errors.IsNotFound(err) {
			return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, err
		}
	}
	r.Log.Info("Deleted Keptn-Gateway resources")

	return ctrl.Result{}, nil
}

func (r *KeptnConfigReconciler) initConfig() {
	r.LastAppliedSpec = &optionsv1alpha1.KeptnConfigSpec{}
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeptnConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&optionsv1alpha1.KeptnConfig{}).
		Complete(r)
}
