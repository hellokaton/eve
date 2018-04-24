package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/biezhi/moe"
	"github.com/olekukonko/tablewriter"
)

func ShowNews() {
	moe := moe.New("loading all news...").Color(moe.Green).Start()
	url := "http://reader.one/api/all/hn,reddit,ph,dn,github,medium,lifehacker?limit=20"
	body := GetRequestBody(url)

	var resp []BaseArticle
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

	shortUrls := make([]string, len(resp))
	for index, article := range resp {
		URL := GetShortURL(article.URL)
		shortUrls[index] = URL
	}

	moe.Stop()
	for index, article := range resp {
		URL := shortUrls[index]
		Title := article.Title
		// if len(Title) > 75 {
		// 	Title = string([]rune(Title)[:72]) + "..."
		// }
		row := []string{Title, URL}
		table.Append(row)
	}
	table.Render()
}
