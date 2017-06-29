package controllers

import (
	"net/http"

	"github.com/gernest/utron/controller"
	"langAnalysis/models"
)

//Index is a controller for Index list
type Index struct {
	controller.BaseController
	Routes []string
}

type countOfLang struct {
	langName string
	count    int
}

//Home page
func (t *Index) Home() {
	hhBase := []*models.HhBase{}
	t.Ctx.DB.Order("counts desc").Find(&hhBase)
	t.Ctx.Data["HHList"] = hhBase
	gitHubBase := []*models.GitHubBase{}
	t.Ctx.DB.Order("counts desc").Find(&gitHubBase)
	t.Ctx.Data["GithubList"] = gitHubBase
	t.Ctx.Template = "index"
	t.HTML(http.StatusOK)
}

//API returns HH data
func (t *Index) Hh() {
	hhBase := []*models.HhBase{}
	t.Ctx.DB.Order("date desc").Find(&hhBase)
	response := make(map[string]map[string][]int)
	for _, hhElem := range hhBase {
		lang := make(map[string][]int)
		for _, hhElemIn := range hhBase {
			if hhElemIn.CityName == hhElem.CityName {
				lang[hhElemIn.LanguageName] = append(lang[hhElemIn.LanguageName], hhElemIn.Counts)
			}
		}
		response[hhElem.CityName] = lang
	}
	t.RenderJSON(response, http.StatusOK)
}

//Controller returns a new controller
func Controller() controller.Controller {
	return &Index{
		Routes: []string{
			"get;/;Home",
			"get;/api/hhdata;Hh",
		},
	}
}
