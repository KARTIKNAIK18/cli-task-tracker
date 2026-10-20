package main

import (
    "time"
)

type Task struct{
	ID int
	Description string
	Status bool
    CreatedAt time.Time 
    UpdatedAt *time.Time
}
