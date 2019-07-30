package common

type Adapter interface {
	Initialize(Configuration)            // Initialize the adapter
	List(string) []*Task                 // Return the adapter Task list to print the scrum
	NextScrum()                          // Hook that is called when a scrum is done
	Add(description string, list string) // Add a new task in the list
}
