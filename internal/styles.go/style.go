package styles

import (
	"fmt"
)


// Item indicates the field that is selected and whether or not the item is existing on the platform.
type Item struct {
	Name string
	DescriptionText string
	Existing bool
	Active bool
}

func (i Item) Title() string {
	title := i.Name
	if i.Active {
		title = fmt.Sprintf("%s %s", title, SuccessStyle.Render("(active)"))
	}
	if i.Existing {
		title = fmt.Sprintf("%s %s", title, HighlightStyle.Render("(existing)"))
	}
	return title
}

func (i Item) FilterValue() string { return i.Name }
func (i Item) Description() string { return i.DescriptionText }