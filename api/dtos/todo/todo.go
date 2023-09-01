package tododtos

import (
	"time"
)

type Todo struct {
	ID int64				`json:"id"`
	Title string			`json:"title" binding:"required"`
	Description string		`json:"description"` 
	Completed bool			`json:"completed"`
	Date time.Time			`json:"date"`
	UpdatedDate time.Time	`json:"updatedDate"`
}

func (td *Todo) MapChanges(newTodo UpdateTodo) {

	if newTodo.Title != nil && td.Title != *newTodo.Title {
		td.Title = *newTodo.Title
	}

	if newTodo.Description != nil && td.Description != *newTodo.Description {
		td.Description = *newTodo.Description
	}

	if newTodo.Completed != nil && td.Completed != *newTodo.Completed {
		td.Completed = *newTodo.Completed
	}

	if newTodo.Date != nil && td.Date != *newTodo.Date {
		td.Date = *newTodo.Date
	}
}
