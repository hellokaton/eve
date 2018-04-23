package main

import "os"

func main() {
	options := ParseArgs(os.Args)
	if options == nil {
		displayUsage()
		os.Exit(0)
	}
	if options.Github {
		ShowGithubTrending()
	}
	if options.News {
		// TODO
	}
	if options.Films {
		// TODO
	}
}
