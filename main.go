package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)
type cup struct{}

func (c cup) Init() (tea.Cmd){
return nil}

func (c cup) Update(msg tea.Msg)(tea.Model, tea.Cmd){
	if keymsg,okay:=msg.(tea.KeyMsg);okay{
		if  keymsg.String()=="q"{
			return c,tea.Quit
		}
	}
	
	return c,nil
}

func (c cup) View() string{
return "hi" }

func main(){
	fmt.Println("hi")
	app:=tea.NewProgram(cup{})
	if _,err:=app.Run();err!=nil{
		fmt.Println("kela una", err)
	}
	
}

