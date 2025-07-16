package data

import (
	"task_manager/models"
	"time"
)

var tasks = []models.Task{
	{ID: 1, Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "In Progress"},
    {ID: 1, Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: 3, Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

var nextTask int = tasks[len(tasks) - 1].ID

/* data layer / Business logics and rules */

func GetAllTasks() []models.Task{
	return tasks
}

func GetTaskByID(id int) (models.Task, bool){ 
	// 2nd return value is for, found or !found
	for _, task := range tasks {
		if task.ID == id {
			return task, true
		}
	}
	return models.Task{}, false
}

func AddTask(newTask models.Task) {
	nextTask++
	newTask.ID = nextTask

	tasks = append(tasks, newTask)
}

func UpdateTask(id int, updatedTask models.Task) bool {
	
	for i, task := range tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if !updatedTask.DueDate.IsZero() {
				tasks[i].DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}
			return true
		}
	}
	return false
}

func DeleteTask(id int) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}