[![CircleCI](https://circleci.com/gh/giantswarm/kyverno-app.svg?style=shield)](https://circleci.com/gh/giantswarm/kyverno-app)

# kyverno

Kyverno is an admission controller offering policy enforcement as a validating or mutating webhook.
It audits or enforces policies for cluster resources, and produces reports about the compliance of the cluster.

It is used to enforce [Pod Security Standards (PSS)][pss-policies] as a replacement for Pod Security Policies (PSPs), as well as many other community-supported policies for various use cases. For more information on the switch from PSP to PSS, see [our blog post][pss-blog].

## Installing

There are 3 ways to install this app onto a workload cluster.

1. [Using our web interface](https://docs.giantswarm.io/ui-api/web/app-platform/#installing-an-app)
2. [Using our API](https://docs.giantswarm.io/api/#operation/createClusterAppV5)
3. Directly creating the [App custom resource](https://docs.giantswarm.io/ui-api/management-api/crd/apps.application.giantswarm.io/) on the management cluster.

## Configuring

#### Kyverno Configurations

Please see the [Kyverno docs][kyverno-docs] or the [configuration reference in this chart](helm/kyverno/#configuration) for configurable values.

See our [full reference page on how to configure applications](https://docs.giantswarm.io/app-platform/app-configuration/) for more details.

## Credit

* https://github.com/kyverno/kyverno

[kyverno-docs]: https://kyverno.io/docs/
[pss-blog]: https://www.giantswarm.io/blog/giant-swarms-farewell-to-psp
[pss-policies]: https://kyverno.io/policies/?policytypes=Pod%2520Security%2520Standards%2520%28Baseline%29%2BPod%2520Security%2520Standards%2520%28Restricted%29
