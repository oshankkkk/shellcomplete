package main

import (
	"bufio"
	"fmt"
	"os"
)

//input filepath is given
func inputReader(filepath string){

fmt.Println("filepath",filepath)
// this loads the whole file into memory
//file,_:=os.ReadFile(filepath)
//fmt.Println(string(file))
file,_:=os.Open(filepath)
defer file.Close()
scanner:=bufio.NewScanner(file)
for scanner.Scan(){
	fmt.Println(scanner.Text())
}

}
func lexer(){}
func parser(){}
