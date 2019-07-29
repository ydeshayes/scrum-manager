package trelloAdapter

import common "github.com/ydeshayes/dev-diary/common"

type Configuration struct {
	Apikey        string
	Token         string
	TrelloBoardId string
}

func getConfiguration(config common.Configuration) Configuration {
	return Configuration{config.TrelloApikey, config.TrelloToken, config.TrelloBoardId}
}
