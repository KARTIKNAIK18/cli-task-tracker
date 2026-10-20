package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/KARTIKNAIK18/cli-task-tracker/task"
)


func main(){

	var taskList []task.Task

	if len(os.Args)<2{
		fmt.Println("U need to give the argumants Along with the Go run .")
		fmt.Println("Commands: add, list, update, delete, done, list-done")
		return
	}

	switch os.Args[1]{
	case "add" :
		if  len(os.Args)<3{
			fmt.Println(`Usage: go run . add "Task description"`)
			return
		}
		task.AddTask(&taskList,os.Args[2])

	case  "list":
		task.ListTask(&taskList)

	case "update":
        if len(os.Args) < 4 {
            fmt.Println(`Usage: go run . update <id> "New description"`)
            return
        }

        id, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println("Invalid task ID:", os.Args[2])
            return
        }
        task.UpdateTask(&taskList, id, os.Args[3])

	case "delete":
        if len(os.Args) < 3 {
            fmt.Println("Usage: go run . delete <id>")
            return
        }

        id, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println("Invalid task ID:", os.Args[2])
            return
        }
        task.Deletetask(&taskList, id)


	 case "done":
        if len(os.Args) < 3 {
            fmt.Println("Usage: go run . done <id>")
            return
        }

        id, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println("Invalid task ID:", os.Args[2])
            return
        }
        task.UpdateStatus(&taskList, id)
	
	 case "list-done":
        task.ListCompleted(&taskList)

	 default:
		fmt.Println("Unknown command:", os.Args[1])

	}

}






























