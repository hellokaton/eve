package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	trending "github.com/andygrunwald/go-trending"
	"github.com/biezhi/moe"
	"github.com/olekukonko/tablewriter"
)

// Repository represents a GitHub repository.
type Repository struct {
	ID              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	FullName        string `json:"full_name,omitempty"`
	Description     string `json:"description,omitempty"`
	Homepage        string `json:"homepage,omitempty"`
	HTMLURL         string `json:"html_url,omitempty"`
	GitURL          string `json:"git_url,omitempty"`
	Language        string `json:"language,omitempty"`
	ForksCount      int    `json:"forks_count,omitempty"`
	StargazersCount int    `json:"stargazers_count,omitempty"`
}

// GithubRepoResp github search response
type GithubRepoResp struct {
	TotalCount int
	Items      []Repository
}

// ShowGithubInfo show github trending or query repo
// default is today
func ShowGithubInfo(q string) {
	if q != "" {
		queryRepo(q)
		return
	}
	trend := trending.NewTrending()
	moe := moe.New("show github trending").Color(moe.Green).Start()
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

func queryRepo(q string) {
	moe := moe.New("query [" + q + "]").Color(moe.Green).Start()
	url := "https://api.github.com/search/repositories?q=" + q
	body := GetRequestBody(url)

	var resp GithubRepoResp
	if err := json.Unmarshal([]byte(body), &resp); err != nil {

	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Project", "Stars", "Language", "Description", "URL"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)

	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgBlueColor})

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor})

	shortUrls := make([]string, len(resp.Items))
	for index, repo := range resp.Items {
		URL := GetShortURL(repo.HTMLURL)
		shortUrls[index] = URL
	}
	moe.Stop()
	for index, repo := range resp.Items {
		Repo := repo.Name
		Desc := repo.Description
		Stars := strconv.Itoa(repo.StargazersCount)
		if len(Repo) > 20 {
			Repo = string([]rune(Repo)[:20]) + "..."
		}
		if len(Desc) > 25 {
			Desc = string([]rune(Desc)[:25]) + "..."
		}
		URL := shortUrls[index]
		row := []string{Repo, Stars, repo.Language, Desc, URL}
		table.Append(row)
	}
	table.Render()
}
