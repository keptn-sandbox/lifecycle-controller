package common

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	lifecyclev1alpha2 "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2"
	"github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2/common"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/instrument/asyncfloat64"
	"go.opentelemetry.io/otel/metric/instrument/asyncint64"
	"go.opentelemetry.io/otel/metric/unit"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	logger       = ctrl.Log.WithName("otel-utils")
	gitCommit    string
	buildTime    string
	buildVersion string
	otelInitOnce sync.Once
)

type otelConfig struct {
	TracerProvider *trace.TracerProvider
	OtelExporter   *trace.SpanExporter
}

// do not export this type to make it accessible only via the GetInstance method (i.e Singleton)
var otelInstance *otelConfig

func GetOtelInstance() *otelConfig {
	// initialize once
	otelInitOnce.Do(func() {
		otelInstance = &otelConfig{}
	})

	return otelInstance
}

func (o *otelConfig) InitOtelCollector(otelCollectorUrl string) error {
	tpOptions, otelExporter, err := GetOTelTracerProviderOptions(otelCollectorUrl)
	if err != nil {
		return err
	}

	o.TracerProvider = trace.NewTracerProvider(tpOptions...)
	otel.SetTracerProvider(o.TracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	o.OtelExporter = &otelExporter
	logger.Info("Successfully initialized OTel collector")
	return nil
}

func (o *otelConfig) ShutDown() {
	if err := o.TracerProvider.Shutdown(context.Background()); err != nil {
		os.Exit(1)
	}
}

func GetOTelTracerProviderOptions(oTelCollectorUrl string) ([]trace.TracerProviderOption, trace.SpanExporter, error) {
	var tracerProviderOptions []trace.TracerProviderOption
	var otelExporter trace.SpanExporter

	stdOutExp, err := newStdOutExporter()
	if err != nil {
		return nil, nil, fmt.Errorf("could not create stdout OTel exporter: %w", err)
	}
	tracerProviderOptions = append(tracerProviderOptions, trace.WithBatcher(stdOutExp))

	if oTelCollectorUrl != "" {
		// try to set OTel exporter for Jaeger
		otelExporter, err := newOTelExporter(oTelCollectorUrl)
		if err != nil {
			// log the error, but do not break if Jaeger exporter cannot be created
			logger.Error(err, "Could not set up OTel exporter")
			return nil, nil, err
		} else if otelExporter != nil {
			tracerProviderOptions = append(tracerProviderOptions, trace.WithBatcher(otelExporter))
		}
	}
	tracerProviderOptions = append(tracerProviderOptions, trace.WithResource(newResource()))

	return tracerProviderOptions, otelExporter, nil
}

func newStdOutExporter() (trace.SpanExporter, error) {
	return stdouttrace.New(
		// Use human readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

func newOTelExporter(oTelCollectorUrl string) (trace.SpanExporter, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, oTelCollectorUrl, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector at %s: %w", oTelCollectorUrl, err)
	}
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}
	return traceExporter, nil
}

func newResource() *resource.Resource {
	r := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.TelemetrySDKLanguageGo,
		semconv.ServiceNameKey.String("keptn-lifecycle-operator"),
		semconv.ServiceVersionKey.String(buildVersion+"-"+gitCommit+"-"+buildTime),
	)
	return r
}

func SetUpKeptnMeters(meter metric.Meter, mgr client.Client) {
	deploymentActiveGauge, err := meter.AsyncInt64().Gauge("keptn.deployment.active", instrument.WithDescription("a gauge keeping track of the currently active Keptn Deployments"))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}
	taskActiveGauge, err := meter.AsyncInt64().Gauge("keptn.task.active", instrument.WithDescription("a simple counter of active Keptn Tasks"))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}
	appActiveGauge, err := meter.AsyncInt64().Gauge("keptn.app.active", instrument.WithDescription("a simple counter of active Keptn Apps"))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}
	evaluationActiveGauge, err := meter.AsyncInt64().Gauge("keptn.evaluation.active", instrument.WithDescription("a simple counter of active Keptn Evaluations"))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}
	appDeploymentIntervalGauge, err := meter.AsyncFloat64().Gauge("keptn.app.deploymentinterval", instrument.WithDescription("a gauge of the interval between deployments"))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}

	appDeploymentDurationGauge, err := meter.AsyncFloat64().Gauge("keptn.app.deploymentduration", instrument.WithDescription("a gauge of the duration of deployments"))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}

	workloadDeploymentIntervalGauge, err := meter.AsyncFloat64().Gauge("keptn.deployment.deploymentinterval", instrument.WithDescription("a gauge of the interval between deployments"))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}

	workloadDeploymentDurationGauge, err := meter.AsyncFloat64().Gauge("keptn.deployment.deploymentduration", instrument.WithDescription("a gauge of the duration of deployments"))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}

	err = meter.RegisterCallback(
		[]instrument.Asynchronous{
			deploymentActiveGauge,
			taskActiveGauge,
			appActiveGauge,
			evaluationActiveGauge,
			appDeploymentIntervalGauge,
			appDeploymentDurationGauge,
			workloadDeploymentIntervalGauge,
			workloadDeploymentDurationGauge,
		},
		func(ctx context.Context) {
			observeActiveInstances(ctx, mgr, deploymentActiveGauge, appActiveGauge, taskActiveGauge, evaluationActiveGauge)
			observeDeploymentInterval(ctx, mgr, appDeploymentIntervalGauge, workloadDeploymentIntervalGauge)
			observeDuration(ctx, mgr, appDeploymentDurationGauge, workloadDeploymentDurationGauge)
		})
	if err != nil {
		fmt.Println("Failed to register callback")
		panic(err)
	}
}

