# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/giantswarm/kyverno-app/compare/v0.14.4...HEAD
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
