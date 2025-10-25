package main

import (
	//"fmt"
	"fmt"
	"os"
)
func main(){
	userhome:=os.Getenv("HOME")	
	fmt.Println(userhome)
	filepath:=userhome+"/.bash_history"
	inputReader(filepath)
}
