kind-create:
	kind create cluster --name kyverno --image=kindest/node:${KUBERNETES_VERSION}

kind-get-kubeconfig:
	kind get kubeconfig --name kyverno > kyverno-kubeconfig.yaml

install-kyverno:
	helm install https://giantswarm.github.io/giantswarm-catalog/kyverno-crds-${KYVERNO_CRDS_VERSION}.tgz --generate-name