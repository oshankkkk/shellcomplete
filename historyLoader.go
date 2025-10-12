package main
import(
	"fmt"
	"os"
	"os/user"
)
func loadhistory()(*os.File,error){

	history:=os.Getenv("HISTFILE")

	if history==""{
		currentuser,err:=user.Current()
		if err!=nil{
			fmt.Println("error",err)
			return nil,err
		}
		history=currentuser.HomeDir+"/.bash_history"
	}
	bashhistory, err:=os.Open(history)
	if err!=nil{
		fmt.Println("error",err)
		return nil,err
	}

	return bashhistory,nil

}

