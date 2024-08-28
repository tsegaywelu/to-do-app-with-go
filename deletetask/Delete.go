package delete

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)
var show=fmt.Println
type Task struct {
	Name            string
	Id              string
	Creation_Date   time.Time
	Task_start_Time time.Time
}

func Delete() {
var  mytask_iD  string

show("please enter the task you want to delete ")
fmt.Scan(&mytask_iD)
// now i am going to read all the tasks from the  file mytask.txt 

file,err:=os.Open("mytask.txt")

if err!=nil{

	show ("there is no file with name mytask.txt ",err)

}
defer file.Close()

var tasks []Task //creating slice with type  task 
scanner:=bufio.NewScanner(file)
for scanner.Scan(){
	var task Task
	err:=json.Unmarshal(scanner.Bytes(),&task)
	if err!=nil{
		show("error when ",err)
	}
tasks=append(tasks, task)
//finding the task with the given id 

var updatedtasks []Task
for _ ,task:=range tasks{
	if strings.TrimSpace(task.Id)!=mytask_iD{ //here the task that we enter it Id will not append to updated tasks 

		updatedtasks=append(updatedtasks, task)
		

	}
}

//writing the updated task back to the file 

file,err1:=os.Create("mytask.txt")

if err1!=nil{
	show("error when creating the file")
}
defer file.Close()

writing:=bufio.NewWriter(file)
for _,insert_tasks:=range updatedtasks{
	taskdata,err:=json.Marshal(insert_tasks)
	if err!=nil{
		show("error ",err)
	}
  _,err4:=writing.Write(taskdata)

  if err4!=nil{
	show("can not write the updated task ",err4)
  }

  _,err5:=writing.WriteString("\n")
  if err5!=nil{
	show("can not write the updated task ",err5)
  }


}


}








}