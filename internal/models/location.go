package models

type ViewState int

type Location struct {
	Name string
	Desc string
	View ViewState
}

func (l Location) GetName() string {
	return l.Name
}

func (l Location) Title() string       { return l.Name }
func (l Location) Description() string { return l.Desc }
func (l Location) FilterValue() string { return l.Name }
