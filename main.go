package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	common "github.com/ydeshayes/dev-diary/common"
	googleCalendar "github.com/ydeshayes/dev-diary/googleCalendar"
	trelloAdapter "github.com/ydeshayes/dev-diary/trelloAdapter"
)

func main() {
	config := common.GetConfiguration()
	adapter := selectAdapter(config)
	adapter.Initialize(config)

	if len(os.Args) == 1 {
		fmt.Println("Please read the README")
	}

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
	case "add-blocker":
		text := os.Args[2]
		adapter.Add(text, "blockers")
	case "scrum":
		lastScrumDate := adapter.LastScrumDate()
		if config.GoogleAppCredentialPath != "" {
			googleCalendar.Connect(config.GoogleAppCredentialPath, lastScrumDate, adapter, false)
		}

		blockers := adapter.List("blockers")
		if len(blockers) == 0 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Any bloker? (leave empty if not):")
			blockerText, _ := reader.ReadString('\n')
			if blockerText != "No\n" && blockerText != "no\n" && blockerText != "\n" {
				adapter.Add(blockerText, "blockers")
			}
		}

		todayList := adapter.List("today")
		nextList := adapter.List("next")
		blockers = adapter.List("blockers")
		scrum := generateScrum(todayList, nextList, blockers, config.ScrumWelcomeText)
		fmt.Print(scrum)
		clipboard.WriteAll(scrum)
		adapter.NextScrum()
		break
	case "scrum-preview":
		lastScrumDate := adapter.LastScrumDate()

		if config.GoogleAppCredentialPath != "" {
			googleCalendar.Connect(config.GoogleAppCredentialPath, lastScrumDate, adapter, true)
		}

		todayList := adapter.List("today")
		nextList := adapter.List("next")
		blockers := adapter.List("blockers")
		scrum := generateScrum(todayList, nextList, blockers, config.ScrumWelcomeText)
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

func generateScrum(todayTasks []*common.Task, nextTasks []*common.Task, blockers []*common.Task, title string) string {
	scrumText := title + "\n" + "Today:\n"
	for _, task := range todayTasks {
		scrumText += " - " + task.Title + "\n"
	}
	scrumText += "Next:\n"
	if len(nextTasks) == 0 {
		scrumText += "  No next task for now\n"
	}
	for _, task := range nextTasks {
		scrumText += " - " + task.Title + "\n"
	}
	scrumText += "Blockers:\n"
	if len(blockers) == 0 {
		scrumText += "  None\n"
	}
	for _, task := range blockers {
		scrumText += " - " + task.Title + "\n"
	}

	return scrumText
}
