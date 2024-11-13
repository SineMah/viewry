package data

type General struct {
	Meta         Meta         `yaml:",inline"`
	Contact      Contact      `yaml:"contact"`
	Presentation Presentation `yaml:"presentation"`
	Slides       []Slide      `yaml:"slides"`
}
