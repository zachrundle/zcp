package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Cluster struct {
	base *BaseCluster
	adv  *AdvancedSettings
}

func (c Cluster) Init() tea.Cmd {
	return tea.SetWindowTitle("zcp")
}

func InitialModel() Cluster {
	return Cluster{
		base: &BaseCluster{
			platform:  []string{"aws", "vsphere", "baremetal"},
			selected: make(map[int]struct{}),
			// You can initialize other fields here too if needed.
		},
		adv: &AdvancedSettings{
			hyperThreading: false, // or whatever default makes sense
		},
	}
}

func (c Cluster) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return c, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if c.base.cursor > 0 {
				c.base.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if c.base.cursor < len(c.base.platform)-1 {
				c.base.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := c.base.selected[c.base.cursor]
			if ok {
				delete(c.base.selected, c.base.cursor)
			} else {
				c.base.selected[c.base.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return c, nil
}

func (c Cluster) View() string {
	// The header
	s := "What platform are you deploying to?\n\n"

	// Iterate over our choices
	for i, choice := range c.base.platform {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if c.base.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := c.base.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
