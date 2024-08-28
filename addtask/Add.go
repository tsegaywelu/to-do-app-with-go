package adder

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"
)

var show=fmt.Println
//lets create task object here and i will stor as json file
type Task struct {
	Name                 string
	Id                   string
	Creation_Date        time.Time
	Task_start_Time      time.Time
}
//custome error for empty task i will call this function 
type EmpthyTask struct{}
func (e * EmpthyTask) Error() string{
	return"you can not register an empty task. plase fill out all the fields!"

}
func Add(channeladder chan int ) {

	var task Task

	task.Creation_Date=time.Now()

	//here i have to ceate Id  using random numbers
	random,err:=rand.Int(rand.Reader,big.NewInt(700000000))
	if err!=nil{
		show("Error when creating random number",err)
	}

	
	task.Id=time.Now().Local().Weekday().String() +random.String()

	// reciving the task name from user input 
	show(" plase enter the task name ")
	fmt.Scan(&task.Name)
	if task.Name == "" {
		show("error",&EmpthyTask{})
	}

	show("enter starting date format: 2006-01-02 16:09 ")

	//reciving the task starting time 
	//here i will use the bufio.NewReader  because there is space between time and  year 
	reader:=bufio.NewReader(os.Stdin)  
	reader.ReadString('\n')  //because the fmt.scan  is omitting the /n and i do this to solve the problem
	AssignedTime_Stringform,_:=reader.ReadString('\n')
//lets remove now the /n  and /r so to do this i can make the folllowing code 
AssignedTime_Stringform=strings.TrimSpace(AssignedTime_Stringform)
AssignedTime_Stringform=strings.ReplaceAll(AssignedTime_Stringform,"\r","")
assignedtime, err1 := time.Parse("2006-01-02 15:04", AssignedTime_Stringform)//i am checking if the dtring is in the time format 
	show(AssignedTime_Stringform)
if err1!=nil{
	show("invalid time format please use format yyyy-mm-dd hh:mm ",err1)
	return
}
// pass assignedtime  to the task.AssignedTime
task.Task_start_Time=assignedtime

//  now lets change the task data  t ojson file 
taskData,err3:=json.Marshal(task)
if err3!=nil{
	log.Fatal(err3)
}
//then lets attach the task to mytask.txt 
file,err4:=os.OpenFile("mytask.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
if err4!=nil{
	log.Fatal(err4)
}

defer file.Close()
_,err5:=file.Write(taskData)
if err5!=nil{
	log.Fatal(err5)

}
// lets write new line /n to the file mytask.txt and another tasks will be added at new line 
_,err6:=file.WriteString("\n")
if err6!=nil{
	log.Fatal(err6)
}
//i am sending this to main and main will jump to other programs 
channeladder<-45

}
