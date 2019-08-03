package common

import "time"

func GetNow() time.Time {
	return time.Now()
}

func GetTomorrow() time.Time {
	return time.Now().Add(time.Hour * 24)
}

func BeginingOfDay() time.Time {
	today := time.Now()
	year, month, day := today.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, today.Location())
}

func EndOfDay(t time.Time) time.Time {
	today := time.Now()
	year, month, day := today.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, today.Location())
}
