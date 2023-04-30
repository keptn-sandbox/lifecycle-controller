---
title: Install KLT
description: Install the Keptn Lifecycle Toolkit
icon: concepts
layout: quickstart
weight: 35
hidechildren: false # this flag hides all sub-pages in the sidebar-multicard.html
---

Two methods are supported for installing the Keptn Lifecycle Toolkit:

* Releases v0.7.0 and later can be installed using
  the [Helm Chart](#use-helm-chart).
  This is the preferred strategy because it allows you to customize your cluster.

* All releases can be installed using
  the [manifests](#use-manifests).
  This is the less-preferred way because it does not support customization.

## Use Helm Chart

Version v0.7.0 and later of the Lifecycle Toolkit
should be installed using Helm Charts.
The command sequence to fetch and install the latest release is:

```shell
helm repo add klt https://charts.lifecycle.keptn.sh
helm repo update
helm upgrade --install keptn klt/klt \
   -n keptn-lifecycle-toolkit-system --create-namespace --wait
```

Note that the `helm repo update` command is used for fresh installs
as well as for upgrades.

Use the `--version <version>` flag on the
`helm upgrade --install` command line to specify a different KLT version.

Use the following command sequence to see a list of available versions:

```shell
helm repo update
helm search repo klt
```

To modify configuration options, download a copy of the
[helm/chart/values.yaml](https://github.com/keptn/lifecycle-toolkit/blob/main/helm/chart/values.yaml)
file, modify some values, and use the modified file to install KLT:

1. Download the `values.yaml` file:

   ```shell
   helm get values RELEASE_NAME [flags] > values.yaml
   ```

1. Edit your local copy to modify some values

1. Install KLT by adding the following string to your `helm upgrade` command line:

   ```shell
   --values=values.yaml
   ```

You can also use the `--set` flag
to specify a value change for the `helm upgrade --install` command.
Configuration options are specified using the format:

```shell
--set key1=value1,key2=value2,....
```

For more information,see

* The [Helm Get Values](https://helm.sh/docs/helm/helm_get_values/)) document

* The [helm-charts](https://github.com/keptn/lifecycle-toolkit/blob/main/helm/chart/README.md) page
  contains the full list of available values.

## Use manifests

All versions of the Lifecycle Toolkit can be installed using manifests,
although we recommend that you use Helm Charts
to install Version 0.7.0 and later
because the Helm Charts allow you to customize your configuration.

Versions 0.6.0 and earlier can only be installed using manifests.

> **Note** When installing Version 0.6.0,
you must first install the `cert-manager` with the following command sequence:

```shell
kubectl apply \
   -f https://github.com/cert-manager/cert-manager/releases/download/v1.11.0/cert-manager.yaml
kubectl wait \
   --for=condition=Available deployment/cert-manager-webhook -n cert-manager --timeout=60s

Use a command sequence like the following
to install the Lifecycle Toolkit from the manifest,
specifying the version you want to install.

```shell
kubectl apply \
   -f https://github.com/keptn/lifecycle-toolkit/releases/download/v0.6.0/manifest.yaml
kubectl wait --for=condition=Available deployment/lifecycle-operator \
   -n keptn-lifecycle-toolkit-system --timeout=120s
```

The Lifecycle Toolkit and its dependencies are now installed and ready to use.
