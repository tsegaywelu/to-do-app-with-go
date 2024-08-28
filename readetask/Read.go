package read

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var show = fmt.Println

type Task struct {
	Name                 string
	Id                   string
	Creation_Date        time.Time
	Task_start_Time      time.Time
}

func Read(channelread chan string  ) {
	
	file,err:=os.Open("mytask.txt")

	if err !=nil{
		show("you have not register any task so please chose 1 to add task ")
		return
	}
	
	defer file.Close()
	scanner:=bufio.NewScanner(file)

for scanner.Scan(){
var task Task
err := json.Unmarshal([ ]byte(  scanner.Text()),   &task)
if err!=nil {
	log.Fatal(err)
}
show("task name :", task.Name)
show("task  ID:", task.Id)
		
		show("task creation  Date:", task.Creation_Date)
		show("task assigning  Time:", task.Task_start_Time)
		show(strings.Repeat("*", 100))  //this is used to separete tasks 

}


//channelread<-"i have finished "  // since i have finished running this code i have to send for main. because he is waiting for this code 

channelread<-"i have finished reading tasks  "


}
	

