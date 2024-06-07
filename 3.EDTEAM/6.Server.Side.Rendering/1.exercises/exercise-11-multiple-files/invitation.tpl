Hello {{ .Name }}, hope you are doing good

{{ .Name }}, remember there is a party this Saturday.

Since you are {{ .Age }} years
{{- if lt .Age 18 }} you cannot enter, but we will send you a gift
{{- else }} we are counting on you for the party, it starts at 8pm
{{ end }}
See you soon!
