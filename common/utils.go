package common

import "time"

func GetNow() time.Time {
	return time.Now()
}

func GetTomorrow() time.Time {
	return time.Now().Add(time.Hour * 24)
}
