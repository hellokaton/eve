package main

import "os"

func main() {
	options := ParseArgs(os.Args)
	if options == nil {
		displayUsage()
		os.Exit(0)
	}
	if options.Github {
		ShowGithubInfo(options.Query)
	}
	if options.News {
		ShowNews()
	}
	if options.Films {
		ShowFilms(options.Query)
	}
	if options.V2EX {
		ShowHotTopic()
	}
}
