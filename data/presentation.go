package data

type Config struct {
	ShowTitle      bool   `yaml:"show_title"`
	ShowPage       bool   `yaml:"show_page"`
	ShowAuthor     bool   `yaml:"show_author"`
	ShowAuthorMeta bool   `yaml:"show_author_meta"`
	ShowPageStatus bool   `yaml:"show_page_status"`
	Theme          string `yaml:"theme"`
	ImageViewer    string `yaml:"asset_viewer"`
	Style          string `yaml:"style"`
	ProgressColor1 string `yaml:"page_status_color1"`
	ProgressColor2 string `yaml:"page_status_color2"`
}

type Presentation struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Config      Config   `yaml:"config"`
	Tags        []string `yaml:"tags"`
}
