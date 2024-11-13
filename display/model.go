package display

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/viewry/data"
	"github.com/viewry/display/finish"
	"github.com/viewry/display/overview"
	"github.com/viewry/display/slide"
	"os/exec"
)

const (
	padding  = 2
	maxWidth = 80
)

type Model struct {
	Slides          []data.Slide
	Contact         data.Contact
	Meta            data.Meta
	Presentation    data.Presentation
	CurrentPosition int
	CurrentAsset    int
	TotalSlides     float64
	Width           int
	Height          int
	Progress        progress.Model
	Viewport        viewport.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case "right":
			m.CurrentAsset = 0
			m.CurrentPosition++

			if m.CurrentPosition > len(m.Slides) {
				m.CurrentPosition = -1
			}
		case "left":
			m.CurrentAsset = 0
			m.CurrentPosition--

			if m.CurrentPosition < -1 {
				m.CurrentPosition = len(m.Slides)
			}
		case "tab":

			if m.CurrentPosition > 0 {
				openAsset(m)

				m.CurrentAsset++

				if m.CurrentAsset >= len(m.Slides[m.CurrentPosition-1].Assets) {
					m.CurrentAsset = 0
				}
			}
		}

	case tea.WindowSizeMsg:
		m.Viewport.Height = msg.Height - lipgloss.Height(getFooter(m))
		m.Viewport.Width = maxWidth
		m.Progress.Width = msg.Width - padding*2 - 4
		m.Width = msg.Width
		m.Height = msg.Height

		m.Viewport.SetContent(getContent(m))
		m.Viewport.GotoTop()
	}

	m.Viewport.SetContent(getContent(m))
	m.Viewport.GotoTop()

	return m, nil
}

func (m Model) View() string {
	content := lipgloss.NewStyle().
		Width(m.Width).
		Align(lipgloss.Center).
		Render(m.Viewport.View())

	return lipgloss.JoinVertical(
		lipgloss.Left,
		content,
		getFooter(m),
	)
}

func getContent(m Model) string {
	content := ""

	switch m.CurrentPosition {
	case -1:
		content, _ = finish.Render(m.Contact, m.Meta, m.Presentation)
	case 0:
		content, _ = overview.Render(m.Contact, m.Meta, m.Presentation)
	default:
		content, _ = slide.Render(
			m.Meta,
			m.Presentation,
			m.Slides[m.CurrentPosition-1],
			m.CurrentPosition,
			len(m.Slides),
		)
	}

	rendered, _ := glamour.Render(content, m.Presentation.Config.Style)

	return rendered
}

func getFooter(m Model) string {
	status := ""
	author := ""
	paddingRight := 4

	if m.CurrentPosition > 0 && m.Presentation.Config.ShowPageStatus {
		status = m.Progress.ViewAs(float64(m.CurrentPosition) / m.TotalSlides)
		paddingLeft := (m.Width - lipgloss.Width(status)) / 2
		paddingRight = paddingLeft

		status = lipgloss.NewStyle().PaddingLeft(paddingLeft).Render(status)
	}

	if m.CurrentPosition > 0 && m.Presentation.Config.ShowAuthor {
		author = m.Meta.Author
		paddingLeft := m.Width - lipgloss.Width(author) - paddingRight

		author = lipgloss.NewStyle().PaddingLeft(paddingLeft).Render(author)
	}

	return status + "\n" + "\n" + author
}

func openAsset(m Model) {

	assets := m.Slides[m.CurrentPosition-1].Assets

	if len(assets) == 0 || m.CurrentAsset >= len(assets) {
		return
	}

	if m.Presentation.Config.ImageViewer == "" {
		return
	}

	cmd := exec.Command(m.Presentation.Config.ImageViewer, assets[m.CurrentAsset])
	_ = cmd.Start()
}
