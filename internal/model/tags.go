//go:generate easyjson -all tags.go
package model

type TagDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Slug  string `json:"slug"`
}
