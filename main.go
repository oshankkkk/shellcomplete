package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"

	
)
type shellcomplete struct{
	historylist []string
	index int
	sortedhistorylist []string
}



func (c *shellcomplete) Init() (tea.Cmd){
	bashistoryfile,err:=getbashhistory();if err!=nil{
		fmt.Println(err)
	}
	c.historylist =readbashhistory(bashistoryfile)
	return nil
}

func (c *shellcomplete) Update(msg tea.Msg)(tea.Model, tea.Cmd){
	if keymsg,okay:=msg.(tea.KeyMsg);okay{
		if  keymsg.String()=="q"{
			return c,tea.Quit
		}
		

	}
	return c,nil
}

func (c *shellcomplete) View() string{
  if len(c.historylist) == 0 {
        return "No history found"
    }
    return fmt.Sprintf("%s", c.historylist[c.index])
 }



func main(){
	fmt.Println("hi")
	app:=tea.NewProgram(&shellcomplete{})
	if _,err:=app.Run();err!=nil{
		fmt.Println("kela una", err)
	}


}

