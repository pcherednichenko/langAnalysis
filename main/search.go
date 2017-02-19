package main

import (
	"strconv"
)

type search interface {
	GetStringParams() string
}

type searchParamsHh struct {
	text string
	area int
}

type searchParamsGithub struct {
	name string
}

func (p searchParamsHh) GetStringParams() string {
	stringParams := "https://api.hh.ru/vacancies"
	stringParams = stringParams + "?text=" + p.text
	stringParams = stringParams + "&area=" + strconv.Itoa(p.area)
	return stringParams
}

func NewParamsHh(text string, area int) searchParamsHh {
	return searchParamsHh{text: text, area: area}
}

func (p searchParamsGithub) GetStringParams() string {
	if p.name == "Visual Basic" {
		p.name = "Visual+Basic"
	}
	stringParams := "https://api.github.com/search/repositories"
	stringParams = stringParams + "?q=language:" + p.name
	return stringParams
}

func NewParamsGithub(name string) searchParamsGithub {
	return searchParamsGithub{name: name}
}
