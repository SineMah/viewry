package overview

import (
	"bytes"
	"github.com/viewry/data"
	"os"
	"text/template"
)

func Render(c data.Contact, m data.Meta, p data.Presentation) (string, error) {
	d := data.Overview{
		Title:          p.Title,
		Description:    p.Description,
		ShowAuthorMeta: p.Config.ShowAuthorMeta,
		ShowTitle:      p.Config.ShowTitle,
		Author:         m.Author,
		Mail:           c.Mail,
		Tags:           p.Tags,
	}

	tmpl, err := template.New("overview").Parse(loadFile())

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
	file := path + "/layout/overview.md"

	t, err := os.ReadFile(file)

	if err != nil {
		return ""
	}

	return string(t)
}
