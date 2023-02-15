package providers

import (
	"context"
	"fmt"
	"github.com/keptn/lifecycle-toolkit/metrics-operator/apis/metrics/v1alpha2"
	"net/http"
	"strings"

	"github.com/go-logr/logr"
	"github.com/keptn/lifecycle-toolkit/metrics-operator/controllers/common/providers/dynatrace"
	"github.com/keptn/lifecycle-toolkit/metrics-operator/controllers/common/providers/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// KeptnSLIProvider is the interface that describes the operations that an SLI provider must implement
type KeptnSLIProvider interface {
	GetMetricValue(ctx context.Context, objective v1alpha2.KeptnMetric, provider v1alpha2.KeptnMetricProvider) (string, []byte, error)
}

// NewProvider is a factory method that chooses the right implementation of KeptnSLIProvider
func NewProvider(provider string, log logr.Logger, k8sClient client.Client) (KeptnSLIProvider, error) {
	switch strings.ToLower(provider) {
	case PrometheusProviderName:
		return &prometheus.KeptnPrometheusProvider{
			HttpClient: http.Client{},
			Log:        log,
		}, nil
	case DynatraceProviderName:
		return &dynatrace.KeptnDynatraceProvider{
			HttpClient: http.Client{},
			Log:        log,
			K8sClient:  k8sClient,
		}, nil
	case DynatraceDQLProviderName:
		return dynatrace.NewKeptnDynatraceDQLProvider(
			k8sClient,
			dynatrace.WithLogger(log),
		), nil
	default:
		return nil, fmt.Errorf("provider %s not supported", provider)
	}
}
