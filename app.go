package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/viewry/data"
	"github.com/viewry/data/style"
	"github.com/viewry/display"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {
	var file string
	debug := false

	if os.Getenv("DEBUG_VIEWRY") == "true" {
		debug = true
	}

	if debug == true {
		file = exampleFile()
	}

	if len(os.Args) == 2 {
		file = os.Args[1]
	}

	g, err := loadFile(file)

	if err != nil {
		log.Fatal(fmt.Printf("Could not load file: %s", file))
	}

	s := style.New()

	g.Presentation.Config.Style = s.GetStyle(g.Presentation.Config.Style)

	m := display.Model{
		Slides:          g.Slides,
		Contact:         g.Contact,
		Meta:            g.Meta,
		Presentation:    g.Presentation,
		CurrentPosition: 0,
		CurrentAsset:    0,
		TotalSlides:     float64(len(g.Slides)),
		Viewport:        viewport.New(50, 10),
		Progress: progress.New(
			progress.WithScaledGradient(
				s.GetProgressColor1(g.Presentation.Config.ProgressColor1),
				s.GetProgressColor2(g.Presentation.Config.ProgressColor2),
			),
		),
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func loadFile(file string) (data.General, error) {
	content, err := os.ReadFile(file)

	var g data.General

	if err != nil {
		return g, err
	}

	err = yaml.Unmarshal(content, &g)
	if err != nil {
		log.Fatalf("struct error: %v", err)
	}

	return g, nil
}

func exampleFile() string {
	dir, err := os.Getwd()

	if err != nil {
		return ""
	}

	return dir + "/examples/test.yml"
}
