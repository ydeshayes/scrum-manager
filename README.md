IN PROGRESS

# Scrum Manager

A Golang cli program that will generate your scrum report for you.
The database is a Trello board for now, but I want to implement more adapters.

## Usage
### Todos:
    ./main add "Add this task to the todo"
You can select todos and move them into the today and next part of the scrum (See `move` command)
### Today:
    ./main add-today "Add this task to the today part of the scrum"
### Next:
    ./main add-next "Add this task to the next part of the scrum"
### Blockers:
    ./main add-blocker "Add this task to the blockers part of the scrum"
### Move:
    ./main move [listName1 listName2] (see possible lists bellow)
- Default is todos to today
### Preview your scrum report (it will copy the scrum report into the clipboard)
    ./main scrum-preview
### Generate your scrum report (it will copy the scrum report into the clipboard)
    ./main scrum
- When you generate the scrum, it will archived the today tasks and move the next tasks into today for the next scrum. If you don't have blockers yet, it will ask you if you want to add one.
- If you activate the google calendar integregation, it will copy all the event that you had since the last scrum into the today part of the new scrum.

## Lists
The lists that you can use are:
- todos
- today
- next
- blockers
## Config
For now the config file need to be in `~/.scrum-manager`
```
{
	"AdapterName": "Trello",
	"TrelloApikey": "",
	"TrelloToken": "",
	"TrelloBoardId": "",

	"ScrumWelcomeText": "Hello! this is my scrum",
    "GoogleAppCredentialPath": "[Optional] Absolute path to the google app credential json file"
}
```
## Google calendar integration
- If you activate the google calendar integregation, it will copy all the event that you had since the last srum into the today part of the new scrum.
- Please visit [quickstart](https://developers.google.com/calendar/quickstart/go) to generate your json crendential file
## Build
```
go build main.go
```