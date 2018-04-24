package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/biezhi/moe"
	"github.com/olekukonko/tablewriter"
)

func ShowTechCrunch() {
	moe := moe.New("loading techcrunch...").Color(moe.Green).Start()
	url := "https://newsapi.org/v2/top-headlines?sources=techcrunch-cn&apiKey=b77ad5166e264aa999d117a8ca7ccba6"
	body := GetRequestBody(url)
	var resp HackNewsResp
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Fatalln("load techcrunch fail")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Title", "URL"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)

	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlackColor})

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor})

	shortUrls := make([]string, len(resp.Articles))
	for index, article := range resp.Articles {
		URL := GetShortURL(article.URL)
		shortUrls[index] = URL
	}

	moe.Stop()
	for index, article := range resp.Articles {
		URL := shortUrls[index]
		Title := strings.Replace(article.Title, " | TechCrunch 中文版", "", -1)
		if len(Title) > 75 {
			Title = string([]rune(Title)[:72]) + "..."
		}
		row := []string{Title, URL}
		table.Append(row)
	}
	table.Render()

}
