kind-create:
	kind create cluster --name kyverno --image=kindest/node:${KUBERNETES_VERSION} --config=tests/ats/kind_config.yaml

kind-get-kubeconfig:
	kind get kubeconfig --name kyverno > kyverno-kubeconfig.yaml

install-vpa-crds:
	kubectl --kubeconfig=kyverno-kubeconfig.yaml apply -f https://raw.githubusercontent.com/kubernetes/autoscaler/vertical-pod-autoscaler-1.3.0/vertical-pod-autoscaler/deploy/vpa-v1-crd-gen.yaml

install-kyverno:
	helm install https://giantswarm.github.io/giantswarm-catalog/kyverno-crds-${KYVERNO_CRDS_VERSION}.tgz --generate-name --kubeconfig=kyverno-kubeconfig.yaml