package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"sort"
)

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
	"Go":                 "'golang'",
	"PHP":                "'php'",
	"C++":                "'c%2B%2B'",
	"C#":                 "C%23",
	"JavaScript":         "'JavaScript'",
	"Java":               "'Java'",
	"Swift":              "'swift'",
	"Objective-C":        "'Objective-C'",
	"Python":             "'Python'",
	"Ruby":               "'Ruby'",
	"Visual Basic":       "Visual%20Basic",
	"Scala":              "Scala",
	"Perl":               "Perl",
	"Erlang":             "Erlang",
}

type githubCounter struct {
	name  string
	count float64
}

type language struct {
	name string
}

type languageResponse struct {
	area    int
	countHh map[string]interface{}
}

type Response struct {
	status int
	body   []byte
}

func (l *language) search(s search) map[string]interface{} {
	req, err := http.NewRequest("GET", s.GetStringParams(), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	Response := &Response{}
	Response.status = resp.StatusCode
	if Response.status != 200 {
		panic(Response.status)
	}
	Response.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var dat map[string]interface{}
	if err := json.Unmarshal(Response.body, &dat); err != nil {
		panic(err)
	}
	return dat
}

func main() {
	in := make(chan int, 2)
	out := make(chan languageResponse, 2)
	for range areaMap {
		go collectDataHh(in, out)
	}
	for id := range areaMap {
		in <- id
	}
	languagesByArea := make(map[int](map[string]interface{}))
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
	}
}

func collectDataHh(in <-chan int, out chan<- languageResponse) {
	area := <-in
	outArrayHh := map[string]interface{}{}
	for name, searchParams := range languages {
		lang := language{name}
		ResponseHh := lang.search(NewParamsHh(searchParams, area))
		outArrayHh[lang.name] = ResponseHh["found"]
	}
	out <- languageResponse{area, outArrayHh}
}

func NewGithubCounter(in <-chan string, out chan <- githubCounter) {
	name := <-in
	lang := language{name}
	langCount := lang.search(NewParamsGithub(name))["total_count"].(float64)
	out <- githubCounter{name, langCount}
}
