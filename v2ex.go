package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/biezhi/moe"
	"github.com/olekukonko/tablewriter"
)

// Topic topic
type Topic struct {
	Title   string `json:"title,omitempty"`
	URL     string `json:"url,omitempty"`
	Content string `json:"content,omitempty"`
}

// GetV2EX get hot topics
func getV2EX() []Topic {
	url := "https://www.v2ex.com/api/topics/hot.json"
	body := GetRequestBody(url)

	var resp []Topic
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil
	}
	return resp
}

func ShowHotTopic() {

	moe := moe.New("正在加载 v2ex...").Color(moe.Green).Start()

	topics := getV2EX()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"标题", "地址"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)

	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor})

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor})

	shortUrls := make([]string, len(topics))
	for index, topic := range topics {
		URL := GetShortURL(topic.URL)
		shortUrls[index] = URL
	}

	moe.Stop()
	for index, topic := range topics {
		URL := shortUrls[index]
		Title := topic.Title
		// 去除空格
		Title = strings.Replace(Title, " ", "", -1)
		// 去除换行符
		Title = strings.Replace(Title, "\n", "", -1)
		fmt.Println(len(Title))
		if len(Title) > 25 {
			Title = string([]rune(Title)[:24]) + "..."
		}
		row := []string{Title, URL}
		table.Append(row)
	}
	table.Render()
}