func observeDuration(ctx context.Context, mgr client.Client, appDeploymentDurationGauge asyncfloat64.Gauge, workloadDeploymentDurationGauge asyncfloat64.Gauge) {

	err := ObserveDeploymentDuration(ctx, mgr, &lifecyclev1alpha2.KeptnAppVersionList{}, appDeploymentDurationGauge)
	if err != nil {
		logger.Error(err, "unable to gather app deployment durations")
	}

	err = ObserveDeploymentDuration(ctx, mgr, &lifecyclev1alpha2.KeptnWorkloadInstanceList{}, workloadDeploymentDurationGauge)
	if err != nil {
		logger.Error(err, "unable to gather workload deployment durations")
	}

}

func observeDeploymentInterval(ctx context.Context, mgr client.Client, appDeploymentIntervalGauge asyncfloat64.Gauge, workloadDeploymentIntervalGauge asyncfloat64.Gauge) error {
	err := ObserveDeploymentInterval(ctx, mgr, &lifecyclev1alpha2.KeptnAppVersionList{}, &lifecyclev1alpha2.KeptnWorkloadInstance{}, appDeploymentIntervalGauge)
	if err != nil {
		logger.Error(err, "unable to gather app deployment intervals")
	}

	err = ObserveDeploymentInterval(ctx, mgr, &lifecyclev1alpha2.KeptnWorkloadInstanceList{}, &lifecyclev1alpha2.KeptnWorkloadInstance{}, workloadDeploymentIntervalGauge)
	if err != nil {
		logger.Error(err, "unable to gather workload deployment intervals")
	}
	return err
}

func observeActiveInstances(ctx context.Context, mgr client.Client, deploymentActiveGauge asyncint64.Gauge, appActiveGauge asyncint64.Gauge, taskActiveGauge asyncint64.Gauge, evaluationActiveGauge asyncint64.Gauge) {

	err := ObserveActiveInstances(ctx, mgr, &lifecyclev1alpha2.KeptnWorkloadInstanceList{}, deploymentActiveGauge)
	if err != nil {
		logger.Error(err, "unable to gather active deployments")
	}
	err = ObserveActiveInstances(ctx, mgr, &lifecyclev1alpha2.KeptnAppVersionList{}, appActiveGauge)
	if err != nil {
		logger.Error(err, "unable to gather active apps")
	}
	err = ObserveActiveInstances(ctx, mgr, &lifecyclev1alpha2.KeptnTaskList{}, taskActiveGauge)
	if err != nil {
		logger.Error(err, "unable to gather active tasks")
	}
	err = ObserveActiveInstances(ctx, mgr, &lifecyclev1alpha2.KeptnEvaluationList{}, evaluationActiveGauge)
	if err != nil {
		logger.Error(err, "unable to gather active evaluations")
	}
}

func SetUpKeptnTaskMeters(meter metric.Meter) common.KeptnMeters {
	deploymentCount, err := meter.SyncInt64().Counter("keptn.deployment.count", instrument.WithDescription("a simple counter for Keptn Deployments"))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}
	deploymentDuration, err := meter.SyncFloat64().Histogram("keptn.deployment.duration", instrument.WithDescription("a histogram of duration for Keptn Deployments"), instrument.WithUnit(unit.Unit("s")))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}
	taskCount, err := meter.SyncInt64().Counter("keptn.task.count", instrument.WithDescription("a simple counter for Keptn Tasks"))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}
	taskDuration, err := meter.SyncFloat64().Histogram("keptn.task.duration", instrument.WithDescription("a histogram of duration for Keptn Tasks"), instrument.WithUnit(unit.Unit("s")))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}
	appCount, err := meter.SyncInt64().Counter("keptn.app.count", instrument.WithDescription("a simple counter for Keptn Apps"))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}
	appDuration, err := meter.SyncFloat64().Histogram("keptn.app.duration", instrument.WithDescription("a histogram of duration for Keptn Apps"), instrument.WithUnit(unit.Unit("s")))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}
	evaluationCount, err := meter.SyncInt64().Counter("keptn.evaluation.count", instrument.WithDescription("a simple counter for Keptn Evaluations"))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}
	evaluationDuration, err := meter.SyncFloat64().Histogram("keptn.evaluation.duration", instrument.WithDescription("a histogram of duration for Keptn Evaluations"), instrument.WithUnit(unit.Unit("s")))
	if err != nil {
		logger.Error(err, "unable to start OTel")
	}

	meters := common.KeptnMeters{
		TaskCount:          taskCount,
		TaskDuration:       taskDuration,
		DeploymentCount:    deploymentCount,
		DeploymentDuration: deploymentDuration,
		AppCount:           appCount,
		AppDuration:        appDuration,
		EvaluationCount:    evaluationCount,
		EvaluationDuration: evaluationDuration,
	}
	return meters
}
