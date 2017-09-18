package main

import (
	"fmt"
	"safari_downloader/conf"
	"safari_downloader/jobs"
)

func main() {
	fmt.Println("#---------safari-downloader----------#")
	//------->configuration
	config := conf.ReadConfig()

	//------> Create file
	err := jobs.CreateDataFile(&config)
	if err != nil {
		config.Logger.Println(err.Error)
	}
}
