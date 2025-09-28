package main

import (
	"bufio"
	"fmt"
	"os/user"
	"os"
	"sort"
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

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return historylist

}
type sortedcmds struct {
	cmd string
	distance int
}
func sortbashhistory(usercmd string, bashhistory []string)[]string{
	var sortedbashhistory []sortedcmds

	for _, cmds := range bashhistory {
		sortedbashhistory = append(sortedbashhistory,sortedcmds{
			cmd:      cmds,
			distance: editDistance(cmds, usercmd),
		})
	}
	sort.Slice(sortedbashhistory, func(i, j int) bool {
		return sortedbashhistory[i].distance < sortedbashhistory[j].distance
	})

	result := []string{}
	for _, sc := range sortedbashhistory {
		result = append(result, sc.cmd)
	}

	return result
}


