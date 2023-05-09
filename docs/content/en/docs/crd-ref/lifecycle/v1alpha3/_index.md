---
title: v1alpha3
description: Reference information for lifecycle.keptn.sh/v1alpha3
---
<!-- markdownlint-disable -->

## Packages
- [lifecycle.keptn.sh/v1alpha3](#lifecyclekeptnshv1alpha3)


## lifecycle.keptn.sh/v1alpha3

Package v1alpha3 contains API Schema definitions for the lifecycle v1alpha3 API group

### Resource Types
- [KeptnApp](#keptnapp)
- [KeptnAppCreationRequest](#keptnappcreationrequest)
- [KeptnAppCreationRequestList](#keptnappcreationrequestlist)
- [KeptnAppList](#keptnapplist)
- [KeptnAppVersion](#keptnappversion)
- [KeptnAppVersionList](#keptnappversionlist)
- [KeptnEvaluation](#keptnevaluation)
- [KeptnEvaluationDefinition](#keptnevaluationdefinition)
- [KeptnEvaluationDefinitionList](#keptnevaluationdefinitionlist)
- [KeptnEvaluationList](#keptnevaluationlist)
- [KeptnEvaluationProvider](#keptnevaluationprovider)
- [KeptnEvaluationProviderList](#keptnevaluationproviderlist)
- [KeptnTask](#keptntask)
- [KeptnTaskDefinition](#keptntaskdefinition)
- [KeptnTaskDefinitionList](#keptntaskdefinitionlist)
- [KeptnTaskList](#keptntasklist)
- [KeptnWorkload](#keptnworkload)
- [KeptnWorkloadInstance](#keptnworkloadinstance)
- [KeptnWorkloadInstanceList](#keptnworkloadinstancelist)
- [KeptnWorkloadList](#keptnworkloadlist)



#### ConfigMapReference





_Appears in:_
- [FunctionSpec](#functionspec)

| Field | Description |
| --- | --- |
| `name` _string_ | Name is the name of the referenced ConfigMap. |




#### EvaluationStatusItem





_Appears in:_
- [KeptnEvaluationStatus](#keptnevaluationstatus)

| Field | Description |
| --- | --- |
| `value` _string_ | Value represents the value of the KeptnMetric being evaluated. |
| `message` _string_ | Message contains additional information about the evaluation of an objective. This can include explanations about why an evaluation has failed (e.g. due to a missed objective), or if there was any error during the evaluation of the objective. |


#### FunctionReference





_Appears in:_
- [FunctionSpec](#functionspec)

| Field | Description |
| --- | --- |
| `name` _string_ | Name is the name of the referenced KeptnTaksDefinition. |


#### FunctionSpec





_Appears in:_
- [KeptnTaskDefinitionSpec](#keptntaskdefinitionspec)

| Field | Description |
| --- | --- |
| `functionRef` _[FunctionReference](#functionreference)_ | FunctionReference allows to reference another KeptnTaskDefinition which contains the source code of the function to be executes for KeptnTasks based on this KeptnTaskDefinition. This can be useful when you have multiple KeptnTaskDefinitions that should execute the same logic, but each with different parameters. |
| `inline` _[Inline](#inline)_ | Inline allows to specify the code that should be executed directly in the KeptnTaskDefinition, as a multi-line string. |
| `httpRef` _[HttpReference](#httpreference)_ | HttpReference allows to point to an HTTP URL containing the code of the function. |
| `configMapRef` _[ConfigMapReference](#configmapreference)_ | ConfigMapReference allows to reference a ConfigMap containing the code of the function. When referencing a ConfigMap, the code of the function must be available as a value of the 'code' key of the referenced ConfigMap. |
| `parameters` _[TaskParameters](#taskparameters)_ | Parameters contains parameters that will be passed to the job that executes the task. |
| `secureParameters` _[SecureParameters](#secureparameters)_ | SecureParameters contains secure parameters that will be passed to the job that executes the task. These will be stored and accessed as secrets in the cluster. |


#### FunctionStatus





_Appears in:_
- [KeptnTaskDefinitionStatus](#keptntaskdefinitionstatus)

| Field | Description |
| --- | --- |
| `configMap` _string_ | ConfigMap indicates the ConfigMap in which the function code is stored. |


#### HttpReference





_Appears in:_
- [FunctionSpec](#functionspec)

| Field | Description |
| --- | --- |
| `url` _string_ | Url is the URL containing the code of the function. |


#### Inline





_Appears in:_
- [FunctionSpec](#functionspec)

| Field | Description |
| --- | --- |
| `code` _string_ | Code contains the code of the function. |


#### ItemStatus





_Appears in:_
- [KeptnAppVersionStatus](#keptnappversionstatus)
- [KeptnWorkloadInstanceStatus](#keptnworkloadinstancestatus)

| Field | Description |
| --- | --- |
| `definitionName` _string_ | DefinitionName is the name of the EvaluationDefinition/TaskDefiniton |
| `name` _string_ | Name is the name of the Evaluation/Task |
| `startTime` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#time-v1-meta)_ | StartTime represents the time at which the Item (Evaluation/Task) started. |
| `endTime` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#time-v1-meta)_ | EndTime represents the time at which the Item (Evaluation/Task) started. |


#### KeptnApp



KeptnApp is the Schema for the keptnapps API

_Appears in:_
- [KeptnAppList](#keptnapplist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnApp`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnAppSpec](#keptnappspec)_ | Spec describes the desired state of the KeptnApp. |


#### KeptnAppCreationRequest



KeptnAppCreationRequest is the Schema for the keptnappcreationrequests API

_Appears in:_
- [KeptnAppCreationRequestList](#keptnappcreationrequestlist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnAppCreationRequest`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnAppCreationRequestSpec](#keptnappcreationrequestspec)_ | Spec describes the desired state of the KeptnAppCreationRequest. |


#### KeptnAppCreationRequestList



KeptnAppCreationRequestList contains a list of KeptnAppCreationRequest



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnAppCreationRequestList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnAppCreationRequest](#keptnappcreationrequest) array_ |  |


#### KeptnAppCreationRequestSpec



KeptnAppCreationRequestSpec defines the desired state of KeptnAppCreationRequest

_Appears in:_
- [KeptnAppCreationRequest](#keptnappcreationrequest)

| Field | Description |
| --- | --- |
| `appName` _string_ | AppName is the name of the KeptnApp the KeptnAppCreationRequest should create if no user-defined object with that name is found. |




#### KeptnAppList



KeptnAppList contains a list of KeptnApp



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnAppList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnApp](#keptnapp) array_ |  |


#### KeptnAppSpec



KeptnAppSpec defines the desired state of KeptnApp

_Appears in:_
- [KeptnApp](#keptnapp)
- [KeptnAppVersionSpec](#keptnappversionspec)

| Field | Description |
| --- | --- |
| `version` _string_ | Version defines the version of the application. For automatically created KeptnApps, the version is a function of all KeptnWorkloads that are part of the KeptnApp. |
| `revision` _integer_ | Revision can be modified to trigger another deployment of a KeptnApp of the same version. This can be used for restarting a KeptnApp which failed to deploy, e.g. due to a failed preDeploymentEvaluation/preDeploymentTask. |
| `workloads` _[KeptnWorkloadRef](#keptnworkloadref) array_ | Workloads is a list of all KeptnWorkloads that are part of the KeptnApp. |
| `preDeploymentTasks` _string array_ | PreDeploymentTasks is a list of all tasks to be performed during the pre-deployment phase of the KeptnApp. The items of this list refer to the names of KeptnTaskDefinitions located in the same namespace as the KeptnApp, or in the KLT namespace. |
| `postDeploymentTasks` _string array_ | PostDeploymentTasks is a list of all tasks to be performed during the post-deployment phase of the KeptnApp. The items of this list refer to the names of KeptnTaskDefinitions located in the same namespace as the KeptnApp, or in the KLT namespace. |
| `preDeploymentEvaluations` _string array_ | PreDeploymentEvaluations is a list of all evaluations to be performed during the pre-deployment phase of the KeptnApp. The items of this list refer to the names of KeptnEvaluationDefinitions located in the same namespace as the KeptnApp, or in the KLT namespace. |
| `postDeploymentEvaluations` _string array_ | PostDeploymentEvaluations is a list of all evaluations to be performed during the post-deployment phase of the KeptnApp. The items of this list refer to the names of KeptnEvaluationDefinitions located in the same namespace as the KeptnApp, or in the KLT namespace. |




#### KeptnAppVersion



KeptnAppVersion is the Schema for the keptnappversions API

_Appears in:_
- [KeptnAppVersionList](#keptnappversionlist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnAppVersion`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnAppVersionSpec](#keptnappversionspec)_ | Spec describes the desired state of the KeptnAppVersion. |


#### KeptnAppVersionList



KeptnAppVersionList contains a list of KeptnAppVersion



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnAppVersionList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnAppVersion](#keptnappversion) array_ |  |


#### KeptnAppVersionSpec



KeptnAppVersionSpec defines the desired state of KeptnAppVersion

_Appears in:_
- [KeptnAppVersion](#keptnappversion)

| Field | Description |
| --- | --- |
| `version` _string_ | Version defines the version of the application. For automatically created KeptnApps, the version is a function of all KeptnWorkloads that are part of the KeptnApp. |
| `revision` _integer_ | Revision can be modified to trigger another deployment of a KeptnApp of the same version. This can be used for restarting a KeptnApp which failed to deploy, e.g. due to a failed preDeploymentEvaluation/preDeploymentTask. |
| `workloads` _[KeptnWorkloadRef](#keptnworkloadref) array_ | Workloads is a list of all KeptnWorkloads that are part of the KeptnApp. |
| `preDeploymentTasks` _string array_ | PreDeploymentTasks is a list of all tasks to be performed during the pre-deployment phase of the KeptnApp. The items of this list refer to the names of KeptnTaskDefinitions located in the same namespace as the KeptnApp, or in the KLT namespace. |
| `postDeploymentTasks` _string array_ | PostDeploymentTasks is a list of all tasks to be performed during the post-deployment phase of the KeptnApp. The items of this list refer to the names of KeptnTaskDefinitions located in the same namespace as the KeptnApp, or in the KLT namespace. |
| `preDeploymentEvaluations` _string array_ | PreDeploymentEvaluations is a list of all evaluations to be performed during the pre-deployment phase of the KeptnApp. The items of this list refer to the names of KeptnEvaluationDefinitions located in the same namespace as the KeptnApp, or in the KLT namespace. |
| `postDeploymentEvaluations` _string array_ | PostDeploymentEvaluations is a list of all evaluations to be performed during the post-deployment phase of the KeptnApp. The items of this list refer to the names of KeptnEvaluationDefinitions located in the same namespace as the KeptnApp, or in the KLT namespace. |
| `appName` _string_ | AppName is the name of the KeptnApp. |
| `previousVersion` _string_ | PreviousVersion is the version of the KeptnApp that has been deployed prior to this version. |
| `traceId` _object (keys:string, values:string)_ | TraceId contains the OpenTelemetry trace ID. |




#### KeptnEvaluation



KeptnEvaluation is the Schema for the keptnevaluations API

_Appears in:_
- [KeptnEvaluationList](#keptnevaluationlist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnEvaluation`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnEvaluationSpec](#keptnevaluationspec)_ | Spec describes the desired state of the KeptnEvaluation. |


#### KeptnEvaluationDefinition



KeptnEvaluationDefinition is the Schema for the keptnevaluationdefinitions API

_Appears in:_
- [KeptnEvaluationDefinitionList](#keptnevaluationdefinitionlist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnEvaluationDefinition`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnEvaluationDefinitionSpec](#keptnevaluationdefinitionspec)_ | Spec describes the desired state of the KeptnEvaluationDefinition. |


#### KeptnEvaluationDefinitionList



KeptnEvaluationDefinitionList contains a list of KeptnEvaluationDefinition



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnEvaluationDefinitionList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnEvaluationDefinition](#keptnevaluationdefinition) array_ |  |


#### KeptnEvaluationDefinitionSpec



KeptnEvaluationDefinitionSpec defines the desired state of KeptnEvaluationDefinition

_Appears in:_
- [KeptnEvaluationDefinition](#keptnevaluationdefinition)

| Field | Description |
| --- | --- |
| `objectives` _[Objective](#objective) array_ | Objectives is a list of objectives that have to be met for a KeptnEvaluation referencing this KeptnEvaluationDefinition to be successful. |




#### KeptnEvaluationList



KeptnEvaluationList contains a list of KeptnEvaluation



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnEvaluationList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnEvaluation](#keptnevaluation) array_ |  |


#### KeptnEvaluationProvider



KeptnEvaluationProvider is the Schema for the keptnevaluationproviders API

_Appears in:_
- [KeptnEvaluationProviderList](#keptnevaluationproviderlist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnEvaluationProvider`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnEvaluationProviderSpec](#keptnevaluationproviderspec)_ |  |


#### KeptnEvaluationProviderList



KeptnEvaluationProviderList contains a list of KeptnEvaluationProvider



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnEvaluationProviderList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnEvaluationProvider](#keptnevaluationprovider) array_ |  |


#### KeptnEvaluationProviderSpec



KeptnEvaluationProviderSpec defines the desired state of KeptnEvaluationProvider

_Appears in:_
- [KeptnEvaluationProvider](#keptnevaluationprovider)

| Field | Description |
| --- | --- |
| `targetServer` _string_ |  |
| `secretKeyRef` _[SecretKeySelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#secretkeyselector-v1-core)_ |  |




#### KeptnEvaluationSpec



KeptnEvaluationSpec defines the desired state of KeptnEvaluation

_Appears in:_
- [KeptnEvaluation](#keptnevaluation)

| Field | Description |
| --- | --- |
| `workload` _string_ | Workload defines the KeptnWorkload for which the KeptnEvaluation is done. |
| `workloadVersion` _string_ | WorkloadVersion defines the version of the KeptnWorkload for which the KeptnEvaluation is done. |
| `appName` _string_ | AppName defines the KeptnApp for which the KeptnEvaluation is done. |
| `appVersion` _string_ | AppVersion defines the version of the KeptnApp for which the KeptnEvaluation is done. |
| `evaluationDefinition` _string_ | EvaluationDefinition refers to the name of the KeptnEvaluationDefinition which includes the objectives for the KeptnEvaluation. The KeptnEvaluationDefinition can be located in the same namespace as the KeptnEvaluation, or in the KLT namespace. |
| `retries` _integer_ | Retries indicates how many times the KeptnEvaluation can be attempted in the case of an error or missed evaluation objective, before considering the KeptnEvaluation to be failed. |
| `retryInterval` _[Duration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#duration-v1-meta)_ | RetryInterval specifies the interval at which the KeptnEvaluation is retried in the case of an error or a missed objective. |
| `failAction` _string_ |  |
| `checkType` _CheckType_ | Type indicates whether the KeptnEvaluation is part of the pre- or postDeployment phase. |




#### KeptnMetricReference





_Appears in:_
- [Objective](#objective)

| Field | Description |
| --- | --- |
| `name` _string_ | Name is the name of the referenced KeptnMetric. |
| `namespace` _string_ | Namespace is the namespace where the referenced KeptnMetric is located. |


#### KeptnTask



KeptnTask is the Schema for the keptntasks API

_Appears in:_
- [KeptnTaskList](#keptntasklist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnTask`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnTaskSpec](#keptntaskspec)_ | Spec describes the desired state of the KeptnTask. |


#### KeptnTaskDefinition



KeptnTaskDefinition is the Schema for the keptntaskdefinitions API

_Appears in:_
- [KeptnTaskDefinitionList](#keptntaskdefinitionlist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnTaskDefinition`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnTaskDefinitionSpec](#keptntaskdefinitionspec)_ | Spec describes the desired state of the KeptnTaskDefinition. |


#### KeptnTaskDefinitionList



KeptnTaskDefinitionList contains a list of KeptnTaskDefinition



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnTaskDefinitionList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnTaskDefinition](#keptntaskdefinition) array_ |  |


#### KeptnTaskDefinitionSpec



KeptnTaskDefinitionSpec defines the desired state of KeptnTaskDefinition

_Appears in:_
- [KeptnTaskDefinition](#keptntaskdefinition)

| Field | Description |
| --- | --- |
| `function` _[FunctionSpec](#functionspec)_ | Function contains the definition for the function that is to be executed in KeptnTasks based on the KeptnTaskDefinitions. |
| `retries` _integer_ | Retries specifies how many times a job executing the KeptnTaskDefinition should be restarted in the case of an unsuccessful attempt. |
| `timeout` _[Duration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#duration-v1-meta)_ | Timeout specifies the maximum time to wait for the task to be completed successfully. If the task does not complete successfully within this time frame, it will be considered to be failed. |




#### KeptnTaskList



KeptnTaskList contains a list of KeptnTask



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnTaskList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnTask](#keptntask) array_ |  |


#### KeptnTaskSpec



KeptnTaskSpec defines the desired state of KeptnTask

_Appears in:_
- [KeptnTask](#keptntask)

| Field | Description |
| --- | --- |
| `workload` _string_ | Workload defines the KeptnWorkload for which the KeptnTask is executed. |
| `workloadVersion` _string_ | WorkloadVersion defines the version of the KeptnWorkload for which the KeptnTask is executed. |
| `app` _string_ | AppName defines the KeptnApp for which the KeptnTask is executed. |
| `appVersion` _string_ | AppVersion defines the version of the KeptnApp for which the KeptnTask is executed. |
| `taskDefinition` _string_ | TaskDefinition refers to the name of the KeptnTaskDefinition which includes the specification for the task to be performed. The KeptnTaskDefinition can be located in the same namespace as the KeptnTask, or in the KLT namespace. |
| `context` _[TaskContext](#taskcontext)_ | Context contains contextual information about the task execution. |
| `parameters` _[TaskParameters](#taskparameters)_ | Parameters contains parameters that will be passed to the job that executes the task. |
| `secureParameters` _[SecureParameters](#secureparameters)_ | SecureParameters contains secure parameters that will be passed to the job that executes the task. These will be stored and accessed as secrets in the cluster. |
| `checkType` _CheckType_ | Type indicates whether the KeptnTask is part of the pre- or postDeployment phase. |
| `retries` _integer_ | Retries indicates how many times the KeptnTask can be attempted in the case of an error before considering the KeptnTask to be failed. |
| `timeout` _[Duration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#duration-v1-meta)_ | Timeout specifies the maximum time to wait for the task to be completed successfully. If the task does not complete successfully within this time frame, it will be considered to be failed. |




#### KeptnWorkload



KeptnWorkload is the Schema for the keptnworkloads API

_Appears in:_
- [KeptnWorkloadList](#keptnworkloadlist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnWorkload`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnWorkloadSpec](#keptnworkloadspec)_ | Spec describes the desired state of the KeptnWorkload. |


#### KeptnWorkloadInstance



KeptnWorkloadInstance is the Schema for the keptnworkloadinstances API

_Appears in:_
- [KeptnWorkloadInstanceList](#keptnworkloadinstancelist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnWorkloadInstance`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnWorkloadInstanceSpec](#keptnworkloadinstancespec)_ | Spec describes the desired state of the KeptnWorkloadInstance. |


#### KeptnWorkloadInstanceList



KeptnWorkloadInstanceList contains a list of KeptnWorkloadInstance



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnWorkloadInstanceList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnWorkloadInstance](#keptnworkloadinstance) array_ |  |


#### KeptnWorkloadInstanceSpec



KeptnWorkloadInstanceSpec defines the desired state of KeptnWorkloadInstance

_Appears in:_
- [KeptnWorkloadInstance](#keptnworkloadinstance)

| Field | Description |
| --- | --- |
| `app` _string_ | AppName is the name of the KeptnApp containing the KeptnWorkload. |
| `version` _string_ | Version defines the version of the KeptnWorkload. |
| `preDeploymentTasks` _string array_ | PreDeploymentTasks is a list of all tasks to be performed during the pre-deployment phase of the KeptnWorkload. The items of this list refer to the names of KeptnTaskDefinitions located in the same namespace as the KeptnApp, or in the KLT namespace. |
| `postDeploymentTasks` _string array_ | PostDeploymentTasks is a list of all tasks to be performed during the post-deployment phase of the KeptnWorkload. The items of this list refer to the names of KeptnTaskDefinitions located in the same namespace as the KeptnWorkload, or in the KLT namespace. |
| `preDeploymentEvaluations` _string array_ | PreDeploymentEvaluations is a list of all evaluations to be performed during the pre-deployment phase of the KeptnWorkload. The items of this list refer to the names of KeptnEvaluationDefinitions located in the same namespace as the KeptnWorkload, or in the KLT namespace. |
| `postDeploymentEvaluations` _string array_ | PostDeploymentEvaluations is a list of all evaluations to be performed during the post-deployment phase of the KeptnWorkload. The items of this list refer to the names of KeptnEvaluationDefinitions located in the same namespace as the KeptnWorkload, or in the KLT namespace. |
| `resourceReference` _[ResourceReference](#resourcereference)_ | ResourceReference is a reference to the Kubernetes resource (Deployment, DaemonSet, StatefulSet or ReplicaSet) the KeptnWorkload is representing. |
| `workloadName` _string_ | WorkloadName is the name of the KeptnWorkload. |
| `previousVersion` _string_ | PreviousVersion is the version of the KeptnWorkload that has been deployed prior to this version. |
| `traceId` _object (keys:string, values:string)_ | TraceId contains the OpenTelemetry trace ID. |




#### KeptnWorkloadList



KeptnWorkloadList contains a list of KeptnWorkload



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha3`
| `kind` _string_ | `KeptnWorkloadList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnWorkload](#keptnworkload) array_ |  |


#### KeptnWorkloadRef



KeptnWorkloadRef refers to a KeptnWorkload that is part of a KeptnApp

_Appears in:_
- [KeptnAppSpec](#keptnappspec)
- [KeptnAppVersionSpec](#keptnappversionspec)
- [WorkloadStatus](#workloadstatus)

| Field | Description |
| --- | --- |
| `name` _string_ | Name is the name of the KeptnWorkload. |
| `version` _string_ | Version is the version of the KeptnWorkload. |


#### KeptnWorkloadSpec



KeptnWorkloadSpec defines the desired state of KeptnWorkload

_Appears in:_
- [KeptnWorkload](#keptnworkload)
- [KeptnWorkloadInstanceSpec](#keptnworkloadinstancespec)

| Field | Description |
| --- | --- |
| `app` _string_ | AppName is the name of the KeptnApp containing the KeptnWorkload. |
| `version` _string_ | Version defines the version of the KeptnWorkload. |
| `preDeploymentTasks` _string array_ | PreDeploymentTasks is a list of all tasks to be performed during the pre-deployment phase of the KeptnWorkload. The items of this list refer to the names of KeptnTaskDefinitions located in the same namespace as the KeptnApp, or in the KLT namespace. |
| `postDeploymentTasks` _string array_ | PostDeploymentTasks is a list of all tasks to be performed during the post-deployment phase of the KeptnWorkload. The items of this list refer to the names of KeptnTaskDefinitions located in the same namespace as the KeptnWorkload, or in the KLT namespace. |
| `preDeploymentEvaluations` _string array_ | PreDeploymentEvaluations is a list of all evaluations to be performed during the pre-deployment phase of the KeptnWorkload. The items of this list refer to the names of KeptnEvaluationDefinitions located in the same namespace as the KeptnWorkload, or in the KLT namespace. |
| `postDeploymentEvaluations` _string array_ | PostDeploymentEvaluations is a list of all evaluations to be performed during the post-deployment phase of the KeptnWorkload. The items of this list refer to the names of KeptnEvaluationDefinitions located in the same namespace as the KeptnWorkload, or in the KLT namespace. |
| `resourceReference` _[ResourceReference](#resourcereference)_ | ResourceReference is a reference to the Kubernetes resource (Deployment, DaemonSet, StatefulSet or ReplicaSet) the KeptnWorkload is representing. |




#### Objective





_Appears in:_
- [KeptnEvaluationDefinitionSpec](#keptnevaluationdefinitionspec)

| Field | Description |
| --- | --- |
| `keptnMetricRef` _[KeptnMetricReference](#keptnmetricreference)_ | KeptnMetricRef references the KeptnMetric that should be evaluated. |
| `evaluationTarget` _string_ | EvaluationTarget specifies the target value for the references KeptnMetric. Needs to start with either '<' or '>', followed by the target value (e.g. '<10'). |


#### ResourceReference





_Appears in:_
- [KeptnWorkloadInstanceSpec](#keptnworkloadinstancespec)
- [KeptnWorkloadSpec](#keptnworkloadspec)

| Field | Description |
| --- | --- |
| `uid` _UID_ |  |
| `kind` _string_ |  |
| `name` _string_ |  |


#### SecureParameters





_Appears in:_
- [FunctionSpec](#functionspec)
- [KeptnTaskSpec](#keptntaskspec)

| Field | Description |
| --- | --- |
| `secret` _string_ | Secret contains the parameters that will be made available to the job executing the KeptnTask via the 'SECRET_DATA' environment variable. The 'SECRET_DATA'  environment variable's content will the same as value of the 'SECRET_DATA' key of the referenced secret. |


#### TaskContext





_Appears in:_
- [KeptnTaskSpec](#keptntaskspec)

| Field | Description |
| --- | --- |
| `workloadName` _string_ | WorkloadName the name of the KeptnWorkload the KeptnTask is being executed for. |
| `appName` _string_ | AppName the name of the KeptnApp the KeptnTask is being executed for. |
| `appVersion` _string_ | AppVersion the version of the KeptnApp the KeptnTask is being executed for. |
| `workloadVersion` _string_ | WorkloadVersion the version of the KeptnWorkload the KeptnTask is being executed for. |
| `taskType` _string_ | TaskType indicates whether the KeptnTask is part of the pre- or postDeployment phase. |
| `objectType` _string_ | ObjectType indicates whether the KeptnTask is being executed for a KeptnApp or KeptnWorkload. |


#### TaskParameters





_Appears in:_
- [FunctionSpec](#functionspec)
- [KeptnTaskSpec](#keptntaskspec)

| Field | Description |
| --- | --- |
| `map` _object (keys:string, values:string)_ | Inline contains the parameters that will be made available to the job executing the KeptnTask via the 'DATA' environment variable. The 'DATA'  environment variable's content will be a json encoded string containing all properties of the map provided. |


#### WorkloadStatus





_Appears in:_
- [KeptnAppVersionStatus](#keptnappversionstatus)

| Field | Description |
| --- | --- |
| `workload` _[KeptnWorkloadRef](#keptnworkloadref)_ | Workload refers to a KeptnWorkload that is part of the KeptnAppVersion. |


