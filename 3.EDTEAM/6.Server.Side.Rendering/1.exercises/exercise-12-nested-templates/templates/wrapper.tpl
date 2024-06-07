{{ define "wrapper" }}
This is a header

{{ template "body" . }}

This will be the footer
{{ end }}