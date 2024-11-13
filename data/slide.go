package data

type Slide struct {
	Type    string   `yaml:"type"`
	Title   string   `yaml:"title"`
	Content string   `yaml:"content"`
	Assets  []string `yaml:"assets"`
}

type Slides struct {
	Items []Slide `yaml:"slides"`
}
