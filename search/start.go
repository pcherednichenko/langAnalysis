package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gernest/utron/app"
	"io/ioutil"
	"langAnalysis/models"
	"net/http"
	"sort"
	"strconv"
	"time"
)

var timeDiffToUpdate int64 = 86400

var areaMap = map[int]string{
	1:   "Москва",
	2:   "Санкт-Петербург",
	3:   "Екатеринбург",
	4:   "Новосибирск",
	66:  "Нижний Новгород",
	88:  "Казань",
	78:  "Самара",
	68:  "Омск",
	104: "Челябинск",
	76:  "Ростов-на-Дону",
	99:  "Уфа",
	24:  "Волгоград",
	72:  "Пермь",
	54:  "Красноярск",
	26:  "Воронеж",
	79:  "Саратов",
	53:  "Краснодар",
	212: "Тольятти",
	96:  "Ижевск",
	98:  "Ульяновск",
	11:  "Барнаул",
	22:  "Владивосток",
	112: "Ярославль",
	35:  "Иркутск",
	95:  "Тюмень",
	29:  "Махачкала",
	102: "Хабаровск",
	70:  "Оренбург",
}

var languages = map[string]string{
	"Go":           "'golang'",
	"PHP":          "'php'",
	"C++":          "'c%2B%2B'",
	"C#":           "C%23",
	"JavaScript":   "'JavaScript'",
	"Java":         "'Java'",
	"Swift":        "'swift'",
	"Objective-C":  "'Objective-C'",
	"Python":       "'Python'",
	"Ruby":         "'Ruby'",
	"Visual Basic": "Visual%20Basic",
	"Scala":        "Scala",
	"Perl":         "Perl",
	"Erlang":       "Erlang",
	"CoffeeScript": "CoffeeScript",
	"Haskell":      "Haskell",
	"HTML":         "HTML",
	"Lua":          "Lua",
	"Matlab":       "Matlab",
	"R":            "'R'",
	"Shell":        "Shell",
	"TeX":          "TeX",
	"Dart":         "Dart",
	"Fortran":      "Fortran",
}

type githubCounter struct {
	name  string
	count int
}

type language struct {
	name string
}

type languageResponse struct {
	area    int
	countHh map[string]int
}

type Response struct {
	status int
	body   []byte
}

func Start(app *app.App) {
	settings := models.Settings{}
	app.Model.DB.Find(&settings)

	nextIteration := settings.Iteration + 1

	if !isReloadNeed(settings) {
		return
	}
	in := make(chan int, 2)
	out := make(chan languageResponse, 2)
	for range areaMap {
		go collectDataHh(in, out)
	}
	for id := range areaMap {
		in <- id
	}
	languagesByArea := make(map[int](map[string]int))
	for range areaMap {
		languagesByAreaTemp := <-out
		languagesByArea[languagesByAreaTemp.area] = languagesByAreaTemp.countHh
	}

	//Sort map
	var keys []int
	for k := range languagesByArea {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(areaMap[k])
		for name, count := range languagesByArea[k] {
			fmt.Println("  ", name, " : ", count)
			hhModel := models.HhBase{
				Date:         time.Now(),
				CityName:     areaMap[k],
				LanguageName: name,
				Counts:       count,
				Iteration:    nextIteration,
			}
			app.Model.DB.Create(&hhModel)
		}
	}

	githubLanguages := make([]githubCounter, len(languages))
	inGit := make(chan string, 2)
	outGit := make(chan githubCounter, 2)
	for range languages {
		go NewGithubCounter(inGit, outGit)
	}
	for name := range languages {
		inGit <- name
	}
	i := 0
	for range languages {
		githubLanguages[i] = <-outGit
		i++
	}
	fmt.Println("________")
	fmt.Println("GitHub language analysis:")
	for _, language := range githubLanguages {
		fmt.Println(" ", language.name, " : ", language.count)
		modelGithub := models.GitHubBase{
			Date:         time.Now(),
			LanguageName: language.name,
			Counts:       language.count,
			Iteration:    nextIteration,
		}
		app.Model.DB.Create(&modelGithub)
	}

	settingsNew := models.Settings{LastUpdateTime: time.Now(), Iteration: nextIteration}
	app.Model.DB.Create(settingsNew)
}

func (l *language) search(s Search) (dat map[string]interface{}, err error) {
	req, err := http.NewRequest("GET", s.GetStringParams(), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	Response := &Response{}
	Response.status = resp.StatusCode
	if Response.status != 200 {
		err = errors.New(strconv.Itoa(Response.status))
		return
	}
	Response.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(Response.body, &dat)
	if err != nil {
		return
	}
	return
}

func collectDataHh(in <-chan int, out chan<- languageResponse) {
	area := <-in
	outArrayHh := map[string]int{}
	for name, searchParams := range languages {
		if name == "R" {
			continue
		}
		lang := language{name}
		ResponseHh, err := lang.search(NewParamsHh(searchParams, area))
		if err != nil {
			close(out)
			fmt.Println(err)
		}
		outArrayHh[lang.name] = int(ResponseHh["found"].(float64))
	}
	out <- languageResponse{area, outArrayHh}
}

func NewGithubCounter(in <-chan string, out chan<- githubCounter) {
	name := <-in
	lang := language{name}
	langCount, err := lang.search(NewParamsGithub(name))
	if err != nil {
		close(out)
		fmt.Println(err)
	}
	out <- githubCounter{name, int(langCount["total_count"].(float64))}
}

func isReloadNeed(settings models.Settings) bool {
	if time.Now().Unix() < settings.LastUpdateTime.Unix()+timeDiffToUpdate {
		fmt.Println("Данные по языкам актуальны, последнее обновление: ")
		fmt.Println(settings.LastUpdateTime)
		return false
	}
	fmt.Println("Обновляем данные по языкам")
	return true
}
