package trelloAdapter

import (
	"encoding/json"
	"fmt"
	"time"

	trello "github.com/adlio/trello"
	common "github.com/ydeshayes/dev-diary/common"
)

type TrelloAdapter struct {
	config Configuration
	client *trello.Client
	board  *trello.Board
}

func (t *TrelloAdapter) Initialize(configuration common.Configuration) {
	t.config = getConfiguration(configuration)
	t.client = trello.NewClient(t.config.Apikey, t.config.Token)

	board, err := t.client.GetBoard(t.config.TrelloBoardId, trello.Defaults())
	if err != nil {
		fmt.Println("Error loading the board " + t.config.TrelloBoardId)
	}
	t.board = board
}

func (t TrelloAdapter) Add(description string, listName string) {
	lists, err := t.board.GetLists(trello.Defaults())
	if err != nil {
		fmt.Println("Error getting the lists")
	}
	if len(lists) == 0 {
		// Create the lists
		lists, err = t.createDefaultLists()
	}

	list, err := t.getlistByName(listName, lists)
	if err != nil {
		fmt.Println("Error loading the list todos")
	}

	card := trello.Card{
		Name:   description,
		Desc:   description,
		IDList: list.ID,
	}

	err = t.client.CreateCard(&card, trello.Defaults())

	if err != nil {
		fmt.Println("Error creating card ")
		fmt.Println(err)
	}
}

func (t TrelloAdapter) List(name string) []*common.Task {
	lists, err := t.board.GetLists(trello.Defaults())
	if err != nil {
		fmt.Println("Error getting the list")
		fmt.Println(err)
	}

	list, err := t.getlistByName(name, lists)
	cards, err := list.GetCards(trello.Defaults())

	return cardsListToTasksList(cards)
}

func (t TrelloAdapter) Move(task common.Task, newListName string) error {
	lists, err := t.board.GetLists(trello.Defaults())
	if err != nil {
		fmt.Println("Error getting the list")
		fmt.Println(err)
		return err
	}
	list, err := t.getlistByName(newListName, lists)
	card, err := t.client.GetCard(task.Id, trello.Defaults())
	if err != nil {
		fmt.Println("Error getting the card " + task.Id)
		return err
	}

	card.MoveToList(list.ID, trello.Defaults())

	return err
}

func (t TrelloAdapter) createDefaultLists() (lists []*trello.List, err error) {
	list, err := t.board.CreateList("todos", trello.Arguments{"pos": "1"})
	lists = append(lists, list)

	if err != nil {
		fmt.Println("Error creating list todos")
		fmt.Println(err)
	}

	list, err = t.board.CreateList("next", trello.Arguments{"pos": "2"})
	lists = append(lists, list)

	if err != nil {
		fmt.Println("Error creating list in_progress")
		fmt.Println(err)
	}

	list, err = t.board.CreateList("today", trello.Arguments{"pos": "3"})
	lists = append(lists, list)

	if err != nil {
		fmt.Println("Error creating list done")
		fmt.Println(err)
	}

	list, err = t.board.CreateList("archives", trello.Arguments{"pos": "4"})
	lists = append(lists, list)

	if err != nil {
		fmt.Println("Error creating list archives")
		fmt.Println(err)
	}

	return
}

func (t TrelloAdapter) getlistByName(name string, lists []*trello.List) (*trello.List, error) {
	for _, list := range lists {
		if list.Name == name {
			return list, nil
		}
	}

	list, err := t.board.CreateList(name, trello.Defaults())

	return list, err
}

func (t TrelloAdapter) NextScrum() {
	lists, err := t.board.GetLists(trello.Defaults())
	todayList, err := t.getlistByName("today", lists)
	nextList, err := t.getlistByName("next", lists)
	blockersList, err := t.getlistByName("blockers", lists)
	archivedList, err := t.getlistByName("archived", lists)

	if err != nil {
		fmt.Println("Error getting lists")
		fmt.Println(err)
	}
	t.archiveList(*todayList, *blockersList, *archivedList)
	moveAllCardsBetweenLists(*nextList, todayList.ID)
}

func (t TrelloAdapter) archiveList(todayList trello.List, blockersList trello.List, archiveList trello.List) {
	cards, err := todayList.GetCards(trello.Arguments{"fields": "idShort,name"})
	if err != nil {
		fmt.Println("Error getting today list cards")
		fmt.Println(err)
	}

	blockers, err := blockersList.GetCards(trello.Defaults())
	if err != nil {
		fmt.Println("Error getting blockers list cards")
		fmt.Println(err)
	}

	// Serialize cards in json?
	serializedTodayCards, err := json.Marshal(cardsListToArchivedTasksList(cards))
	if err != nil {
		fmt.Println("Error serialize cards")
		fmt.Println(err)
	}
	serializedBlockersCards, err := json.Marshal(cardsListToArchivedTasksList(blockers))
	if err != nil {
		fmt.Println("Error serialize cards")
		fmt.Println(err)
	}
	date := common.GetNow()
	card := &trello.Card{
		Name:   date.Format("02-01-2006"),
		Desc:   string(serializedTodayCards) + " {------} " + string(serializedBlockersCards),
		IDList: archiveList.ID,
	}

	err = t.client.CreateCard(card, trello.Defaults())
	if err != nil {
		fmt.Println("Error Archiving!")
		fmt.Println(err)
	}
	for _, card := range cards {
		card.Update(trello.Arguments{"closed": "true"})
	}
}

func (t TrelloAdapter) LastScrumDate() time.Time {
	lists, err := t.board.GetLists(trello.Defaults())
	if err != nil {
		fmt.Println("Error getting the list")
		fmt.Println(err)
	}

	list, err := t.getlistByName("archived", lists)
	cards, _ := list.GetCards(trello.Arguments{"limit": "1"})

	return cards[0].CreatedAt()
}
