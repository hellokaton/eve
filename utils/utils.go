package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"reflect"
)

// BaseArticle topic
type BaseArticle struct {
	Title   string `json:"title,omitempty"`
	URL     string `json:"url,omitempty"`
	Content string `json:"content,omitempty"`
}

// Options terminal args
type Options struct {
	News        bool
	Github      bool
	Films       bool
	V2EX        bool
	Zhihu       bool
	JianDan     bool
	HackNews    bool
	Reddit      bool
	ProductHunt bool
	Medium      bool
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
	case "zhihu":
		options.Zhihu = true
		break
	case "jiandan":
		options.JianDan = true
		break
	case "reddit":
		options.Reddit = true
		break
	case "medium":
		options.Medium = true
		break
	case "v2ex":
		options.V2EX = true
		break
	case "hn":
		options.HackNews = true
		break
	case "ph":
		options.ProductHunt = true
		break
	case "help":
		return nil
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
	if url == "" {
		return ""
	}
	var urlArr []map[string]interface{}
	body := GetRequestBody(reqURL)
	if err := json.Unmarshal(body, &urlArr); err != nil {
		fmt.Println("获取短链接失败", string(body), "\n", err)
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

func MapToString(structs []BaseArticle, fieldName string) []string {
	result := make([]string, len(structs))
	for _, item := range structs {
		v := reflect.ValueOf(item)
		f := v.FieldByName(fieldName)
		result = append(result, f.String())
	}
	return result
}

var wg sync.WaitGroup

func GetShortURLArray(longUrls []string) map[string]*string {

	shortUrls := make(map[string]*string, len(longUrls))
	for _, url := range longUrls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			URL := GetShortURL(url)
			shortUrls[url] = &URL
		}(url)
	}
	wg.Wait()
	return shortUrls
}
