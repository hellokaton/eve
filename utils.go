package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	// Version of namebeta
	Version = "0.1"
	logo    = `
______   __   __ ______    
/\  ___\ /\ \ / //\  ___\   
\ \  __\ \ \ \'/ \ \  __\   
 \ \_____\\ \__|  \ \_____\ 
  \/_____/ \/_/    \/_____/ 
  
eve v0.0.1

everyday explore.

Inspired by https://github.com/biezhi/eve
`
)

// BaseArticle topic
type BaseArticle struct {
	Title   string `json:"title,omitempty"`
	URL     string `json:"url,omitempty"`
	Content string `json:"content,omitempty"`
}

func displayUsage() {
	fmt.Println(logo)
	fmt.Println("Usage:")
	fmt.Println("eve news")
	fmt.Println("eve github [search keyword]")
	fmt.Println("eve v2ex")
	fmt.Println("eve hn (HackNews)")
	fmt.Println("eve tc (TechCrunch)")
	// fmt.Println("eve zhihu")
	fmt.Println("eve ph (Product Hunt)")
	fmt.Println("eve medium")
}

// Options terminal args
type Options struct {
	News        bool
	Github      bool
	Films       bool
	V2EX        bool
	HackNews    bool
	TechCrunch  bool
	ProductHunt bool
	Query       string
}

// ParseArgs parse terminal arguments
func ParseArgs(args []string) *Options {
	if len(os.Args) == 1 {
		return nil
	}
	options := &Options{}

	switch args[1] {
	case "news":
		options.News = true
		break
	case "github":
		options.Github = true
		break
	case "films":
		options.Films = true
		break
	case "v2ex":
		options.V2EX = true
		break
	case "hn":
		options.HackNews = true
		break
	case "tc":
		options.TechCrunch = true
		break
	case "ph":
		options.ProductHunt = true
		break
	default:
		break
	}
	if len(os.Args) == 3 {
		options.Query = args[2]
	}
	return options
}

// GetShortURL get sina short url
func GetShortURL(url string) string {
	reqURL := "http://api.t.sina.com.cn/short_url/shorten.json?source=3271760578&url_long=" + url
	var urlArr []map[string]interface{}
	if err := json.Unmarshal([]byte(GetRequestBody(reqURL)), &urlArr); err != nil {
		fmt.Println(err)
		return ""
	}
	return urlArr[0]["url_short"].(string)
}

// GetRequestBody http get body
func GetRequestBody(reqURL string) []byte {
	resp, err := http.Get(reqURL)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func RemoveSpace(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	return s
}
