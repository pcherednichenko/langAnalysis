package search

import (
	"strconv"
)

type Search interface {
	GetStringParams() string
}

type paramsHh struct {
	text string
	area int
}

type paramsGithub struct {
	name string
}

func (p paramsHh) GetStringParams() string {
	stringParams := "https://api.hh.ru/vacancies"
	stringParams = stringParams + "?text=" + p.text
	stringParams = stringParams + "&area=" + strconv.Itoa(p.area)
	return stringParams
}

func NewParamsHh(text string, area int) paramsHh {
	return paramsHh{text: text, area: area}
}

func (p paramsGithub) GetStringParams() string {
	if p.name == "Visual Basic" {
		p.name = "Visual+Basic"
	}
	stringParams := "https://api.github.com/search/repositories"
	stringParams = stringParams + "?q=language:" + p.name + "&stars:>=3"
	return stringParams
}

func NewParamsGithub(name string) paramsGithub {
	return paramsGithub{name: name}
}
