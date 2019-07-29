package common

import "time"

type Task struct {
	Id               string
	Title            string
	Description      string
	CreationDateTime time.Time
	StartDateTime    time.Time
	DoneDateTime     time.Time
}
