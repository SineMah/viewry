package finish

import (
	"bytes"
	"github.com/viewry/data"
	"github.com/viewry/data/layout"
	"os"
	"text/template"
)

func Render(c data.Contact, m data.Meta, p data.Presentation, l *layout.Layout) (string, error) {
	d := data.Overview{
		Title:          p.Title,
		Description:    p.Description,
		ShowAuthorMeta: p.Config.ShowAuthorMeta,
		Author:         m.Author,
		Mail:           c.Mail,
		Tags:           p.Tags,
	}

	tmpl, err := template.New("finish").Parse(l.Finish)

	if err != nil {
		return "", err
	}

	var renderedOutput bytes.Buffer
	err = tmpl.Execute(&renderedOutput, d)

	if err != nil {
		return "", err
	}

	return renderedOutput.String(), nil
}

func loadFile() string {
	path, _ := os.Getwd()
	file := path + "/layout/finish.md"

	t, err := os.ReadFile(file)

	if err != nil {
		return ""
	}

	return string(t)
}
