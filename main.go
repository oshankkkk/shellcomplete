package main

import ( 

	tea "github.com/charmbracelet/bubbletea"
	"fmt"
	"os"
	)

func main() {

	//fmt.Println("Got:", os.Args[1])
	var bashhistory []string
	usercmd:=os.Args[1]
	historyfile,_:=getbashhistory()	
	bashhistory=readbashhistory(historyfile)
	sugessionlist:=sortbashhistory(usercmd,bashhistory)
	fmt.Println(sugessionlist)
	p := tea.NewProgram(initialModel());
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)	
	}
}
