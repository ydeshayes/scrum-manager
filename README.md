IN PROGRESS

# Scrum Manager

A Golang cli program that will generate your scrum report for you.
The database is a Trello board for now, but I want to implement more adapters.


## Usage
### Todos:
    ./main add "Add this task to the todo"
In the future it will be possible to select todos and move them into the today and next part of the scrum
### Today:
    ./main add-today "Add this task to the today part of the scrum"
### Next:
    ./main add-next "Add this task to the next part of the scrum"
### Move:
    ./main move [listName1 listName2] (see possible lists bellow)
- Default is todos to today
### Preview your scrum report (it will copy the scrum report into the clipboard)
    ./main scrum-preview
### Generate your scrum report (it will copy the scrum report into the clipboard)
    ./main scrum
- When you generate the scrum, it will archived the today tasks and move the next tasks into today for the next scrum.

## Lists
The lists that you can use are:
- todos
- today
- next
## Config
For now the config file need to be in `~/.scrum-manager`
```
{
	"AdapterName": "Trello",
	"TrelloApikey": "",
	"TrelloToken": "",
	"TrelloBoardId": "",

	"ScrumWelcomeText": "Hello! this is my scrum"
}
```