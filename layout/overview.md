{{if .ShowTitle}}# {{ .Title }}{{end}}

### {{ .Description }}

{{if .ShowAuthorMeta}}
```
{{ .Author }}
{{ .Mail }}
```
{{end}}

> {{- range $index, $tag := .Tags -}} {{- if $index}}, {{end}}{{ $tag }} {{- end }}