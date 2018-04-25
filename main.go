package main

import (
	"os"
	"github.com/biezhi/eve/utils"
	"fmt"
)

var (
	logo    = `
______   __   __ ______    
/\  ___\ /\ \ / //\  ___\   
\ \  __\ \ \ \'/ \ \  __\   
 \ \_____\\ \__|  \ \_____\ 
  \/_____/ \/_/    \/_____/ 
  
eve v0.0.2

everyday explore.

Inspired by https://github.com/biezhi/eve
`
)

func green(s string) {
	fmt.Println(utils.Green(s))
}

func displayUsage() {
	fmt.Println(utils.Red(logo))

	green("Usage:\n")
	green("    eve <option> [query]\n")
	green("    eve news             (show all news)")
	green("    eve github [keyword] (show github trending or search repo)")
	green("    eve v2ex             (show v2ex hot topics)")
	green("    eve hn               (show HackNews hot topics)")
	green("    eve tc               (show TechCrunch hot topics)")
	green("    eve ph               (show Product Hunt hot product)")
	green("    eve medium")
	green("    eve help             (show usage)")
}

func main() {
	options := utils.ParseArgs(os.Args)
	if options == nil {
		displayUsage()
		os.Exit(0)
	}
	if options.News {
		ShowNews()
	}
	if options.Github {
		ShowGithubInfo(options.Query)
	}
	if options.Films {
		ShowFilms(options.Query)
	}
	if options.V2EX {
		ShowHotTopic()
	}
	if options.HackNews {
		ShowHackNews()
	}
	if options.TechCrunch {
		ShowTechCrunch()
	}
}
