package common

import "time"

type Adapter interface {
	Initialize(Configuration)            // Initialize the adapter
	List(string) []*Task                 // Return the adapter Task list to print the scrum
	Move(Task, string) error             // Move the task into the list that have the name ...
	NextScrum()                          // Hook that is called when a scrum is done
	Add(description string, list string) // Add a new task in the list
	LastScrumDate() time.Time            // Get the last scrum date
}
