package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/andygrunwald/go-trending"
	"github.com/biezhi/moe"
	"github.com/olekukonko/tablewriter"
)

func main() {
	trend := trending.NewTrending()

	moe := moe.New("loading").Color(moe.Green).Start()
	// Show projects of today
	projects, err := trend.GetProjects(trending.TimeToday, "")

	if err != nil {
		log.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Project", "Stars", "Description", "URL"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)

	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgBlueColor})

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor})

	var shortUrls []string = make([]string, len(projects))
	for index, project := range projects {
		URL := getShortURL(project.URL.String())
		shortUrls[index] = URL
	}

	moe.Stop()

	for index, project := range projects {
		Repo := project.RepositoryName
		Desc := project.Description
		if len(Repo) > 20 {
			Repo = string([]rune(Repo)[:20]) + "..."
		}
		if len(Desc) > 25 {
			Desc = string([]rune(Desc)[:25]) + "..."
		}
		URL := shortUrls[index]
		row := []string{Repo, strconv.Itoa(project.Stars), Desc, URL}
		table.Append(row)
	}
	table.Render()

}

type SinaUrl struct {
	UrlShort string `json:url_short`
	UrlLong  string `json:url_long`
	Type     int    `json:type`
}

func getShortURL(url string) string {
	endpoint := "http://api.t.sina.com.cn/short_url/shorten.json?source=3271760578&url_long=" + url
	// fmt.Println("Url:", endpoint)
	resp, _ := http.Get(endpoint)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var ms []map[string]interface{}

	if err := json.Unmarshal(body, &ms); err != nil {
		fmt.Println(err)
		return ""
	}
	result := ms[0]["url_short"]
	return result.(string)
}
