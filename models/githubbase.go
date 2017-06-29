package models

import "time"

type GitHubBase struct {
	Date         time.Time
	LanguageName string
	Counts       int
	Iteration    int
}
