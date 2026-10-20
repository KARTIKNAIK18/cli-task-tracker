package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)




func CreateFile(){
	file,err:=os.Create("tasks.json")
	if err!=nil{
		fmt.Println("Error :",err)
		return
	}

	defer file.Close()
}





func AddTask(temp *[]Task,task string){

	err := LoadData(temp)
    if err != nil && !os.IsNotExist(err) {
        fmt.Println("Error loading tasks:", err)
        return
    }

		var Dtask Task
		Id:=1
		for _, t := range *temp {
   			 if t.ID >= Id {
        	Id = t.ID + 1
    		}
		}

			Dtask=Task{
				ID: Id,
				Description: task,
				Status: false,
				CreatedAt: time.Now(),
				UpdatedAt: nil,


			}
		
		*temp=append(*temp, Dtask)

		data,err:=json.MarshalIndent(*temp,"","  ")

		if err!=nil {
			fmt.Println("Error :",err)
			return
		}



		err1:=os.WriteFile("tasks.json",[]byte(data),0644)

		if err1!=nil {
			fmt.Println("error in Adding task :",err)
			return
		}

		fmt.Println("Added task Succesfully")
	}





func updateStatus(task *[]Task,taskId int){
	LoadData(task)

	if len(*task)==0 {
		fmt.Println("No tasks found!")
		return
	}

	Found:=false

	for i, v := range *task {
		if v.ID==taskId {
			(*task)[i].Status=true
			Found=true
			break
		}
	}

	if !Found {
		fmt.Println("Error: Task ID not found")
		return
	}
	SaveData(task)
	fmt.Println("Task update as Complted successfully!")
}








//-----------------------------
func Deletetask(task *[]Task,taskId int){
	
	LoadData(task)

	if len(*task)==0 {
		fmt.Println("No tasks found!")
		return
	}

	Found:=false

	for i, v := range *task {
		if v.ID==taskId {
			*task=append((*task)[:i],(*task)[i+1:]...)
			Found=true
			break
		}
	}
	if Found {
		id:=1
    for i := 0; i < len(*task); i++ {
        (*task)[i].ID = id
		id++
    }
}
		

	if !Found {
		fmt.Println("Error: Task ID not found")
		return
	}


	
	SaveData(task)
	fmt.Println("Task deleted successfully!")
}





func UpdateTask(task *[]Task,TaskId int,Update string){

	if e:=LoadData(task); e!=nil{
			  fmt.Println("Error loading tasks:", e)
    		  return
	}
	found := false
	for i := range *task{
		if (*task)[i].ID == TaskId {
			(*task)[i].Description=Update
			Now:=time.Now()
			(*task)[i].UpdatedAt=&Now
			found=true
			break
		}
	}

	if !found {
		errors.New("task Not Found")
	}

	if SaveData(task)!=nil{
		errors.New("Data Not Saved!!")
	}
	fmt.Println("data Updated Sucessfully")
}
