# kyverno

Kyverno is a policy engine designed for Kubernetes. It can validate, mutate,
and generate configurations using admission controls and background scans.
Kyverno policies are Kubernetes resources and do not require learning a new
language.

**Homepage:** <https://github.com/giantswarm/kyverno-app>

## Source Code

* <https://github.com/kyverno/kyverno>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
|  | kyverno | 3.8.2 |
|  | policy-reporter | 3.10.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| global.podSecurityStandards.enforced | bool | `true` |  |
| global.image.registry | string | `"gsoci.azurecr.io"` |  |
| ciliumNetworkPolicy.enabled | bool | `true` |  |
| ciliumNetworkPolicy.admissionControllerExtraEgress.enabled | bool | `false` |  |
| ciliumNetworkPolicy.admissionControllerExtraEgress.dnsSelector.rules | list | `[]` |  |
| ciliumNetworkPolicy.admissionControllerExtraEgress.fqdnsConnection.port | string | `"443"` |  |
| ciliumNetworkPolicy.admissionControllerExtraEgress.fqdnsConnection.protocol | string | `"TCP"` |  |
| ciliumNetworkPolicy.admissionControllerExtraEgress.fqdnsConnection.rules | list | `[]` |  |
| crds.install | bool | `false` |  |
| crds.resources.requests.cpu | string | `"100m"` |  |
| crds.resources.requests.memory | string | `"256Mi"` |  |
| crds.resources.limits.cpu | string | `"750m"` |  |
| crds.resources.limits.memory | string | `"1024Mi"` |  |
| image.registry | string | `"gsoci.azurecr.io"` |  |
| image.repository | string | `"giantswarm/kubectl"` |  |
| image.tag | string | `"v1.34.3"` | Image tag Defaults to `latest` if omitted |
| verticalPodAutoscaler.admissionController.enabled | bool | `true` |  |
| verticalPodAutoscaler.admissionController.containerPolicies | object | `{}` |  |
| verticalPodAutoscaler.backgroundController.enabled | bool | `true` |  |
| verticalPodAutoscaler.backgroundController.containerPolicies.minAllowed.cpu | string | `"100m"` |  |
| verticalPodAutoscaler.backgroundController.containerPolicies.minAllowed.memory | string | `"100Mi"` |  |
| verticalPodAutoscaler.backgroundController.containerPolicies.maxAllowed.cpu | int | `4` |  |
| verticalPodAutoscaler.backgroundController.containerPolicies.maxAllowed.memory | string | `"8Gi"` |  |
| verticalPodAutoscaler.cleanupController.enabled | bool | `true` |  |
| verticalPodAutoscaler.cleanupController.containerPolicies | object | `{}` |  |
| verticalPodAutoscaler.reportsController.enabled | bool | `true` |  |
| verticalPodAutoscaler.reportsController.containerPolicies.minAllowed.cpu | string | `"100m"` |  |
| verticalPodAutoscaler.reportsController.containerPolicies.minAllowed.memory | string | `"100Mi"` |  |
| verticalPodAutoscaler.reportsController.containerPolicies.maxAllowed.cpu | int | `6` |  |
| verticalPodAutoscaler.reportsController.containerPolicies.maxAllowed.memory | string | `"24Gi"` |  |
| verticalPodAutoscaler.kyvernoPlugin.enabled | bool | `true` |  |
| verticalPodAutoscaler.kyvernoPlugin.containerPolicies.minAllowed.cpu | string | `"50m"` |  |
| verticalPodAutoscaler.kyvernoPlugin.containerPolicies.minAllowed.memory | string | `"50Mi"` |  |
| verticalPodAutoscaler.policyReporter.enabled | bool | `true` |  |
| verticalPodAutoscaler.policyReporter.containerPolicies.minAllowed.cpu | string | `"50m"` |  |
| verticalPodAutoscaler.policyReporter.containerPolicies.minAllowed.memory | string | `"50Mi"` |  |
| verticalPodAutoscaler.ui.enabled | bool | `true` |  |
| verticalPodAutoscaler.ui.containerPolicies.minAllowed.cpu | string | `"50m"` |  |
| verticalPodAutoscaler.ui.containerPolicies.minAllowed.memory | string | `"50Mi"` |  |
| cleanupPolicies.trivyOperator.enabled | bool | `false` | Enable Trivy cleanup policy to delete old Trivy Operator reports |
| cleanupPolicies.trivyOperator.schedule | string | `"37 * * * *"` | Cleanup schedule |
| cleanupPolicies.trivyOperator.olderThan | string | `"336h"` | Reports older than this duration will be deleted. Needs to be specified in hours. Defaults to 2 weeks (336 hours). |
| policyExceptions.enablePolexPolicy | bool | `true` |  |
| policyExceptions.allowedPolexNamespaces[0] | string | `"policy-exceptions"` |  |
| policyExceptions.polexPolicyMessage | string | `"PolicyExceptions are not allowed to be created in the {{ request.namespace }} namespace. Please contact a cluster administrator for assistance."` |  |
| policyExceptions.enableChartOperatorPolex | bool | `true` |  |
| policyExceptions.enableCiliumPolex | bool | `true` |  |
| policyExceptions.enableWildcardMatchPolicy | bool | `true` |  |
| policyExceptions.enableNoisyContextsPolicy | bool | `true` |  |
| policyExceptions.enableAwsCloudControllerManagerPolex | bool | `true` |  |
| policyExceptions.enableAwsEbsCsiDriverPolex | bool | `true` |  |
| policyExceptions.enableAzureCloudControllerManagerPolex | bool | `true` |  |
| policyExceptions.enableAzureCloudNodeManagerPolex | bool | `true` |  |
| upgradeJob.enabled | bool | `false` |  |
| monitoring.podLogs.enabled | bool | `true` |  |
| monitoring.podLogs.tenant | string | `"giantswarm"` |  |
| monitoring.podLogs.annotations | object | `{}` |  |
| monitoring.podLogs.labels | object | `{}` |  |
| monitoring.podLogs.relabelings | list | `[]` |  |
| kyverno.openreports.enabled | bool | `false` | Enable OpenReports feature in controllers |
| kyverno.openreports.installCrds | bool | `false` | Whether to install CRDs from the upstream OpenReports chart. Setting this to true requires enabled to also be true. |
| kyverno.reportsServer.enabled | bool | `false` | Enable reports-server deployment alongside Kyverno |
| kyverno.reportsServer.waitForReady | bool | `true` | Wait for reports-server to be ready before starting Kyverno components |
| kyverno.reportsServer.readinessTimeout | int | `300` | Timeout for waiting for reports-server readiness (in seconds) |
| kyverno.test.image.registry | string | `"gsoci.azurecr.io"` |  |
| kyverno.test.image.repository | string | `"giantswarm/kyverno-readiness-checker"` |  |
| kyverno.crds.install | bool | `false` | We manage CRDs with an outside job so this needs to be disabled |
| kyverno.crds.migration.enabled | bool | `true` |  |
| kyverno.crds.migration.image.registry | string | `"gsoci.azurecr.io"` |  |
| kyverno.crds.migration.image.repository | string | `"giantswarm/kyverno-cli"` |  |
| kyverno.grafana.enabled | bool | `false` | Enable grafana dashboard creation. |
| kyverno.grafana.grafanaDashboard | object | `{"allowCrossNamespaceImport":true,"create":false,"folder":"kyverno","matchLabels":{"dashboards":"grafana"}}` | create GrafanaDashboard custom resource referencing to the configMap. |
| kyverno.config.excludeKyvernoNamespace | bool | `true` | Exclude Kyverno namespace Determines if default Kyverno namespace exclusion is enabled for webhooks and resourceFilters |
| kyverno.config.enableDefaultRegistryMutation | bool | `true` | Enable registry mutation for container images. Enabled by default. |
| kyverno.config.defaultRegistry | string | `"docker.io"` | The registry hostname used for the image mutation. |
| kyverno.config.maxContextSize | string | 2Mi | Maximum cumulative size of context data during policy evaluation. Supports Kubernetes quantity format (e.g., 100Mi, 2Gi) or plain bytes (e.g., 2097152). Limits memory used by context variables to prevent unbounded growth. Increase if policies legitimately need large context data (e.g., processing large ConfigMaps). Set to 0 to disable the limit (not recommended for production). |
| kyverno.config.resourceFiltersExcludeNamespaces | list | `[]` | resourceFilter namespace exclude Namespaces to exclude from the default resourceFilters |
| kyverno.config.resourceFiltersExclude | list | `[]` | resourceFilters exclude list Items to exclude from config.resourceFilters |
| kyverno.config.resourceFiltersIncludeNamespaces | list | `[]` | resourceFilter namespace include Namespaces to include to the default resourceFilters |
| kyverno.config.resourceFiltersInclude | list | `[]` | resourceFilters include list Items to include to config.resourceFilters |
| kyverno.config.webhooks | object | `{"namespaceSelector":{"matchExpressions":[{"key":"kubernetes.io/metadata.name","operator":"NotIn","values":["kube-system"]}]}}` | Defines the `namespaceSelector`/`objectSelector` in the webhook configurations. The Kyverno namespace is excluded if `excludeKyvernoNamespace` is `true` (default) |
| kyverno.metricsConfig.create | bool | `true` | Create the configmap. |
| kyverno.metricsConfig.annotations | object | `{}` | Additional annotations to add to the configmap. |
| kyverno.metricsConfig.namespaces.include | list | `[]` | List of namespaces to capture metrics for. |
| kyverno.metricsConfig.namespaces.exclude | list | `[]` | list of namespaces to NOT capture metrics for. |
| kyverno.metricsConfig.bucketBoundaries | list | `[0.005,0.01,0.025,0.05,0.1,0.25,0.5,1,2.5,5,10,15,20,25,30]` | Configures the bucket boundaries for all Histogram metrics |
| kyverno.features.admissionReports.enabled | bool | `true` | Enables the feature |
| kyverno.features.aggregateReports.enabled | bool | `true` | Enables the feature |
| kyverno.features.policyReports.enabled | bool | `true` | Enables the feature |
| kyverno.features.validatingAdmissionPolicyReports.enabled | bool | `true` | Enables the feature |
| kyverno.features.mutatingAdmissionPolicyReports.enabled | bool | `false` | Enables the feature |
| kyverno.features.reporting.validate | bool | `true` | Enables the feature |
| kyverno.features.reporting.mutate | bool | `true` | Enables the feature |
| kyverno.features.reporting.mutateExisting | bool | `true` | Enables the feature |
| kyverno.features.reporting.imageVerify | bool | `true` | Enables the feature |
| kyverno.features.reporting.generate | bool | `true` | Enables the feature |
| kyverno.features.autoUpdateWebhooks.enabled | bool | `true` | Enables the feature |
| kyverno.features.backgroundScan.enabled | bool | `true` | Enables the feature |
| kyverno.features.backgroundScan.backgroundScanWorkers | int | `2` | Number of background scan workers |
| kyverno.features.backgroundScan.backgroundScanInterval | string | `"1h"` | Background scan interval |
| kyverno.features.backgroundScan.skipResourceFilters | bool | `true` | Skips resource filters in background scan |
| kyverno.features.configMapCaching.enabled | bool | `true` | Enables the feature |
| kyverno.features.controllerRuntimeMetrics.bindAddress | string | `":8080"` | Bind address for controller-runtime metrics (use "0" to disable it) |
| kyverno.features.deferredLoading.enabled | bool | `true` | Enables the feature |
| kyverno.features.dumpPayload.enabled | bool | `false` | Enables the feature |
| kyverno.features.forceFailurePolicyIgnore.enabled | bool | `false` | Enables the feature |
| kyverno.features.generateValidatingAdmissionPolicy.enabled | bool | `true` | Enables the feature |
| kyverno.features.generateMutatingAdmissionPolicy.enabled | bool | `false` | Enables the feature |
| kyverno.features.dumpPatches.enabled | bool | `false` | Enables the feature |
| kyverno.features.globalContext.maxApiCallResponseLength | int | `2000000` | Maximum allowed response size from API Calls. A value of 0 bypasses checks (not recommended) |
| kyverno.features.logging.format | string | `"text"` | Logging format |
| kyverno.features.logging.verbosity | int | `2` | Logging verbosity |
| kyverno.features.omitEvents.eventTypes | list | `["PolicyApplied","PolicyError","PolicySkipped"]` | Events which should not be emitted (possible values `PolicyViolation`, `PolicyApplied`, `PolicyError`, and `PolicySkipped`) |
| kyverno.features.policyExceptions.enabled | bool | `true` | Enables the feature |
| kyverno.features.policyExceptions.namespace | string | `"*"` | Restrict policy exceptions to a single namespace |
| kyverno.features.protectManagedResources.enabled | bool | `false` | Enables the feature |
| kyverno.features.registryClient.allowInsecure | bool | `false` | Allow insecure registry |
| kyverno.features.registryClient.credentialHelpers | list | `["default","google","amazon","azure","github"]` | Enable registry client helpers |
| kyverno.features.ttlController.reconciliationInterval | string | `"1m"` | Reconciliation interval for the label based cleanup manager |
| kyverno.features.tuf.enabled | bool | `false` | Enables the feature |
| kyverno.features.tuf.root | string | `nil` | Path to Tuf root |
| kyverno.features.tuf.rootRaw | string | `nil` | Raw Tuf root |
| kyverno.features.tuf.mirror | string | `nil` | Tuf mirror |
| kyverno.customLabels."giantswarm.io/service-type" | string | `"managed"` |  |
| kyverno.customLabels."application.giantswarm.io/team" | string | `"shield"` |  |
| kyverno.admissionController.autoscaling.enabled | bool | `true` | Enable horizontal pod autoscaling |
| kyverno.admissionController.autoscaling.minReplicas | int | `3` | Minimum number of pods |
| kyverno.admissionController.autoscaling.maxReplicas | int | `10` | Maximum number of pods |
| kyverno.admissionController.autoscaling.targetCPUUtilizationPercentage | int | `80` | Target CPU utilization percentage |
| kyverno.admissionController.autoscaling.behavior | object | `{}` | Configurable scaling behavior |
| kyverno.admissionController.featuresOverride | object | `{"admissionReports":{"backPressureThreshold":1000}}` | Overrides features defined at the root level |
| kyverno.admissionController.featuresOverride.admissionReports.backPressureThreshold | int | `1000` | Max number of admission reports allowed in flight until the admission controller stops creating new ones |
| kyverno.admissionController.rbac.create | bool | `true` | Create RBAC resources |
| kyverno.admissionController.rbac.serviceAccount.name | string | `"kyverno-admission-controller"` | The ServiceAccount name |
| kyverno.admissionController.rbac.serviceAccount.annotations | object | `{}` | Annotations for the ServiceAccount |
| kyverno.admissionController.rbac.clusterRole.extraResources | list | `[]` | Extra resource permissions to add in the cluster role |
| kyverno.admissionController.createSelfSignedCert | bool | `false` | Create self-signed certificates at deployment time. The certificates won't be automatically renewed if this is set to `true`. |
| kyverno.admissionController.replicas | int | `3` | Desired number of pods |
| kyverno.admissionController.podLabels | object | `{}` | Additional labels to add to each pod |
| kyverno.admissionController.podAnnotations | object | `{}` | Additional annotations to add to each pod |
| kyverno.admissionController.updateStrategy | object | See [values.yaml](values.yaml) | Deployment update strategy. Ref: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#strategy |
| kyverno.admissionController.priorityClassName | string | `"giantswarm-critical"` | Optional priority class |
| kyverno.admissionController.apiPriorityAndFairness | bool | `true` | Change `apiPriorityAndFairness` to `true` if you want to insulate the API calls made by Kyverno admission controller activities. This will help ensure Kyverno stability in busy clusters. Ref: https://kubernetes.io/docs/concepts/cluster-administration/flow-control/ |
| kyverno.admissionController.nodeSelector | object | `{}` | Node labels for pod assignment |
| kyverno.admissionController.tolerations | list | `[{"effect":"NoSchedule","key":"node.cluster.x-k8s.io/uninitialized","operator":"Exists"}]` | List of node taints to tolerate |
| kyverno.admissionController.antiAffinity.enabled | bool | `true` | Pod antiAffinities toggle. Enabled by default but can be disabled if you want to schedule pods to the same node. |
| kyverno.admissionController.podAntiAffinity | object | See [values.yaml](values.yaml) | Pod anti affinity constraints. |
| kyverno.admissionController.podAffinity | object | `{}` | Pod affinity constraints. |
| kyverno.admissionController.nodeAffinity | object | `{"preferredDuringSchedulingIgnoredDuringExecution":[{"preference":{"matchExpressions":[{"key":"karpenter.sh/capacity-type","operator":"NotIn","values":["spot"]}]},"weight":50}]}` | Node affinity constraints. |
| kyverno.admissionController.topologySpreadConstraints | list | `[]` | Topology spread constraints |
| kyverno.admissionController.podSecurityContext | object | `{}` | Security context for the pod |
| kyverno.admissionController.podDisruptionBudget.minAvailable | int | `2` | Configures the minimum available pods for disruptions. Cannot be used if `maxUnavailable` is set. |
| kyverno.admissionController.podDisruptionBudget.unhealthyPodEvictionPolicy | string | `"AlwaysAllow"` | Configures the maximum unavailable pods for disruptions. Cannot be used if `minAvailable` is set. maxUnavailable: |
| kyverno.admissionController.tufRootMountPath | string | `"/.sigstore"` | A writable volume to use for the TUF root initialization. |
| kyverno.admissionController.sigstoreVolume | object | `{"emptyDir":{}}` | Volume to be mounted in pods for TUF/cosign work. |
| kyverno.admissionController.imagePullSecrets | list | `[]` | Image pull secrets |
| kyverno.admissionController.initContainer.image.registry | string | `"gsoci.azurecr.io"` | Image registry |
| kyverno.admissionController.initContainer.image.repository | string | `"giantswarm/kyvernopre"` | Image repository |
| kyverno.admissionController.initContainer.resources.limits | object | `{"cpu":"100m","memory":"256Mi"}` | Pod resource limits |
| kyverno.admissionController.initContainer.resources.requests | object | `{"cpu":"10m","memory":"64Mi"}` | Pod resource requests |
| kyverno.admissionController.initContainer.securityContext | object | `{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"privileged":false,"readOnlyRootFilesystem":true,"runAsGroup":65534,"runAsNonRoot":true,"runAsUser":65534,"seccompProfile":{"type":"RuntimeDefault"}}` | Container security context |
| kyverno.admissionController.initContainer.extraArgs | object | `{}` | Additional container args. |
| kyverno.admissionController.initContainer.extraEnvVars | list | `[]` | Additional container environment variables. |
| kyverno.admissionController.container.image.registry | string | `"gsoci.azurecr.io"` | Image registry |
| kyverno.admissionController.container.image.repository | string | `"giantswarm/kyverno"` | Image repository |
| kyverno.admissionController.container.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| kyverno.admissionController.container.resources.limits | object | `{"memory":"384Mi"}` | Pod resource limits |
| kyverno.admissionController.container.resources.requests | object | `{"cpu":"100m","memory":"128Mi"}` | Pod resource requests |
| kyverno.admissionController.container.securityContext | object | `{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"privileged":false,"readOnlyRootFilesystem":true,"runAsGroup":65534,"runAsNonRoot":true,"runAsUser":65534,"seccompProfile":{"type":"RuntimeDefault"}}` | Container security context |
| kyverno.admissionController.container.extraArgs | object | `{}` | Additional container args. |
| kyverno.admissionController.container.extraEnvVars | list | `[]` | Additional container environment variables. |
| kyverno.admissionController.extraInitContainers | list | `[]` | Array of extra init containers |
| kyverno.admissionController.extraContainers | list | `[]` | Array of extra containers to run alongside kyverno |
| kyverno.admissionController.service.port | int | `443` | Service port. |
| kyverno.admissionController.service.type | string | `"ClusterIP"` | Service type. |
| kyverno.admissionController.service.annotations | object | `{}` | Service annotations. |
| kyverno.admissionController.metricsService.create | bool | `true` | Create service. |
| kyverno.admissionController.metricsService.port | int | `8000` | Service port. Kyverno's metrics server will be exposed at this port. |
| kyverno.admissionController.metricsService.type | string | `"ClusterIP"` | Service type. |
| kyverno.admissionController.metricsService.annotations | object | `{}` | Service annotations. |
| kyverno.admissionController.networkPolicy.enabled | bool | `false` | When true, use a NetworkPolicy to allow ingress to the webhook This is useful on clusters using Calico and/or native k8s network policies in a default-deny setup. |
| kyverno.admissionController.networkPolicy.ingressFrom | list | `[]` | A list of valid from selectors according to https://kubernetes.io/docs/concepts/services-networking/network-policies. |
| kyverno.admissionController.serviceMonitor.enabled | bool | `true` | Create a `ServiceMonitor` to collect Prometheus metrics. |
| kyverno.admissionController.serviceMonitor.additionalLabels | object | `{}` | Additional labels |
| kyverno.admissionController.serviceMonitor.interval | string | `"30s"` | Interval to scrape metrics |
| kyverno.admissionController.serviceMonitor.scrapeTimeout | string | `"25s"` | Timeout if metrics can't be retrieved in given time interval |
| kyverno.admissionController.serviceMonitor.secure | bool | `false` | Is TLS required for endpoint |
| kyverno.admissionController.serviceMonitor.tlsConfig | object | `{}` | TLS Configuration for endpoint |
| kyverno.admissionController.tracing.enabled | bool | `false` | Enable tracing |
| kyverno.admissionController.tracing.address | string | `nil` | Traces receiver address |
| kyverno.admissionController.tracing.port | string | `nil` | Traces receiver port |
| kyverno.admissionController.tracing.creds | string | `""` | Traces receiver credentials |
| kyverno.admissionController.metering.disabled | bool | `false` | Disable metrics export |
| kyverno.admissionController.metering.config | string | `"prometheus"` | Otel configuration, can be `prometheus` or `grpc` |
| kyverno.admissionController.metering.port | int | `8000` | Prometheus endpoint port |
| kyverno.admissionController.metering.collector | string | `""` | Otel collector endpoint |
| kyverno.admissionController.metering.creds | string | `""` | Otel collector credentials |
| kyverno.admissionController.profiling.enabled | bool | `false` | Enable profiling |
| kyverno.admissionController.profiling.port | int | `6060` | Profiling endpoint port |
| kyverno.admissionController.profiling.serviceType | string | `"ClusterIP"` | Service type. |
| kyverno.backgroundController.featuresOverride | object | `{}` | Overrides features defined at the root level |
| kyverno.backgroundController.enabled | bool | `true` | Enable background controller. |
| kyverno.backgroundController.rbac.create | bool | `true` | Create RBAC resources |
| kyverno.backgroundController.rbac.serviceAccount.name | string | `"kyverno-background-controller"` | Service account name |
| kyverno.backgroundController.rbac.serviceAccount.annotations | object | `{}` | Annotations for the ServiceAccount |
| kyverno.backgroundController.rbac.clusterRole.extraResources | list | `[]` | Extra resource permissions to add in the cluster role |
| kyverno.backgroundController.image.registry | string | `"gsoci.azurecr.io"` | Image registry |
| kyverno.backgroundController.image.repository | string | `"giantswarm/background-controller"` | Image repository |
| kyverno.backgroundController.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| kyverno.backgroundController.imagePullSecrets | list | `[]` | Image pull secrets |
| kyverno.backgroundController.updateStrategy | object | See [values.yaml](values.yaml) | Deployment update strategy. Ref: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#strategy |
| kyverno.backgroundController.priorityClassName | string | `"giantswarm-critical"` | Optional priority class |
| kyverno.backgroundController.hostNetwork | bool | `false` | Change `hostNetwork` to `true` when you want the pod to share its host's network namespace. Useful for situations like when you end up dealing with a custom CNI over Amazon EKS. Update the `dnsPolicy` accordingly as well to suit the host network mode. |
| kyverno.backgroundController.dnsPolicy | string | `"ClusterFirst"` | `dnsPolicy` determines the manner in which DNS resolution happens in the cluster. In case of `hostNetwork: true`, usually, the `dnsPolicy` is suitable to be `ClusterFirstWithHostNet`. For further reference: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy. |
| kyverno.backgroundController.extraArgs | object | `{}` | Extra arguments passed to the container on the command line |
| kyverno.backgroundController.resources.limits | object | `{"cpu":"750m","memory":"2Gi"}` | Pod resource limits |
| kyverno.backgroundController.resources.requests | object | `{"cpu":"500m","memory":"1Gi"}` | Pod resource requests |
| kyverno.backgroundController.nodeSelector | object | `{}` | Node labels for pod assignment |
| kyverno.backgroundController.tolerations | list | `[]` | List of node taints to tolerate |
| kyverno.backgroundController.antiAffinity.enabled | bool | `true` | Pod antiAffinities toggle. Enabled by default but can be disabled if you want to schedule pods to the same node. |
| kyverno.backgroundController.podAntiAffinity | object | See [values.yaml](values.yaml) | Pod anti affinity constraints. |
| kyverno.backgroundController.podAffinity | object | `{}` | Pod affinity constraints. |
| kyverno.backgroundController.nodeAffinity | object | `{}` | Node affinity constraints. |
| kyverno.backgroundController.topologySpreadConstraints | list | `[]` | Topology spread constraints. |
| kyverno.backgroundController.podSecurityContext | object | `{}` | Security context for the pod |
| kyverno.backgroundController.securityContext | object | `{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"privileged":false,"readOnlyRootFilesystem":true,"runAsGroup":65534,"runAsNonRoot":true,"runAsUser":65534,"seccompProfile":{"type":"RuntimeDefault"}}` | Security context for the containers |
| kyverno.backgroundController.podDisruptionBudget.minAvailable | int | `1` | Configures the minimum available pods for disruptions. Cannot be used if `maxUnavailable` is set. |
| kyverno.backgroundController.metricsService.create | bool | `true` | Create service. |
| kyverno.backgroundController.metricsService.port | int | `8000` | Service port. Metrics server will be exposed at this port. |
| kyverno.backgroundController.metricsService.type | string | `"ClusterIP"` | Service type. |
| kyverno.backgroundController.metricsService.annotations | object | `{}` | Service annotations. |
| kyverno.backgroundController.networkPolicy.enabled | bool | `false` | When true, use a NetworkPolicy to allow ingress to the webhook This is useful on clusters using Calico and/or native k8s network policies in a default-deny setup. |
| kyverno.backgroundController.networkPolicy.ingressFrom | list | `[]` | A list of valid from selectors according to https://kubernetes.io/docs/concepts/services-networking/network-policies. |
| kyverno.backgroundController.serviceMonitor.enabled | bool | `true` | Create a `ServiceMonitor` to collect Prometheus metrics. |
| kyverno.backgroundController.serviceMonitor.additionalLabels | object | `{}` | Additional labels |
| kyverno.backgroundController.serviceMonitor.interval | string | `"30s"` | Interval to scrape metrics |
| kyverno.backgroundController.serviceMonitor.scrapeTimeout | string | `"25s"` | Timeout if metrics can't be retrieved in given time interval |
| kyverno.backgroundController.serviceMonitor.secure | bool | `false` | Is TLS required for endpoint |
| kyverno.backgroundController.serviceMonitor.tlsConfig | object | `{}` | TLS Configuration for endpoint |
| kyverno.backgroundController.tracing.enabled | bool | `false` | Enable tracing |
| kyverno.backgroundController.tracing.address | string | `nil` | Traces receiver address |
| kyverno.backgroundController.tracing.port | string | `nil` | Traces receiver port |
| kyverno.backgroundController.tracing.creds | string | `""` | Traces receiver credentials |
| kyverno.backgroundController.metering.disabled | bool | `false` | Disable metrics export |
| kyverno.backgroundController.metering.config | string | `"prometheus"` | Otel configuration, can be `prometheus` or `grpc` |
| kyverno.backgroundController.metering.port | int | `8000` | Prometheus endpoint port |
| kyverno.backgroundController.metering.collector | string | `""` | Otel collector endpoint |
| kyverno.backgroundController.metering.creds | string | `""` | Otel collector credentials |
| kyverno.backgroundController.server | object | `{"port":9443}` | backgroundController server port in case you are using hostNetwork: true, you might want to change the port the backgroundController is listening to |
| kyverno.backgroundController.profiling.enabled | bool | `false` | Enable profiling |
| kyverno.backgroundController.profiling.port | int | `6060` | Profiling endpoint port |
| kyverno.backgroundController.profiling.serviceType | string | `"ClusterIP"` | Service type. |
| kyverno.cleanupController.featuresOverride | object | `{}` | Overrides features defined at the root level |
| kyverno.cleanupController.enabled | bool | `true` | Enable cleanup controller. |
| kyverno.cleanupController.rbac.create | bool | `true` | Create RBAC resources |
| kyverno.cleanupController.rbac.serviceAccount.name | string | `"kyverno-cleanup-controller"` | Service account name |
| kyverno.cleanupController.rbac.serviceAccount.annotations | object | `{}` | Annotations for the ServiceAccount |
| kyverno.cleanupController.rbac.clusterRole.extraResources | list | `[]` | Extra resource permissions to add in the cluster role |
| kyverno.cleanupController.createSelfSignedCert | bool | `false` | Create self-signed certificates at deployment time. The certificates won't be automatically renewed if this is set to `true`. |
| kyverno.cleanupController.image.registry | string | `"gsoci.azurecr.io"` | Image registry |
| kyverno.cleanupController.image.repository | string | `"giantswarm/cleanup-controller"` | Image repository |
| kyverno.cleanupController.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| kyverno.cleanupController.imagePullSecrets | list | `[]` | Image pull secrets |
| kyverno.cleanupController.updateStrategy | object | See [values.yaml](values.yaml) | Deployment update strategy. Ref: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#strategy |
| kyverno.cleanupController.priorityClassName | string | `"giantswarm-critical"` | Optional priority class |
| kyverno.cleanupController.hostNetwork | bool | `false` | Change `hostNetwork` to `true` when you want the pod to share its host's network namespace. Useful for situations like when you end up dealing with a custom CNI over Amazon EKS. Update the `dnsPolicy` accordingly as well to suit the host network mode. |
| kyverno.cleanupController.dnsPolicy | string | `"ClusterFirst"` | `dnsPolicy` determines the manner in which DNS resolution happens in the cluster. In case of `hostNetwork: true`, usually, the `dnsPolicy` is suitable to be `ClusterFirstWithHostNet`. For further reference: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy. |
| kyverno.cleanupController.extraArgs | object | `{}` | Extra arguments passed to the container on the command line |
| kyverno.cleanupController.resources.limits | object | `{"cpu":"125m","memory":"128Mi"}` | Pod resource limits |
| kyverno.cleanupController.resources.requests | object | `{"cpu":"100m","memory":"64Mi"}` | Pod resource requests |
| kyverno.cleanupController.nodeSelector | object | `{}` | Node labels for pod assignment |
| kyverno.cleanupController.tolerations | list | `[]` | List of node taints to tolerate |
| kyverno.cleanupController.antiAffinity.enabled | bool | `true` | Pod antiAffinities toggle. Enabled by default but can be disabled if you want to schedule pods to the same node. |
| kyverno.cleanupController.podAntiAffinity | object | See [values.yaml](values.yaml) | Pod anti affinity constraints. |
| kyverno.cleanupController.podAffinity | object | `{}` | Pod affinity constraints. |
| kyverno.cleanupController.nodeAffinity | object | `{}` | Node affinity constraints. |
| kyverno.cleanupController.topologySpreadConstraints | list | `[]` | Topology spread constraints. |
| kyverno.cleanupController.podSecurityContext | object | `{}` | Security context for the pod |
| kyverno.cleanupController.securityContext | object | `{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"privileged":false,"readOnlyRootFilesystem":true,"runAsGroup":65534,"runAsNonRoot":true,"runAsUser":65534,"seccompProfile":{"type":"RuntimeDefault"}}` | Security context for the containers |
| kyverno.cleanupController.podDisruptionBudget.minAvailable | int | `1` | Configures the minimum available pods for disruptions. Cannot be used if `maxUnavailable` is set. |
| kyverno.cleanupController.service.port | int | `443` | Service port. |
| kyverno.cleanupController.service.type | string | `"ClusterIP"` | Service type. |
| kyverno.cleanupController.service.annotations | object | `{}` | Service annotations. |
| kyverno.cleanupController.metricsService.create | bool | `true` | Create service. |
| kyverno.cleanupController.metricsService.port | int | `8000` | Service port. Metrics server will be exposed at this port. |
| kyverno.cleanupController.metricsService.type | string | `"ClusterIP"` | Service type. |
| kyverno.cleanupController.metricsService.annotations | object | `{}` | Service annotations. |
| kyverno.cleanupController.networkPolicy.enabled | bool | `false` | When true, use a NetworkPolicy to allow ingress to the webhook This is useful on clusters using Calico and/or native k8s network policies in a default-deny setup. |
| kyverno.cleanupController.networkPolicy.ingressFrom | list | `[]` | A list of valid from selectors according to https://kubernetes.io/docs/concepts/services-networking/network-policies. |
| kyverno.cleanupController.serviceMonitor.enabled | bool | `true` | Create a `ServiceMonitor` to collect Prometheus metrics. |
| kyverno.cleanupController.serviceMonitor.additionalLabels | object | `{}` | Additional labels |
| kyverno.cleanupController.serviceMonitor.interval | string | `"30s"` | Interval to scrape metrics |
| kyverno.cleanupController.serviceMonitor.scrapeTimeout | string | `"25s"` | Timeout if metrics can't be retrieved in given time interval |
| kyverno.cleanupController.serviceMonitor.secure | bool | `false` | Is TLS required for endpoint |
| kyverno.cleanupController.serviceMonitor.tlsConfig | object | `{}` | TLS Configuration for endpoint |
| kyverno.cleanupController.tracing.enabled | bool | `false` | Enable tracing |
| kyverno.cleanupController.tracing.address | string | `nil` | Traces receiver address |
| kyverno.cleanupController.tracing.port | string | `nil` | Traces receiver port |
| kyverno.cleanupController.tracing.creds | string | `""` | Traces receiver credentials |
| kyverno.cleanupController.metering.disabled | bool | `false` | Disable metrics export |
| kyverno.cleanupController.metering.config | string | `"prometheus"` | Otel configuration, can be `prometheus` or `grpc` |
| kyverno.cleanupController.metering.port | int | `8000` | Prometheus endpoint port |
| kyverno.cleanupController.metering.collector | string | `""` | Otel collector endpoint |
| kyverno.cleanupController.metering.creds | string | `""` | Otel collector credentials |
| kyverno.cleanupController.server | object | `{"port":9443}` | cleanupController server port in case you are using hostNetwork: true, you might want to change the port the cleanupController is listening to |
| kyverno.cleanupController.profiling.enabled | bool | `false` | Enable profiling |
| kyverno.cleanupController.profiling.port | int | `6060` | Profiling endpoint port |
| kyverno.cleanupController.profiling.serviceType | string | `"ClusterIP"` | Service type. |
| kyverno.reportsController.featuresOverride | object | `{}` | Overrides features defined at the root level |
| kyverno.reportsController.enabled | bool | `true` | Enable reports controller. |
| kyverno.reportsController.rbac.create | bool | `true` | Create RBAC resources |
| kyverno.reportsController.rbac.serviceAccount.name | string | `"kyverno-reports-controller"` | Service account name |
| kyverno.reportsController.rbac.serviceAccount.annotations | object | `{}` | Annotations for the ServiceAccount |
| kyverno.reportsController.rbac.clusterRole.extraResources | list | `[{"apiGroups":["apiextensions.k8s.io"],"resources":["customresourcedefinitions"],"verbs":["list"]}]` | Extra resource permissions to add in the cluster role |
| kyverno.reportsController.image.registry | string | `"gsoci.azurecr.io"` | Image registry |
| kyverno.reportsController.image.repository | string | `"giantswarm/reports-controller"` | Image repository |
| kyverno.reportsController.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| kyverno.reportsController.imagePullSecrets | list | `[]` | Image pull secrets |
| kyverno.reportsController.updateStrategy | object | See [values.yaml](values.yaml) | Deployment update strategy. Ref: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#strategy |
| kyverno.reportsController.priorityClassName | string | `"giantswarm-critical"` | Optional priority class |
| kyverno.reportsController.apiPriorityAndFairness | bool | `true` | Change `apiPriorityAndFairness` to `true` if you want to insulate the API calls made by Kyverno reports controller activities. This will help ensure Kyverno reports stability in busy clusters. Ref: https://kubernetes.io/docs/concepts/cluster-administration/flow-control/ |
| kyverno.reportsController.hostNetwork | bool | `false` | Change `hostNetwork` to `true` when you want the pod to share its host's network namespace. Useful for situations like when you end up dealing with a custom CNI over Amazon EKS. Update the `dnsPolicy` accordingly as well to suit the host network mode. |
| kyverno.reportsController.dnsPolicy | string | `"ClusterFirst"` | `dnsPolicy` determines the manner in which DNS resolution happens in the cluster. In case of `hostNetwork: true`, usually, the `dnsPolicy` is suitable to be `ClusterFirstWithHostNet`. For further reference: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy. |
| kyverno.reportsController.extraArgs | object | `{}` | Extra arguments passed to the container on the command line |
| kyverno.reportsController.resources.limits | object | `{"cpu":"750m","memory":"2Gi"}` | Pod resource limits |
| kyverno.reportsController.resources.requests | object | `{"cpu":"500m","memory":"1Gi"}` | Pod resource requests |
| kyverno.reportsController.nodeSelector | object | `{}` | Node labels for pod assignment |
| kyverno.reportsController.tolerations | list | `[]` | List of node taints to tolerate |
| kyverno.reportsController.antiAffinity.enabled | bool | `true` | Pod antiAffinities toggle. Enabled by default but can be disabled if you want to schedule pods to the same node. |
| kyverno.reportsController.podAntiAffinity | object | See [values.yaml](values.yaml) | Pod anti affinity constraints. |
| kyverno.reportsController.podAffinity | object | `{}` | Pod affinity constraints. |
| kyverno.reportsController.nodeAffinity | object | `{}` | Node affinity constraints. |
| kyverno.reportsController.topologySpreadConstraints | list | `[]` | Topology spread constraints. |
| kyverno.reportsController.podSecurityContext | object | `{}` | Security context for the pod |
| kyverno.reportsController.securityContext | object | `{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"privileged":false,"readOnlyRootFilesystem":true,"runAsGroup":65534,"runAsNonRoot":true,"runAsUser":65534,"seccompProfile":{"type":"RuntimeDefault"}}` | Security context for the containers |
| kyverno.reportsController.podDisruptionBudget.minAvailable | int | `1` | Configures the minimum available pods for disruptions. Cannot be used if `maxUnavailable` is set. |
| kyverno.reportsController.tufRootMountPath | string | `"/.sigstore"` | A writable volume to use for the TUF root initialization. |
| kyverno.reportsController.sigstoreVolume | object | `{"emptyDir":{}}` | Volume to be mounted in pods for TUF/cosign work. |
| kyverno.reportsController.metricsService.create | bool | `true` | Create service. |
| kyverno.reportsController.metricsService.port | int | `8000` | Service port. Metrics server will be exposed at this port. |
| kyverno.reportsController.metricsService.type | string | `"ClusterIP"` | Service type. |
| kyverno.reportsController.metricsService.annotations | object | `{}` | Service annotations. |
| kyverno.reportsController.networkPolicy.enabled | bool | `false` | When true, use a NetworkPolicy to allow ingress to the webhook This is useful on clusters using Calico and/or native k8s network policies in a default-deny setup. |
| kyverno.reportsController.networkPolicy.ingressFrom | list | `[]` | A list of valid from selectors according to https://kubernetes.io/docs/concepts/services-networking/network-policies. |
| kyverno.reportsController.serviceMonitor.enabled | bool | `true` | Create a `ServiceMonitor` to collect Prometheus metrics. |
| kyverno.reportsController.serviceMonitor.additionalLabels | object | `{}` | Additional labels |
| kyverno.reportsController.serviceMonitor.interval | string | `"30s"` | Interval to scrape metrics |
| kyverno.reportsController.serviceMonitor.scrapeTimeout | string | `"25s"` | Timeout if metrics can't be retrieved in given time interval |
| kyverno.reportsController.serviceMonitor.secure | bool | `false` | Is TLS required for endpoint |
| kyverno.reportsController.serviceMonitor.tlsConfig | object | `{}` | TLS Configuration for endpoint |
| kyverno.reportsController.tracing.enabled | bool | `false` | Enable tracing |
| kyverno.reportsController.tracing.address | string | `nil` | Traces receiver address |
| kyverno.reportsController.tracing.port | string | `nil` | Traces receiver port |
| kyverno.reportsController.tracing.creds | string | `nil` | Traces receiver credentials |
| kyverno.reportsController.metering.disabled | bool | `false` | Disable metrics export |
| kyverno.reportsController.metering.config | string | `"prometheus"` | Otel configuration, can be `prometheus` or `grpc` |
| kyverno.reportsController.metering.port | int | `8000` | Prometheus endpoint port |
| kyverno.reportsController.metering.collector | string | `nil` | Otel collector endpoint |
| kyverno.reportsController.metering.creds | string | `nil` | Otel collector credentials |
| kyverno.reportsController.server | object | `{"port":9443}` | reportsController server port in case you are using hostNetwork: true, you might want to change the port the reportsController is listening to |
| kyverno.reportsController.profiling.enabled | bool | `false` | Enable profiling |
| kyverno.reportsController.profiling.port | int | `6060` | Profiling endpoint port |
| kyverno.reportsController.profiling.serviceType | string | `"ClusterIP"` | Service type. |
| kyverno.reportsController.sanityChecks | bool | `true` | Enable sanity check for reports CRDs |
| kyverno.webhooksCleanup.enabled | bool | `true` | Create a helm pre-delete hook to cleanup webhooks. |
| kyverno.webhooksCleanup.autoDeleteWebhooks.enabled | bool | `false` | Allow webhooks controller to delete webhooks using finalizers |
| kyverno.webhooksCleanup.image.registry | string | `"gsoci.azurecr.io"` | Image registry |
| kyverno.webhooksCleanup.image.repository | string | `"giantswarm/kubectl"` | Image repository |
| kyverno.webhooksCleanup.image.tag | string | `"v1.34.3"` | Image tag Defaults to `latest` if omitted |
| policy-reporter.fullnameOverride | string | `"kyverno-policy-reporter"` |  |
| policy-reporter.podLabels."app.kubernetes.io/component" | string | `"policy-reporter"` |  |
| policy-reporter.image.registry | string | `"gsoci.azurecr.io"` |  |
| policy-reporter.image.repository | string | `"giantswarm/policy-reporter"` |  |
| policy-reporter.resources.limits.cpu | string | `"30m"` |  |
| policy-reporter.resources.limits.memory | string | `"100Mi"` |  |
| policy-reporter.resources.requests.cpu | string | `"5m"` |  |
| policy-reporter.resources.requests.memory | string | `"30Mi"` |  |
| policy-reporter.serviceAccount.name | string | `"policyreporter-sa"` |  |
| policy-reporter.ui.enabled | bool | `true` |  |
| policy-reporter.ui.podLabels."app.kubernetes.io/component" | string | `"ui"` |  |
| policy-reporter.ui.image.registry | string | `"gsoci.azurecr.io"` |  |
| policy-reporter.ui.image.repository | string | `"giantswarm/policy-reporter-ui"` |  |
| policy-reporter.ui.resources.limits.cpu | string | `"10m"` |  |
| policy-reporter.ui.resources.limits.memory | string | `"16Mi"` |  |
| policy-reporter.ui.resources.requests.cpu | string | `"1m"` |  |
| policy-reporter.ui.resources.requests.memory | string | `"8Mi"` |  |
| policy-reporter.ui.serviceAccount.create | bool | `true` |  |
| policy-reporter.ui.serviceAccount.name | string | `"policyreporter-ui-sa"` |  |
| policy-reporter.plugin.kyverno.enabled | bool | `true` |  |
| policy-reporter.plugin.kyverno.podLabels."app.kubernetes.io/component" | string | `"plugin"` |  |
| policy-reporter.plugin.kyverno.image.registry | string | `"gsoci.azurecr.io"` |  |
| policy-reporter.plugin.kyverno.image.repository | string | `"giantswarm/kyverno-plugin"` |  |
| policy-reporter.plugin.kyverno.resources.limits.cpu | string | `"50m"` |  |
| policy-reporter.plugin.kyverno.resources.limits.memory | string | `"100Mi"` |  |
| policy-reporter.plugin.kyverno.resources.requests.cpu | string | `"10m"` |  |
| policy-reporter.plugin.kyverno.resources.requests.memory | string | `"30Mi"` |  |
| policy-reporter.global.labels | object | `{}` |  |
| policy-reporter.monitoring.enabled | bool | `true` |  |
| policy-reporter.monitoring.serviceMonitor.namespace | string | `"kyverno"` |  |
| policy-reporter.monitoring.serviceMonitor.relabelings[0].action | string | `"labeldrop"` |  |
| policy-reporter.monitoring.serviceMonitor.relabelings[0].regex | string | `"pod|service|container"` |  |
| policy-reporter.monitoring.serviceMonitor.relabelings[1].targetLabel | string | `"instance"` |  |
| policy-reporter.monitoring.serviceMonitor.relabelings[1].replacement | string | `"policy-reporter"` |  |
| policy-reporter.monitoring.serviceMonitor.relabelings[1].action | string | `"replace"` |  |
