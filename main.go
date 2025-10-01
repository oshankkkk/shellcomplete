package main

import ( 

	tea "github.com/charmbracelet/bubbletea"
	"fmt")

func main(){
	fmt.Println("hi")
	app:=tea.NewProgram(& popup{})
	if _,err:=app.Run();err!=nil{
		fmt.Println("kela una", err)
	}


}

