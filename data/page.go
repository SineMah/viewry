package data

type Page struct {
	Title          string
	Content        string
	ShowPageStatus bool
	ShowAuthor     bool
	Author         string
	CurrentPage    int
	TotalPages     int
}
