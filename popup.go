package main

import (

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type popup struct{
	userinput textinput.Model
	showsuggesionbox bool	
	suggestions []int


} 
func InitialModel() popup{
	userinput:=textinput.New()
	userinput.Placeholder ="ye:"
	userinput.Focus()
	return popup{
		userinput: userinput , 
		showsuggesionbox: false,
	}


}
func (p *popup) Init()(tea.Cmd){
	return textinput.Blink
	
}
func (p *popup) Update(msg tea.Msg)(tea.Model , tea.Cmd){
	switch msg:= msg.(type){
	case tea.KeyMsg:
		switch msg.String(){
			case "tab":
				p.showsuggesionbox=true
			case "esc","ctrl+c":
				return p, tea.Quit

		}

}
		var cmd tea.Cmd
		p.userinput, cmd= p.userinput.Update(msg)
		return p,cmd


}

func (p *popup) View()string{
      output := "> " + p.userinput.View()

	if p.showsuggesionbox{
		popupContent := "Option 1\nOption 2\nOption 3"
		output += "\n" + popupContent
	}

	return output

}
