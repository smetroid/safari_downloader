package main

import (
	"flag"
	"os"
	"safari_downloader/conf"
	"safari_downloader/jobs"
	"strings"

	"github.com/fatih/color"
	"github.com/hifx/banner"
)

func main() {
	//-------banner
	printName("getfile")

	msg := color.New(color.Bold, color.FgHiBlue).PrintlnFunc()
	errfun := color.New(color.Bold, color.FgHiRed).PrintlnFunc()

	//-------command arguments
	link := flag.String("l", "", "hint : -l  https://www.safaribooksonline.com [url]")
	username := flag.String("u", "", "hint : -u username [username]")
	password := flag.String("p", "", "hint : -p password [password]")
	dest := flag.String("d", "", "hint : -p password [password]")

	flag.Parse()
	//------url
	if *link == "" {
		errfun("Error : please provide site url")
		os.Exit(-1)
	} else {
		i := strings.Index(*link, ".html")
		if i > 0 {
			msg("Info : enter home page url for downloading all videos")
		}
	}
	//------username
	if *username == "" {
		errfun("Error : please provide username")
		os.Exit(-1)
	}
	//------password
	if *password == "" {
		errfun("Error : please provide password")
		os.Exit(-1)
	}

	//-----read configuration
	config, err := conf.ReadConfig(link, username, password, dest)
	if err != nil {
		errfun(err.Error())
		config.Logger.Println(err.Error())
		os.Exit(-1)
	}
	//-------create files
	err = jobs.CreateDataFile(&config)
	if err != nil {
		errfun(err.Error())
		config.Logger.Println(err.Error())
		os.Exit(-1)
	}

	//------download all files
	err = jobs.DownloadFiles(&config)
	if err != nil {
		errfun(err.Error())
		config.Logger.Println(err.Error())
		os.Exit(-1)
	}
}

//PrintName prints the app name
func printName(str string) {
	color.New(color.FgCyan).Add(color.Bold).Println(banner.PrintS(str))
}
