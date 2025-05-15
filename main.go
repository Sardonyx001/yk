package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Status int

const (
	Ongoing Status = iota
	Done
)

type Task struct {
	status      Status
	title, desc string
}

func (t Task) Title() string       { return t.title }
func (t Task) Description() string { return t.desc }
func (t Task) FilterValue() string { return t.title }

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
		Task{status: Ongoing, title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		Task{status: Ongoing, title: "Nutella", desc: "It's good on toast"},
		Task{status: Ongoing, title: "Bitter melon", desc: "It cools you down"},
		Task{status: Ongoing, title: "Nice socks", desc: "And by that I mean socks without holes"},
		Task{status: Ongoing, title: "Eight hours of sleep", desc: "I had this once"},
		Task{status: Ongoing, title: "Cats", desc: "Usually"},
		Task{status: Ongoing, title: "Plantasia, the album", desc: "My plants love it too"},
		Task{status: Ongoing, title: "Pour over coffee", desc: "It takes forever to make though"},
		Task{status: Ongoing, title: "VR", desc: "Virtual reality...what is there to say?"},
		Task{status: Ongoing, title: "Noguchi Lamps", desc: "Such pleasing organic forms"},
		Task{status: Ongoing, title: "Linux", desc: "Pretty much the best OS"},
		Task{status: Ongoing, title: "Business school", desc: "Just kidding"},
		Task{status: Ongoing, title: "Pottery", desc: "Wet clay is a great feeling"},
		Task{status: Ongoing, title: "Shampoo", desc: "Nothing like clean hair"},
		Task{status: Ongoing, title: "Table tennis", desc: "It’s surprisingly exhausting"},
		Task{status: Ongoing, title: "Milk crates", desc: "Great for packing in your extra stuff"},
		Task{status: Ongoing, title: "Afternoon tea", desc: "Especially the tea sandwich part"},
		Task{status: Ongoing, title: "Stickers", desc: "The thicker the vinyl the better"},
		Task{status: Ongoing, title: "20° Weather", desc: "Celsius, not Fahrenheit"},
		Task{status: Ongoing, title: "Warm light", desc: "Like around 2700 Kelvin"},
		Task{status: Ongoing, title: "The vernal equinox", desc: "The autumnal equinox is pretty good too"},
		Task{status: Ongoing, title: "Gaffer’s tape", desc: "Basically sticky fabric"},
		Task{status: Ongoing, title: "Terrycloth", desc: "In other words, towel fabric"},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "やること"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
