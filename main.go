package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"

	
)
type cup struct{
	historylist []string
	index int
}


func (c *cup) Init() (tea.Cmd){
	bashistoryfile,err:=getbashhistory();if err!=nil{
		fmt.Println(err)
	}
	c.historylist =readbashhistory(bashistoryfile)
	return nil
}

func (c *cup) Update(msg tea.Msg)(tea.Model, tea.Cmd){
	if keymsg,okay:=msg.(tea.KeyMsg);okay{
		if  keymsg.String()=="q"{
			return c,tea.Quit
		}
		c.index++

	}
	return c,nil
}

func (c *cup) View() string{
  if len(c.historylist) == 0 {
        return "No history found"
    }
    return fmt.Sprintf("%s", c.historylist[c.index])
 }



func main(){
	fmt.Println("hi")
	app:=tea.NewProgram(&cup{})
	if _,err:=app.Run();err!=nil{
		fmt.Println("kela una", err)
	}

}

