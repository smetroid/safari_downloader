package jobs

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"safari_downloader/conf"
	"strings"

	"github.com/fatih/color"
)

//CreateDataFile create formated files with neccessary data from url
func CreateDataFile(config *conf.Config) error {
	errfun := color.New(color.Bold, color.FgHiRed).PrintlnFunc()

	//---------get url
	res, err := http.Get(config.Url)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	//-------- read all contents from url
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fullData := string(html)
	index := strings.Index(fullData, "Table of Contents")
	if index < 0 {
		errfun("Table of Contents not exists")
		return errors.New("Table of Contents not exists")
	} else if !strings.Contains(fullData, "href") {
		return errors.New("There is no documents exits in this link")
	}

	subData := fullData[index:]
	re, err := regexp.Compile(`<a.*>`)
	if err != nil {
		return err
	}
	//-----------remove file if exist
	if _, err := os.Stat(config.DataFile); os.IsExist(err) {
		err = os.Remove(config.DataFile)
		if err != nil {
			return err
		}
	}
	//---------create new result file
	dataFile, err := os.Create(config.DataFile)
	if err != nil {
		return err
	}
	defer dataFile.Close()

	result := re.FindAllString(subData, -1)
	for _, v := range result {
		//----------extract url from the line
		urlreg, err := regexp.Compile(`href=\".*html\"`)
		if err != nil {
			return err
		}
		url := urlreg.FindString(v)

		finalurl := strings.TrimLeft(strings.TrimLeft(url, "href="), " ")
		if len(finalurl) != 0 {
			dataFile.WriteString("l=" + finalurl + "\n")

		}
		//---------extract heading from the line
		headreg, err := regexp.Compile(`>.*<`)
		if err != nil {
			return err
		}
		//--------checking for empty heading
		h := headreg.FindString(v)
		if len(h) != 0 {
			dataFile.WriteString("f=" + h[1:len(h)-1] + "\n")
		}
	}
	return nil
}
