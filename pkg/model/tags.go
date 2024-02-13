//go:generate easyjson -all tags.go
package model

type TagDTO struct {
	// ID is string.
	//
	// It could be any string, but mostly is UUID v4 string.
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Slug  string `json:"slug"`
}
