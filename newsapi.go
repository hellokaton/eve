package main

import (
	"github.com/biezhi/moe"
	"github.com/biezhi/eve/utils"
	"encoding/json"
	"log"
	"github.com/olekukonko/tablewriter"
	"os"
)

type HackNewsResp struct {
	Articles     []utils.BaseArticle `json:"articles,omitempty"`
	Status       string              `json:"status,omitempty"`
	TotalResults int                 `json:"totalResults,omitempty"`
}

func ShowNews(q string) {
	tip := "loading all news..."
	url := "http://reader.one/api/all/hn,reddit,ph,dn,github,medium,lifehacker?limit=20"
	if q != "" {
		tip = "loading " + q + " ..."
		url = "http://reader.one/api/news/" + q + "?limit=20"
	}

	moe := moe.New(tip).Color(moe.Green).Start()
	body := utils.GetRequestBody(url)

	var resp []utils.BaseArticle
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Fatalln("load news fail")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(70)

	table.SetHeader([]string{"Title", "URL"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)

	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlackColor})

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor})

	items := utils.MapToString(resp, "URL")
	shortUrls := utils.GetShortURLArray(items)

	moe.Stop()
	for _, article := range resp {
		URL := shortUrls[article.URL]
		Title := article.Title
		row := []string{Title, *URL}
		table.Append(row)
	}
	table.Render()
}

func ShowNewsApi(source string) {
	moe := moe.New("loading " + source + "...").Color(moe.Green).Start()
	url := "https://newsapi.org/v2/top-headlines?sources=" + source + "&apiKey=b77ad5166e264aa999d117a8ca7ccba6"
	body := utils.GetRequestBody(url)
	var resp HackNewsResp
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Fatalln("load " + source + " fail")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(70)

	table.SetHeader([]string{"Title", "URL"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)

	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlackColor})

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor})

	items := utils.MapToString(resp.Articles, "URL")
	shortUrls := utils.GetShortURLArray(items)

	moe.Stop()
	for _, article := range resp.Articles {
		URL := shortUrls[article.URL]
		Title := article.Title
		if len(Title) > 75 {
			Title = string([]rune(Title)[:72]) + "..."
		}
		row := []string{Title, *URL}
		table.Append(row)
	}
	table.Render()

}

// GetV2EX get hot topics
func getV2EX() []utils.BaseArticle {
	url := "https://www.v2ex.com/api/topics/hot.json"
	body := utils.GetRequestBody(url)

	var resp []utils.BaseArticle
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil
	}
	return resp
}

func ShowHotTopic() {
	moe := moe.New("loading v2ex...").Color(moe.Green).Start()

	topics := getV2EX()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(70)

	table.SetHeader([]string{"Title", "URL"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)

	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlackColor})

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor})

	shortUrls := make([]string, len(topics))
	for index, topic := range topics {
		URL := utils.GetShortURL(topic.URL)
		shortUrls[index] = URL
	}

	moe.Stop()
	for index, topic := range topics {
		URL := shortUrls[index]
		Title := utils.RemoveSpace(topic.Title)
		if len(Title) > 34 {
			Title = string([]rune(Title)[:31]) + "..."
		}
		row := []string{Title, URL}
		table.Append(row)
	}
	table.Render()
}
