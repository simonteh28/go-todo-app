package tododtos

import "time"

// Update fields are in pointer to indicate that field is optional
type UpdateTodo struct {
	Title *string 		`json:"title,omitempty"`
	Description *string	`json:"description"`
	Completed *bool		`json:"completed"`
	Date *time.Time		`json:"date"`
}