package common

import (
	"os"
	"path"

	gonfig "github.com/tkanos/gonfig"
)

type Configuration struct {
	AdapterName   string
	TrelloApikey  string
	TrelloToken   string
	TrelloBoardId string

	ScrumWelcomeText string
}

func GetConfiguration() Configuration {
	configuration := Configuration{}
	homePath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	err = gonfig.GetConf(path.Join(homePath, ".scrum-manager"), &configuration)
	if err != nil {
		panic(err)
	}

	return configuration
}
