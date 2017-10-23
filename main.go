package main

import (
	"bufio"
	"net/url"
	"os"
	"safari_downloader/conf"
	"safari_downloader/jobs"
	"strings"

	"github.com/fatih/color"
	"github.com/hifx/banner"
)

//UserInputs represent user input values
type UserInputs struct {
	Username string
	Password string
	Link     string
	DType    string
	Dest     string
}

func main() {
	//-------banner
	printName("getfile")

	msg := color.New(color.Bold, color.FgGreen).PrintFunc()
	errfun := color.New(color.FgRed).PrintlnFunc()

	reader := bufio.NewReader(os.Stdin)
	inputs := UserInputs{}

	//---- username
	for inputs.Username == "" {
		msg("\nEnter username :")
		inputs.Username, _ = reader.ReadString('\n')
		inputs.Username = strings.TrimSpace(inputs.Username)
		if inputs.Username == "" {
			errfun("\nEnter username for safaribooksonline!")
		}
	}
	//---- password
	for inputs.Password == "" {
		msg("\nEnter password :")
		inputs.Password, _ = reader.ReadString('\n')
		inputs.Password = strings.TrimSpace(inputs.Password)
		if inputs.Password == "" {
			errfun("\nEnter password for safaribooksonline !")
		}
	}
	//---- link
	var validLink bool
	for inputs.Link == "" || !validLink {
		msg("\nEnter url :")
		inputs.Link, _ = reader.ReadString('\n')
		inputs.Link = strings.TrimSpace(inputs.Link)
		if inputs.Link == "" {
			errfun("\nEnter URL to home page of tutorial!")
		} else {
			_, err := url.ParseRequestURI(inputs.Link)
			if err != nil {
				errfun("\nEnter valid URL!")
			} else {
				//--- check this is url for home page of tutorials
				if strings.Contains(inputs.Link, "html") {
					errfun("\nEnter URL to home page for getting all videos!")
				} else {
					validLink = true
				}
			}
		}
	}
	//---- document type
	var validDType bool
	for inputs.DType == "" || !validDType {
		msg("\nEnter document type(video/pdf) :")
		inputs.DType, _ = reader.ReadString('\n')
		inputs.DType = strings.TrimSpace(inputs.DType)
		if inputs.DType == "" {
			errfun("\nEnter document type!")
		} else {
			if inputs.DType != "video" && inputs.DType != "pdf" {
				errfun("\nEnter valid document type!")
			} else {
				validDType = true
			}
		}
	}
	//--- custom location need or not?
	var res string
	var validRes bool
	if res == "" || !validRes {
		msg("\n you need any custom location[default : $HOME/Documents/safari] for store files[y/n] :")
		res, _ = reader.ReadString('\n')
		res = strings.TrimSpace(res)
		if res == "" {
			errfun("please enter response!")
		} else {
			if res != "n" && res != "y" {
				errfun("please enter valid response!")
			} else {
				validRes = true
			}
		}
	}
	//--- read custom location
	if res == "y" {
		var validDest bool
		for inputs.Dest == "" || !validDest {
			msg("\n Enter custom location :")
			inputs.Dest, _ = reader.ReadString('\n')
			inputs.Dest = strings.TrimSpace(inputs.Dest)
			if inputs.Dest == "" {
				errfun("please enter custom location!")
			} else {
				_, err := os.Stat(inputs.Dest)
				if err != nil {
					errfun("enter valid custom location!")
				} else {
					validDest = true
				}
			}
		}

	}
	//-----read configuration
	config, err := conf.ReadConfig(&inputs.Link, &inputs.Username, &inputs.Password, &inputs.Dest, &inputs.DType)
	if err != nil {
		errfun(err.Error())
		config.Logger.Println(err.Error())
		os.Exit(-1)
	}
	//-------create files
	err = jobs.CreateDataFile(&config)
	if err != nil { // err = jobs.DownloadFiles(&config)
		// if err != nil {
		// 	errfun(err.Error())
		// 	config.Logger.Println(err.Error())
		// 	os.Exit(-1)
		// }

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
