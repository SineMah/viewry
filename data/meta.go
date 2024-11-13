package data

type Meta struct {
	Author string `yaml:"author"`
	Date   string `yaml:"date"`
}

type Contact struct {
	Mail  string `yaml:"mail"`
	Phone string `yaml:"phone"`
}
