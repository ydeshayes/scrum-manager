package trelloAdapter

import (
	"fmt"

	trello "github.com/adlio/trello"
	common "github.com/ydeshayes/dev-diary/common"
)

func moveAllCardsBetweenLists(list trello.List, newListId string) {
	cards, err := list.GetCards(trello.Defaults())

	if err != nil {
		fmt.Println("Error getting list cards")
		fmt.Println(err)
	}

	for _, card := range cards {
		card.MoveToList(newListId, trello.Defaults())
	}
}

func moveCardToList(card trello.Card, newListId string) {
	card.MoveToList(newListId, trello.Defaults())
}

func cardsListToTasksList(cards []*trello.Card) []*common.Task {
	var tasks []*common.Task

	for _, card := range cards {
		tasks = append(tasks, cardToTask(*card))
	}

	return tasks
}

func cardsListToArchivedTasksList(cards []*trello.Card) []*common.ArchivedTask {
	var tasks []*common.ArchivedTask

	for _, card := range cards {
		tasks = append(tasks, &common.ArchivedTask{Id: card.ID})
	}

	return tasks
}

func cardToTask(card trello.Card) *common.Task {
	task := common.Task{Id: card.ID, Title: card.Name, Description: card.Desc, CreationDateTime: card.CreatedAt(), StartDateTime: card.CreatedAt()}

	if card.DueComplete {
		task.DoneDateTime = *card.DateLastActivity
	}

	return &task
}
