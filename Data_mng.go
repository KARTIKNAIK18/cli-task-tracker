package main

import (
	"encoding/json"
	"fmt"
	"os"
)









func SaveData(tasks *[]Task) error {
	updatedData, err := json.MarshalIndent(*tasks, "", "  ")
	if err != nil {
		return  err
	}
	err1:= os.WriteFile("tasks.json", updatedData, 0644)

	if err1 != nil {
		return  err1
	}
	return nil
}



func LoadData(tasks *[]Task)error{

	 data, err := os.ReadFile("tasks.json")
    if err != nil {
        return err
    }

    if err := json.Unmarshal(data, tasks); err != nil {
        return err
    }

    return nil
}




func ListCompleted(task *[]Task){

	if e:=LoadData(task); e!=nil{
			  fmt.Println("Error loading tasks:", e)
    		  return
	}

	LoadData(task)

	if len(*task)==0 {
		fmt.Println("No tasks found!")
		return
	}

	for _, v := range *task {
			statusCheack :="Not Completed"
			Updated:="Not modified"
			if v.Status {
				if v.UpdatedAt!=nil {
					Updated=v.UpdatedAt.Format("2006-01-02 15:04")
				}
				statusCheack= "Complted"
				 fmt.Printf("%-7d %-60s %-15s %-25s %-25s\n",
        v.ID,
        v.Description,
        statusCheack,
        v.CreatedAt.Format("2006-01-02 15:04"),
        Updated,
    )
			}
	}
}





func ListTask(task *[]Task){
	if e:=LoadData(task); e!=nil{
			  fmt.Println("Error loading tasks:", e)
    		  return
	}

	LoadData(task)

	if len(*task)==0 {
		fmt.Println("No tasks found!")
		return
	}

	fmt.Println("==================================================================================")
	fmt.Printf("%-7s %-60s %-15s %-15v %-15v\n", "ID", "TASK LIST", "STATUS", "CREATED_AT", "UPDATED_AT")
	fmt.Println("==================================================================================")
	
	for _, v := range *task {
		statusCheack :="Not Completed"
		Updated:="Not modified"

		if v.UpdatedAt!=nil {
			Updated=v.UpdatedAt.Format("2006-01-02 15:04")
		}
		
		if v.Status {
			statusCheack= "Complted"
		}

			 fmt.Printf("%-7d %-60s %-15s %-25s %-25s\n",
        v.ID,
        v.Description,
        statusCheack,
        v.CreatedAt.Format("2006-01-02 15:04"),
        Updated,
    )
	}

}