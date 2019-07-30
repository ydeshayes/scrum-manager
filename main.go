package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	common "github.com/ydeshayes/dev-diary/common"
	trelloAdapter "github.com/ydeshayes/dev-diary/trelloAdapter"
)

func main() {
	config := common.GetConfiguration()
	adapter := selectAdapter(config)
	adapter.Initialize(config)

	switch os.Args[1] {
	// Add task to todo
	case "add":
		text := os.Args[2]
		fmt.Println("text: " + text)
		adapter.Add(text, "todos")
		break
	case "add-today":
		text := os.Args[2]
		adapter.Add(text, "today")
	case "add-next":
		text := os.Args[2]
		adapter.Add(text, "next")
	case "scrum":
		todayList := adapter.List("today")
		nextList := adapter.List("next")
		scrum := generateScrum(todayList, nextList, config.ScrumWelcomeText)
		fmt.Print(scrum)
		clipboard.WriteAll(scrum)
		// TODO:
		// Ask something to add for today?
		// Ask something to add next?
		// Ask any bloker?
		adapter.NextScrum()
		break
	case "scrum-preview":
		todayList := adapter.List("today")
		nextList := adapter.List("next")
		scrum := generateScrum(todayList, nextList, config.ScrumWelcomeText)
		fmt.Print(scrum)
		clipboard.WriteAll(scrum)
		break
	case "next":
		break
	case "done":
		break
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

func selectAdapter(config common.Configuration) common.Adapter {
	var adapter common.Adapter
	if config.AdapterName == "Trello" {
		adapter = &trelloAdapter.TrelloAdapter{}
	}
	return adapter
}

func generateScrum(todayTasks []*common.Task, nextTasks []*common.Task, title string) string {
	scrumText := title + "\n" + "Today:\n"
	for _, task := range todayTasks {
		scrumText += " - " + task.Title + "\n"
	}
	scrumText += "Next:\n"
	for _, task := range nextTasks {
		scrumText += " - " + task.Title + "\n"
	}

	return scrumText
}
