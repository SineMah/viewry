package slide

import (
	"bytes"
	"github.com/viewry/data"
	"os"
	"text/template"
)

func Render(m data.Meta, p data.Presentation, s data.Slide, currentPage int, totalPages int) (string, error) {
	d := data.Page{
		Title:          s.Title,
		Content:        s.Content,
		ShowPageStatus: p.Config.ShowPageStatus,
		ShowAuthor:     p.Config.ShowAuthor,
		ShowImage:      len(s.Assets) > 0,
		Author:         m.Author,
		CurrentPage:    currentPage,
		TotalPages:     totalPages,
	}

	tmpl, err := template.New("slide").Parse(loadFile())

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
	file := path + "/layout/slide.md"

	t, err := os.ReadFile(file)

	if err != nil {
		return ""
	}

	return string(t)
}
