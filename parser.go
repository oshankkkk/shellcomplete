package main

import (
	"fmt"
	"os"
)

//input filepath is given
func inputReader(filepath string){
fmt.Println("filepath",filepath)
file,_:=os.ReadFile(filepath)
fmt.Println(string(file))

}
func lexer(){}
func parser(){}
