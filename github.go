package main

import (
	"log"
	"os"
	"strconv"

	trending "github.com/andygrunwald/go-trending"
	"github.com/biezhi/moe"
	"github.com/olekukonko/tablewriter"
)

// ShowGithubTrending show github trending
// default is today
func ShowGithubTrending() {
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

	shortUrls := make([]string, len(projects))
	for index, project := range projects {
		URL := GetShortURL(project.URL.String())
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
