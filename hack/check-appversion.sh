#!/usr/bin/env bash
# The chart's appVersion is the Kyverno version its vendored subchart deploys.
set -euo pipefail

chart_dir="${1:-helm/kyverno}"
top="$(yq '.appVersion' "${chart_dir}/Chart.yaml")"
sub="$(yq '.appVersion' "${chart_dir}/charts/kyverno/Chart.yaml")"

if [[ "${top}" != "${sub}" ]]; then
  echo "${chart_dir}/Chart.yaml appVersion is ${top}, but the kyverno subchart deploys ${sub}: set appVersion to ${sub}" >&2
  exit 1
fi
echo "appVersion ${top} matches the kyverno subchart"
