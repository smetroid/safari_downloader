package main

import (
	"safari_downloader/conf"
	"safari_downloader/jobs"

	"github.com/fatih/color"
	"github.com/hifx/banner"
)

func main() {
	printName("getfile")

	info := color.New(color.Bold, color.FgHiMagenta).PrintlnFunc()
	errfun := color.New(color.Bold, color.FgHiRed).PrintlnFunc()

	info("#step  :  safari_downloader end.")
	config := conf.ReadConfig()

	//-------create file
	err := jobs.CreateDataFile(&config)
	if err != nil {
		errfun(err.Error())
		config.Logger.Println(err.Error())
	}

	//------download all files
	err = jobs.DownloadFiles(&config)
	if err != nil {
		errfun(err.Error())
		config.Logger.Println(err.Error())
	}
}

//PrintName prints the app name
func printName(str string) {
	color.New(color.FgCyan).Add(color.Bold).Println(banner.PrintS(str))
}
