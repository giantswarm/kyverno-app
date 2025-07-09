# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Enable observability platform log collection.

## [0.20.0] - 2025-06-12

### Changed

- Update `kyverno` to upstream version v1.14.2.
- Update `kyverno-policy-reporter` to upstream version v3.1.0.
- Update PolicyExceptions from `v2beta1` to `v2`.
- Add karpenter spot anti-affinity to kyverno admission controller.
- Update PolicyException for Cilium.

### Notes

This release includes an upstream update. Please refer to the following Release Notes from upstream for the latest changes:

- [v1.14.0](https://github.com/kyverno/kyverno/releases/tag/v1.14.0)
- [v1.14.1](https://github.com/kyverno/kyverno/releases/tag/v1.14.1)
- [v1.14.2](https://github.com/kyverno/kyverno/releases/tag/v1.14.2)

## [0.19.0] - 2025-02-20

### Changed

- Update `kyverno` to upstream version v1.13.4.
- Use GVK for specifying Kinds in core-policies.
- Add `runAsGroup` to container security contexts.

## [0.18.1] - 2024-10-01

### Changed

- Update `Kyverno` to upstream version v1.12.6.
- Update `kyverno-policy-reporter` to upstream version v2.20.2.

## [0.18.0] - 2024-09-25

### Changed

- Update `Kyverno` to upstream version v1.12.5.

## [0.17.16] - 2024-08-29

### Changed

- Split Cilium PolicyExceptions per component.
- Add rules to cilium-agent PolicyException.
- Restrict Policy and ClusterPolicy to kyverno.io/v1 for wildcard policy matching

### Removed

- Remove Helm `hooks` annotations from default Policies and PolicyExceptions.

## [0.17.15] - 2024-07-18

### Changed

- Set VPA max 6 CPU / 24Gi memory and adjust default requests/limits for `reports-controller`.
- Set VPA max 4 CPU / 8Gi memory and adjust default requests/limits for `background-controller`.
- Set starting CPU limit of request+25% for `cleanup-controller`.

### Removed

- Disable Kyverno CRDs install Job in favor of `kyverno-crds` App.

## [0.17.14] - 2024-06-12

### Changed

- Remove duplicated key from the `allow-kyverno-policy-reporter-talk-to-kyverno-ui` CiliumNetworkPolicy.
- Add Helm annotations to core policies.

## [0.17.13] - 2024-05-23

### Changed

- Disable Trivy cleanup policy by default.

## [0.17.12] - 2024-05-22

### Changed

- Fix template issue with DNS rules on the `admission-controller` CiliumNetworkPolicy.

## [0.17.11] - 2024-05-16

### Added

- Add cleanup policy to remove old `trivy-operator` resources.

### Changed

- Enable `cleanup-controller` with VericalPodAutoscaler by default.
- Add missing ingress to `cleanup-controller` CiliumNetworkPolicy.
- Add `before-hook-creation` delete-policy for upstream hooks.

## [0.17.10] - 2024-04-30

### Added

- Add Helm labels and annotations for easy CRD adoption in the future.

### Changed

- Adapt Kyverno Policy Reporter CiliumNetworkPolicy to allow for DNS resolution of the `kyverno-ui` service.
- Disable AdmissionReports and ClusterAdmissionReports cleanup jobs.

## [0.17.9] - 2024-04-03

### Changed

- Revert `kubectl` images to v1.29.2.

## [0.17.8] - 2024-04-03

### Added

- Add new CiliumNetworkPolicy section to allow for DNS and FQDNs connections.

### Changed

- Don't push to vsphere-app-collection, capz-app-collection, capa-app-collection or cloud-director-app-collection. We started to consume kyverno-app from security-bundle.

## [0.17.7] - 2024-02-28

### Changed

- Disable v1.10 upgrade jobs. They are not needed when migrating from v1.10 to v1.11+.

## [0.17.6] - 2024-02-22

### Fixed

- Add missing ingress to Cleanup Controller CiliumNetworkPolicy.

## [0.17.5] - 2024-02-08

### Fixed

- Add missing CiliumNetworkPolicies for pre-delete and post-ugprade hooks.

## [0.17.4] - 2024-02-07

### Changed

- Fix label selector `kyverno-policy-reporter` to talk to `kyverno-ui` rule.
- Add `policy-exceptions` namespace if it doesn't exist.

## [0.17.3] - 2024-01-26

### Changed

- Allow `kyverno-policy-reporter` to talk to `kyverno-ui`.

## [0.17.2] - 2024-01-25

### Changed

- Enable CiliumNetworkPolicies by default.
- Enable API Priority and Fairness.

## [0.17.1] - 2024-01-25

### Changed

- Bump to upstream version v1.11.4.

## [0.17.0] - 2024-01-15

### Changed

- Bump to upstream version v1.11.2.
- Enable the Policy Reporter subchart by default.

## [0.16.4] - 2023-12-06

### Added

- Add missing `app.kubernetes.io/` labels to all the pods.
- Add `CiliumNetworkPolicy` for individual controllers:
  - `kyverno-admission-controller`
  - `kyverno-background-controller`
  - `kyverno-reports-controller`
  - `kyverno-cleanup-controller`
  - `kyverno-cleanup-jobs`
  - `kyverno-plugin`
  - `kyverno-policy-reporter`

## [0.16.3] - 2023-11-29

### Added

- Added Policy Exceptions for `azure-cloud-node-manager`.

## [0.16.2] - 2023-11-16

### Changed

- Change the `kubectl apply` command of the `crd-install` job to use the `--force-conflicts` flag.

## [0.16.1] - 2023-10-25

### Changed

- Removed unused values and update schema.

## [0.16.0] - 2023-10-24

### Added

- Added Policy Exceptions for `aws-cloud-controller-manager`, `aws-ebs-csi-driver`, `azure-cloud-controller-manager` and `cilium`.
- Add Policy Exception for `chart-operator` ServiceAccount.
- Change `psp.enabled` value to `global.podSecurityStandards.enforced`

## [0.15.2] - 2023-09-01

### Changed

- Removed check capabilities from VerticalPodAutoscaler resources.
- Dropped `CPU` from Reports Controller VerticalPodAutoscaler controlled resources.

## [0.15.1] - 2023-08-16

### Changed

- Added non-root UID/GID for the cleanup jobs.
- Re-enable `PolicyViolation` events for all controllers.

## [0.15.0] - 2023-08-16

### Added

- Updated Kyverno to upstream version v1.10.2.
- Add core Policy `disallow-noisy-policy-contexts`.

## [0.14.10] - 2023-07-04

### Added

- Add cilium network policies support (enabled with `ciliumNetworkPolicy.enabled=true`).

### Changed

- Allow PolicyExceptions in `flux-giantswarm` namespace.

## [0.14.9] - 2023-06-26

### Added

- Add push to `capz-app-collection` on release to circleci
- Make `scrapeTimeout` and `interval` fields configurable on Policy Reporter monitoring ServiceMonitors.

## [0.14.8] - 2023-06-15

### Changed

- Modified the PSPs logic and moved the value to the parent chart.

## [0.14.7] - 2023-06-09

### Added

- Add resourceFilter for excluding Giant Swarm's `chart-operator` from custom policies.

## [0.14.6] - 2023-06-06

### Added

- Add ClusterPolicy `restrict-policy-kind-wildcards` to prevent running (Cluster)Policies which match all API Kinds.
- Add PolicyException for Giant Swarm's `chart-operator`.

## [0.14.5] - 2023-05-16

### Added

- Add a webhooks cleanup job for ensuring deletion of Kyverno webhooks on chart uninstall.

### Changed

- Replace deprecated toleration `node-role.kubernetes.io/master` with `node-role.kubernetes.io/control-plane` on `CRD` install job.

## [0.14.4] - 2023-05-03

### Added

- Add Kyverno policy controlling `PolicyException` namespaces.
- Set a default namespace of `policy-exceptions` for customer `PolicyException` resources.

## [0.14.3] - 2023-04-11

### Added

- Adds `VerticalPodAutoscaler` for `kyverno-plugin` deployment.

### Changed

- `VericalPodAutoscaler` can now be enabled for individual components.
- Removes GiantSwarm specific labels from `monitoring` ServiceMonitors.
- Update `kyverno-policy-reporter` to upstream version 2.14.0 / chart version 2.18.0.

## [0.14.2] - 2023-03-28

### Changed

- Don't push to `openstack-app-collection`.
- Rename `vmware-app-collection` to `vsphere-app-collection`.
- Consider PolicyExceptions from all namespaces.

## [0.14.1] - 2023-03-22

### Changed

- Update `kyverno` to upstream version 1.9.2 / chart version 2.7.2.

## [0.14.0] - 2023-02-23

### Changed

- Update `kyverno` to upstream version 1.9.0 / chart version 2.7.0.
- Update `kyverno-policy-reporter` to upstream version 2.12.0 / chart version 2.16.0.
- Adds `giantswarm.io/monitoring` annotation to kyverno service & plugin.

## [0.13.1] - 2022-12-21

### Changed

- Excludes `kube-system` namespace from webhooks.
- Bump `kyverno-plugin` resources.

## [0.13.0] - 2022-12-13

### Changed

- Update `kyverno` to upstream version 1.8.4 / chart version 2.6.4.
- Update `kyverno-policy-reporter` to upstream version 2.11.0 / chart version 2.14.0.

## [0.12.2] - 2022-11-29

### Added

- Added `VerticalPodAutoscaler` to `policy-reporter`.

## [0.12.1] - 2022-11-22

### Changed

- Increased `kyverno-policy-reporter` resources limits.

## [0.12.0] - 2022-11-22

### Changed

- Update `kyverno` to upstream version 1.8.2 / chart version 2.6.2.
- Update `kyverno-policy-reporter` to upstream version 2.10.3 / chart version 2.13.4.

## [0.11.8] - 2022-11-09

### Added

- Added `CiliumNetworkPolicy` for the CRD install job.

## [0.11.7] - 2022-11-07

### Fixed

- Fix broken links in documentation.
- Don't remove the `crd-install` job when the job fails, so that we can investigate.

## [0.11.6] - 2022-10-19

### Changed

- Push `kyverno` to the cloud-director app collection.
- Update `kyverno-policy-reporter` to upstream version 2.10.1 / chart version 2.13.1.

## [0.11.5] - 2022-10-11

### Changed

- Change namespace on GCP management clusters to `kyverno`.

## [0.11.4] - 2022-10-11

### Added

- Push `kyverno` to the GCP app collection.

### Changed

- Add Service and Deployment annotations for Giant Swarm platform monitoring.
- Change target namespace for Giant Swarm management clusters.
- Build with `app-build-suite`.
- Add `app-test-suite` tests.

## [0.11.3] - 2022-10-07

### Changed

- Update `kyverno` to upstream version 1.7.4 / chart version 2.5.4.
- Update `kyverno-policy-reporter` to upstream version 2.10.0 / chart version 2.13.0.

## [0.11.2] - 2022-09-26

### Changed

- Update `kyverno` to upstream version 1.7.3 / chart version 2.5.3.

## [0.11.1] - 2022-08-23

### Changed

- Increase maximum sustained and burst Kubernetes client rate limits to 75 and 150 requests/second, respectively.
- Update `policy-reporter` to v2.11.1 / app v2.8.0.

## [0.11.0] - 2022-08-17

### Changed

- Update `kyverno` to upstream version 1.7.2 / chart version 2.5.2.
- Use pre-install CRD install Job to remove storage version `v1alpha1` for several Kyverno CRDs.
- Set Kyverno to use the `giantswarm-critical` PriorityClass.
- Limit maximum ReportChangeRequests per namespace to 100.
- Split PolicyReports into one report per policy to support the RCR limiting and avoid cases where a report doesn't fit into etcd.

## [0.10.3] - 2022-08-09

### Changed

- Fix Circle CI job name for pushing to app collection.

## [0.10.2] - 2022-08-05

### Changed

- Update `policy-reporter` to version 2.9.1 / app version 2.6.1.

## [0.10.1] - 2022-06-09

### Fixed

- Update the kyverno CRDS to align with v1.6.2 version.

## [0.10.0] - 2022-04-05

### Changed

- Break `kyverno-policies` chart into its own [separate app](https://github.com/giantswarm/kyverno-policies).
- Change `kyverno` chart to a subtree tracking the upstream [`kyverno` chart](https://github.com/kyverno/kyverno/tree/main/charts/kyverno).

## [0.9.1] - 2022-03-14

### Fixed

- Change to using `maxUnavailable` to match existing configs.

## [0.9.0] - 2022-03-14

### Changed

- Enable Pod Disruption Budget by default.

## [0.8.1] - 2022-02-18

- Make PDB version conditional based on available API.

## [0.8.0] - 2022-02-03

### Removed

- Disable policy-reporter monitoring (ServiceMonitors) by default.

## [0.7.1] - 2021-12-17

### Removed

- Transition job cleanup.

## [0.7.0] - 2021-11-08

### Changed

- Update `kyverno` to version 1.5.0.

## [0.6.3] - 2021-10-27

### Changed

- Add appropriate labels to CRDs.

## [0.6.2] - 2021-10-18

### Changed

- Change icon URL to an SVG from our own server

## [0.6.1] - 2021-09-16

### Changed

- Use Giant Swarm retagged images.

## [0.6.0] - 2021-09-16

### Added

- Add `policy-reporter` UI and monitoring.

## [0.5.0] - 2021-09-09

### Added

- Use vertical pod autoscaling (VPA)

## [0.4.0] - 2021-08-26

### Added

- Add annotation which sets `ludacris` as app owner.
- Allow ingress on port 8000 for scraping metrics.
- Add ServiceMonitor for metrics to be scraped.
- Push `kyverno` to `giantswarm` catalog.

## [0.3.0] - 2021-07-09

### Added

- Add annotation to use `config-controller`.

### Changed

- \[CI/CD\] Push kyverno to app collections.
- Update Kyverno to `v1.4.1`.

## [0.2.0] - 2021-06-25

### Changed

- Disable default policies.

## [0.1.0] - 2021-06-24

### Changed

- Enable high availability (3 replicas).
- Update Kyverno to `1.4.0`.
- Reduce memory limit.

## [0.0.5] - 2021-06-03

### Changed

- Update Kyverno to `1.3.6`

## [0.0.4] - 2021-05-24

## [0.0.3] - 2021-03-24

## [0.0.2] - 2021-03-23

## [0.0.1] - 2021-03-19

[Unreleased]: https://github.com/giantswarm/kyverno-app/compare/v0.20.0...HEAD
[0.20.0]: https://github.com/giantswarm/kyverno-app/compare/v0.19.0...v0.20.0
[0.19.0]: https://github.com/giantswarm/kyverno-app/compare/v0.18.1...v0.19.0
[0.18.1]: https://github.com/giantswarm/kyverno-app/compare/v0.18.0...v0.18.1
[0.18.0]: https://github.com/giantswarm/kyverno-app/compare/v0.17.16...v0.18.0
[0.17.16]: https://github.com/giantswarm/kyverno-app/compare/v0.17.15...v0.17.16
[0.17.15]: https://github.com/giantswarm/kyverno-app/compare/v0.17.14...v0.17.15
[0.17.14]: https://github.com/giantswarm/kyverno-app/compare/v0.17.13...v0.17.14
[0.17.13]: https://github.com/giantswarm/kyverno-app/compare/v0.17.12...v0.17.13
[0.17.12]: https://github.com/giantswarm/kyverno-app/compare/v0.17.11...v0.17.12
[0.17.11]: https://github.com/giantswarm/kyverno-app/compare/v0.17.10...v0.17.11
[0.17.10]: https://github.com/giantswarm/kyverno-app/compare/v0.17.9...v0.17.10
[0.17.9]: https://github.com/giantswarm/kyverno-app/compare/v0.17.8...v0.17.9
[0.17.8]: https://github.com/giantswarm/kyverno-app/compare/v0.17.7...v0.17.8
[0.17.7]: https://github.com/giantswarm/kyverno-app/compare/v0.17.6...v0.17.7
[0.17.6]: https://github.com/giantswarm/kyverno-app/compare/v0.17.5...v0.17.6
[0.17.5]: https://github.com/giantswarm/kyverno-app/compare/v0.17.4...v0.17.5
[0.17.4]: https://github.com/giantswarm/kyverno-app/compare/v0.17.3...v0.17.4
[0.17.3]: https://github.com/giantswarm/kyverno-app/compare/v0.17.2...v0.17.3
[0.17.2]: https://github.com/giantswarm/kyverno-app/compare/v0.17.1...v0.17.2
[0.17.1]: https://github.com/giantswarm/kyverno-app/compare/v0.17.0...v0.17.1
[0.17.0]: https://github.com/giantswarm/kyverno-app/compare/v0.16.4...v0.17.0
[0.16.4]: https://github.com/giantswarm/kyverno-app/compare/v0.16.3...v0.16.4
[0.16.3]: https://github.com/giantswarm/kyverno-app/compare/v0.16.2...v0.16.3
[0.16.2]: https://github.com/giantswarm/kyverno-app/compare/v0.16.1...v0.16.2
[0.16.1]: https://github.com/giantswarm/kyverno-app/compare/v0.16.0...v0.16.1
[0.16.0]: https://github.com/giantswarm/kyverno-app/compare/v0.15.2...v0.16.0
[0.15.2]: https://github.com/giantswarm/kyverno-app/compare/v0.15.1...v0.15.2
[0.15.1]: https://github.com/giantswarm/kyverno-app/compare/v0.15.0...v0.15.1
[0.15.0]: https://github.com/giantswarm/kyverno-app/compare/v0.14.10...v0.15.0
[0.14.10]: https://github.com/giantswarm/kyverno-app/compare/v0.14.9...v0.14.10
[0.14.9]: https://github.com/giantswarm/kyverno-app/compare/v0.14.8...v0.14.9
[0.14.8]: https://github.com/giantswarm/kyverno-app/compare/v0.14.7...v0.14.8
[0.14.7]: https://github.com/giantswarm/kyverno-app/compare/v0.14.6...v0.14.7
[0.14.6]: https://github.com/giantswarm/kyverno-app/compare/v0.14.5...v0.14.6
[0.14.5]: https://github.com/giantswarm/kyverno-app/compare/v0.14.4...v0.14.5
[0.14.4]: https://github.com/giantswarm/kyverno-app/compare/v0.14.3...v0.14.4
[0.14.3]: https://github.com/giantswarm/kyverno-app/compare/v0.14.2...v0.14.3
[0.14.2]: https://github.com/giantswarm/kyverno-app/compare/v0.14.1...v0.14.2
[0.14.1]: https://github.com/giantswarm/kyverno-app/compare/v0.14.0...v0.14.1
[0.14.0]: https://github.com/giantswarm/kyverno-app/compare/v0.13.1...v0.14.0
[0.13.1]: https://github.com/giantswarm/kyverno-app/compare/v0.13.0...v0.13.1
[0.13.0]: https://github.com/giantswarm/kyverno-app/compare/v0.12.2...v0.13.0
[0.12.2]: https://github.com/giantswarm/kyverno-app/compare/v0.12.1...v0.12.2
[0.12.1]: https://github.com/giantswarm/kyverno-app/compare/v0.12.0...v0.12.1
[0.12.0]: https://github.com/giantswarm/kyverno-app/compare/v0.11.8...v0.12.0
[0.11.8]: https://github.com/giantswarm/kyverno-app/compare/v0.11.7...v0.11.8
[0.11.7]: https://github.com/giantswarm/kyverno-app/compare/v0.11.6...v0.11.7
[0.11.6]: https://github.com/giantswarm/kyverno-app/compare/v0.11.5...v0.11.6
[0.11.5]: https://github.com/giantswarm/kyverno-app/compare/v0.11.4...v0.11.5
[0.11.4]: https://github.com/giantswarm/kyverno-app/compare/v0.11.3...v0.11.4
[0.11.3]: https://github.com/giantswarm/kyverno-app/compare/v0.11.2...v0.11.3
[0.11.2]: https://github.com/giantswarm/kyverno-app/compare/v0.11.1...v0.11.2
[0.11.1]: https://github.com/giantswarm/kyverno-app/compare/v0.11.0...v0.11.1
[0.11.0]: https://github.com/giantswarm/kyverno-app/compare/v0.10.3...v0.11.0
[0.10.3]: https://github.com/giantswarm/kyverno-app/compare/v0.10.2...v0.10.3
[0.10.2]: https://github.com/giantswarm/kyverno-app/compare/v0.10.1...v0.10.2
[0.10.1]: https://github.com/giantswarm/kyverno-app/compare/v0.10.0...v0.10.1
[0.10.0]: https://github.com/giantswarm/kyverno-app/compare/v0.9.1...v0.10.0
[0.9.1]: https://github.com/giantswarm/kyverno-app/compare/v0.9.0...v0.9.1
[0.9.0]: https://github.com/giantswarm/kyverno-app/compare/v0.8.1...v0.9.0
[0.8.1]: https://github.com/giantswarm/kyverno-app/compare/v0.8.0...v0.8.1
[0.8.0]: https://github.com/giantswarm/kyverno-app/compare/v0.7.1...v0.8.0
[0.7.1]: https://github.com/giantswarm/kyverno-app/compare/v0.7.0...v0.7.1
[0.7.0]: https://github.com/giantswarm/kyverno-app/compare/v0.6.3...v0.7.0
[0.6.3]: https://github.com/giantswarm/kyverno-app/compare/v0.6.2...v0.6.3
[0.6.2]: https://github.com/giantswarm/kyverno-app/compare/v0.6.1...v0.6.2
[0.6.1]: https://github.com/giantswarm/kyverno-app/compare/v0.6.0...v0.6.1
[0.6.0]: https://github.com/giantswarm/kyverno-app/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/giantswarm/kyverno-app/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/giantswarm/kyverno-app/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/kyverno-app/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/kyverno-app/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/kyverno-app/compare/v0.0.5...v0.1.0
[0.0.5]: https://github.com/giantswarm/kyverno-app/compare/v0.0.4...v0.0.5
[0.0.4]: https://github.com/giantswarm/kyverno-app/compare/v0.0.3...v0.0.4
[0.0.3]: https://github.com/giantswarm/kyverno-app/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/giantswarm/kyverno-app/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/giantswarm/kyverno-app/releases/tag/v0.0.1
