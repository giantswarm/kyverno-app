{{/* vim: set filetype=mustache: */}}

{{- define "kyverno.deployment.replicas" -}}
  {{- if not (eq . nil) -}}
  {{- if not (kindIs "string" .) -}}
  {{- if eq (int .) 0 -}}
    {{- fail "Kyverno does not support running with 0 replicas. Please provide a non-zero integer value." -}}
  {{- end -}}
  {{- end -}}
  {{- end -}}
  {{- . -}}
{{- end -}}
