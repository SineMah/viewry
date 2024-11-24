package layout

var overview = "{{if .ShowTitle}}# {{ .Title }}{{end}}\n\n### {{ .Description }}\n\n{{if .ShowAuthorMeta}}\n```\n{{ .Author }}\n{{ .Mail }}\n```\n{{end}}\n\n> {{- range $index, $tag := .Tags -}} {{- if $index}}, {{end}}{{ $tag }} {{- end }}"
var slide = "# {{ .Title }} {{if .ShowImage}}◢{{end}}\n\n{{ .Content }}\n\n"
var finish = "# End\n\n"

type Layout struct {
	Overview string
	Slide    string
	Finish   string
}

func New() *Layout {

	return &Layout{
		Overview: overview,
		Slide:    slide,
		Finish:   finish,
	}
}

func (l *Layout) Load() {
	// TODO Load custom layout from files
}
