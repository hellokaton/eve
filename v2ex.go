package main

import (
	"encoding/json"
	"os"

	"github.com/biezhi/moe"
	"github.com/olekukonko/tablewriter"
)

// GetV2EX get hot topics
func getV2EX() []BaseArticle {
	url := "https://www.v2ex.com/api/topics/hot.json"
	body := GetRequestBody(url)

	var resp []BaseArticle
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
		URL := GetShortURL(topic.URL)
		shortUrls[index] = URL
	}

	moe.Stop()
	for index, topic := range topics {
		URL := shortUrls[index]
		Title := RemoveSpace(topic.Title)
		row := []string{Title, URL}
		table.Append(row)
	}
	table.Render()
}
