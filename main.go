package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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
	case "move":
		var listName string
		var listNameDestination string
		if len(os.Args) < 4 {
			listName = "todos"
			listNameDestination = "today"
			fmt.Println("Defaulted to todos -> today")
		} else {
			listName = os.Args[2]
			listNameDestination = os.Args[3]
		}

		todosList := adapter.List(listName)

		if len(todosList) == 0 {
			fmt.Println("Empty " + listName + " :)")
			break
		}

		fmt.Println("List of todos:")
		for index, task := range todosList {
			fmt.Println(strconv.Itoa(index) + " - " + task.Title)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please enter the index(ex) comma separated: ")
		text, _ := reader.ReadString('\n')

		indexes := strings.Split(text, ",")
		for _, indexString := range indexes {
			// Convert indexString to int
			index, err := strconv.Atoi(strings.Trim(indexString, "\n "))
			if err != nil {
				fmt.Println(err)
				fmt.Println("Error on index (invalid) " + indexString)
				continue
			}

			adapter.Move(*todosList[index], listNameDestination)
		}
		break
	case "next":
		break
	case "done":
		break
	default:
		fmt.Println("expected 'add' or 'move' or 'scrum-preview' or 'scrum' subcommands")
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
