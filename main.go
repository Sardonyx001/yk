package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type status int

const (
	Ongoing status = iota
	Done
)

type item struct {
	status      status
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
	items := []list.Item{
		item{status: Ongoing, title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		item{status: Ongoing, title: "Nutella", desc: "It's good on toast"},
		item{status: Ongoing, title: "Bitter melon", desc: "It cools you down"},
		item{status: Ongoing, title: "Nice socks", desc: "And by that I mean socks without holes"},
		item{status: Ongoing, title: "Eight hours of sleep", desc: "I had this once"},
		item{status: Ongoing, title: "Cats", desc: "Usually"},
		item{status: Ongoing, title: "Plantasia, the album", desc: "My plants love it too"},
		item{status: Ongoing, title: "Pour over coffee", desc: "It takes forever to make though"},
		item{status: Ongoing, title: "VR", desc: "Virtual reality...what is there to say?"},
		item{status: Ongoing, title: "Noguchi Lamps", desc: "Such pleasing organic forms"},
		item{status: Ongoing, title: "Linux", desc: "Pretty much the best OS"},
		item{status: Ongoing, title: "Business school", desc: "Just kidding"},
		item{status: Ongoing, title: "Pottery", desc: "Wet clay is a great feeling"},
		item{status: Ongoing, title: "Shampoo", desc: "Nothing like clean hair"},
		item{status: Ongoing, title: "Table tennis", desc: "It’s surprisingly exhausting"},
		item{status: Ongoing, title: "Milk crates", desc: "Great for packing in your extra stuff"},
		item{status: Ongoing, title: "Afternoon tea", desc: "Especially the tea sandwich part"},
		item{status: Ongoing, title: "Stickers", desc: "The thicker the vinyl the better"},
		item{status: Ongoing, title: "20° Weather", desc: "Celsius, not Fahrenheit"},
		item{status: Ongoing, title: "Warm light", desc: "Like around 2700 Kelvin"},
		item{status: Ongoing, title: "The vernal equinox", desc: "The autumnal equinox is pretty good too"},
		item{status: Ongoing, title: "Gaffer’s tape", desc: "Basically sticky fabric"},
		item{status: Ongoing, title: "Terrycloth", desc: "In other words, towel fabric"},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "My Fave Things"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
