package models

import "time"

type HhBase struct {
	Date         time.Time
	CityName     string
	LanguageName string
	Counts       int
	Iteration    int
}
