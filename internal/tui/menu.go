package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zachrundle/zcp/internal/model"
	"github.com/zachrundle/zcp/internal/setup"
	"github.com/zachrundle/zcp/internal/ui"
)

type Cluster struct {
	base *model.BaseCluster
	adv  *model.AdvancedSettings
	tui  *setup.TUIConfig
}

func (c Cluster) Init() tea.Cmd {
	return tea.SetWindowTitle("zcp")
}

func InitialModel(cfg setup.TUIConfig) Cluster {
	return Cluster{
		base: &model.BaseCluster{
			Platform: []string{"aws", "vsphere", "baremetal"},
			Selected: make(map[int]struct{}),
			// You can initialize other fields here too if needed.
		},
		adv: &model.AdvancedSettings{
			HyperThreading: false, // or whatever default makes sense
		},
		tui: &cfg,
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
			if c.base.Cursor > 0 {
				c.base.Cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if c.base.Cursor < len(c.base.Platform)-1 {
				c.base.Cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := c.base.Selected[c.base.Cursor]
			if ok {
				delete(c.base.Selected, c.base.Cursor)
			} else {
				c.base.Selected[c.base.Cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return c, nil
}

func (c Cluster) View() string {
	ts := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF69B4")).
		MarginBottom(1).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lipgloss.Color("#00FFFF")).
		PaddingBottom(1)

	b := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#00FFFF")).
		Padding(1, 2).
		Width(min(c.tui.Width-4, 80))

	h := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FFFF")).
		Bold(true)

	f := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FFFF")).
		MarginTop(1)

	t := ts.Render("menu")

	l := strings.Join(ui.TextLogo, "\n")
	ls := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF69B4")).
		Render(l)

	var welcomeScreen string
	awsSSOLogin := h.Render("'aws sso login --profile my-profile'")
	awsConfigure := h.Render("'aws configure'")

	welcomeScreen = fmt.Sprintf(`To use zcp, you need to ensure that you have logged into aws cli.

If you are using AWS SSO, execute the following:
%s

If you are using API keys, execute the following:
%s`,
		awsSSOLogin, awsConfigure)

	bx := b.Render(welcomeScreen)
	ft := f.Render(c.tui.KeyPrompt)

	pt := max(0, (c.tui.Height-lipgloss.Height(t)-lipgloss.Height(bx)-lipgloss.Height(ft)-4)/2)
	pts := strings.Repeat("\n", pt)

	return  ls + pts + lipgloss.JoinVertical(lipgloss.Center,
		t,
		bx,
		ft,
	)
}
