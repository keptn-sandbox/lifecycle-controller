# ⚠️This repository is still in experimental phase - we discourage any use in production. The code is provided with no expectation of support or maintenance.

# Keptn Lifecycle Controller

This repository hosts a prototype of the Keptn Lifecycle Controller.
The goal of this prototype is to introduce a more “cloud-native” approach for pre- and post-deployment, as well as the concept of application health checks.

The Keptn Lifecycle Controller is composed of the following components:

- Keptn Lifecycle Operator
- Keptn Scheduler

The Keptn Lifecycle Operator contains several controllers for Keptn CRDs and a Mutating Webhook.
The Keptn Scheduler ensures that Pods are started only after the pre-deployment checks have finished.

## Architecture

![](./assets/architecture.jpg)

A Kubernetes Manifest, which is annotated with Keptn specific annotations, gets applied to the Kubernetes Cluster.
Afterward, the Keptn Scheduler gets injected (via Mutating Webhook), and Kubernetes Events for Pre-Deployment are sent to the event stream.
The Event Controller watches for events and triggers a Kubernetes Job to fullfil the Pre-Deployment.
After the Pre-Deployment has finished, the Keptn Scheduler schedules the Pod to be deployed.
The Application and Service Controllers watchfor the workload resources to finish and then generate a Post-Deployment Event.
After the Post-Deployment checks, SLOs can be validated using an interface for retrieving SLI data from a provider, e.g, [Prometheus](https://prometheus.io/).
Finally, Keptn Lifecycle Controller exposes Metrics and Traces of the whole Deployment cycle with [OpenTelemetry](https://opentelemetry.io/).

## How to use

TBD

## How to install (development)

**Prerequisites:**

The lifecycle controllers includes a Mutating Webhook which requires TLS certificates to be mounted as a volume in its pod. The certificate creation
is handled automatically by [cert-manager](https://cert-manager.io). To install **cert-manager**, follow their [installation instructions](https://cert-manager.io/docs/installation/).

When cert-manager is installed, use the following commands to deploy the operator:

```bash
DOCKER_REGISTRY=<YOUR_DOCKER_REGISTRY>
DOCKER_TAG=<YOUR_DOCKER_TAG>

cd operator

make docker-build docker-push IMG=${DOCKER_REGISTRY}/${DOCKER_TAG}:latest
make deploy IMG=${DOCKER_REGISTRY}/${DOCKER_TAG}:latest
```


## License

Please find more information in the [LICENSE](LICENSE) file.