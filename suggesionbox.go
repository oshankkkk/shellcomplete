package main

import (
	"strings"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type suggestionbox struct {
	showsuggestions bool
	suggestions     []string
}
var suggestionBoxStyle = lipgloss.NewStyle().
    Background(lipgloss.Color("#303446")). 
    Foreground(lipgloss.Color("#f2d5cf")).     
    Border(lipgloss.RoundedBorder()).
Padding(1,2 ) 

func initialModel() suggestionbox{
	// get the user input data
	// match and sort it
	// put it into the suggestions
	// view it

	return suggestionbox{
		showsuggestions: true,
		suggestions:     []string{"ls", "cd", "git status", "go run"},
	}
}

func (s suggestionbox) Init() tea.Cmd {
	return nil
}

func (s suggestionbox) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg:=msg.(type){
	case tea.KeyMsg:
		switch msg.String(){
		case "q":
			return s,tea.Quit

		}




	}
	return s,nil
}

func (s suggestionbox) View() string {
	var builder strings.Builder
	if s.showsuggestions {
		for _, s := range s.suggestions {
			builder.WriteString( " " + s + "\n")

		}
	}


	list:=builder.String()
	return suggestionBoxStyle.Render(list)	

}


