package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	showsuggestions bool
	suggestions     []string
}

func initialModel() model {
	return model{
		showsuggestions: true,
		suggestions:     []string{"ls", "cd", "git status", "go run"},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg:=msg.(type){
	case tea.KeyMsg:
		switch msg.String(){
		case "q":
			return m,tea.Quit

		}




	}
	return m,nil
}

func (m model) View() string {
	var builder strings.Builder
	if m.showsuggestions {
		for _, s := range m.suggestions {
			builder.WriteString( " " + s + "\n")
		}
	}
	return builder.String()

}


