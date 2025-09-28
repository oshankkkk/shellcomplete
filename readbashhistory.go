package main

import (
	"bufio"
	"fmt"
	"os/user"
	"os"
)
func getbashhistory()(*os.File,error){
	currentuser,err:=user.Current()
	if err!=nil{
		fmt.Println("error",err)
		return nil,err
	}
	homedir:=currentuser.HomeDir	

	bashhistory, err:=os.Open(homedir+"/.bash_history")
	if err!=nil{
		fmt.Println("error",err)
		return nil,err
	}

	return bashhistory,nil
}
func readbashhistory(historyfile *os.File) []string{
	historylist:=[]string{}
	defer historyfile.Close() 

	scanner := bufio.NewScanner(historyfile)
	for scanner.Scan() {
		line := scanner.Text()      
		historylist=append(historylist,line)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return historylist

}
