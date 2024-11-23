package style

type Style struct {
	Styles       []string
	DefaultStyle string
}

var defaultStyle = "dracula"

func New() *Style {
	s := Style{
		DefaultStyle: defaultStyle,
	}

	s.AddStyle(defaultStyle)
	s.AddStyle("ascii")
	s.AddStyle("dark")
	s.AddStyle("light")
	s.AddStyle("notty")
	s.AddStyle("pink")
	s.AddStyle("tokyo-night")

	return &s
}

func (s *Style) AddStyle(style string) {
	s.Styles = append(s.Styles, style)
}

func (s *Style) GetStyle(style string) string {

	for _, t := range s.Styles {

		if t == style {
			return t
		}
	}

	return s.DefaultStyle
}
