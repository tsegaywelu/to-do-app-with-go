package main

import (
	adder "TASKSAPP/addtask"
	delete "TASKSAPP/deletetask"
	read "TASKSAPP/readetask"
	update "TASKSAPP/updatetask"
	"fmt"
	"time"
)

var show = fmt.Println

func main() {
// i want to measure the time that takes my code to run 
start:=time.Now();


  channelread:=make(chan string )
	go read.Read(channelread)
	messae_from_read:=<-channelread
	show(messae_from_read)
	var choose_operation int

	show("         1 add task")
	show("         2 delete task")
	show("         3 update task")
	show("         4 to exit the program")
	fmt.Scan(&choose_operation)

switch choose_operation{
case 1:
	channeladder:=make(chan int )
	go adder.Add(channeladder)
	message:=<-channeladder
	show(message)
	
	case 2:
		delete.Delete()
	case 3:
		update.Update()
	case 4:
		return
}
 duration:=time.Since(start)

 show("time taken to run is:",duration)
main() //lets make our code recursive 
}