package main

import ( 

	tea "github.com/charmbracelet/bubbletea"
	"fmt"
	"os"
	)

func main() {
	p := tea.NewProgram(initialModel());
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)	
}
}
