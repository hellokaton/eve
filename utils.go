package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	// Version of namebeta
	Version = "0.1"
	logo    = `
	888                          .d888                   
	888                         d88P"                    
	888                         888                      
.d88888  .d88b.  888  888       888888 888  888 88888b.  
d88" 888 d8P  Y8b 888  888       888    888  888 888 "88b 
888  888 88888888 Y88  88P       888    888  888 888  888 
Y88b 888 Y8b.      Y8bd8P        888    Y88b 888 888  888 
"Y88888  "Y8888    Y88P         888     "Y88888 888  888 
														 
Dev Fun v0.0.1

Inspired by https://github.com/biezhi/dev-fun

`
)

func displayUsage() {
	fmt.Println(logo)
	fmt.Println("Usage:")
	fmt.Println("dev-fun github [search keyword]")
	fmt.Println("dev-fun news")
	fmt.Println("dev-fun films [film name]")
}

// Options terminal args
type Options struct {
	Github bool
	News   bool
	Films  bool
	Query  string
}

// ParseArgs parse terminal arguments
func ParseArgs(args []string) *Options {
	if len(os.Args) == 1 {
		return nil
	}
	options := &Options{}

	switch args[1] {
	case "github":
		options.Github = true
		break
	case "news":
		options.News = true
		break
	case "films":
		options.Films = true
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
	body := []byte(GetRequestBody(reqURL))

	var urlArr []map[string]interface{}
	if err := json.Unmarshal(body, &urlArr); err != nil {
		fmt.Println(err)
		return ""
	}
	result := urlArr[0]["url_short"]
	return result.(string)
}

// GetRequestBody http get body
func GetRequestBody(reqURL string) string {
	resp, err := http.Get(reqURL)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
