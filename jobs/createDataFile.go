package jobs

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"safari_downloader/conf"
	"strings"
)

//CreateDataFile create formated files with neccessary data from url
func CreateDataFile(config *conf.Config) error {
	//--------->get url
	res, err := http.Get(config.Url)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	//--------> read all contents from url
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fullData := string(html)
	index := strings.Index(fullData, "Table of Contents")
	if index < 0 {
		return errors.New("Table of Contents not exists")
	} else if !strings.Contains(fullData, "href") {
		return errors.New("There is no documents exits in this link")
	}

	subData := fullData[index:]
	re, err := regexp.Compile(`<a.*>`)
	if err != nil {
		return err
	}

	dataFile, err := os.Create(config.DataFile)
	if err != nil {
		return err
	}
	defer dataFile.Close()

	result := re.FindAllString(subData, -1)
	for _, v := range result {
		//---------->extract url from the line
		urlreg, err := regexp.Compile(`href=\".*html\"`)
		if err != nil {
			return err
		}
		u := urlreg.FindString(v)
		dataFile.WriteString(u+"\n")
		//--------->extract heading from the line
		headreg , err := regexp.Compile(`>.*<`)
		if err != nil{
			return err
		}
		//-------->checking for empty heading
		h := headreg.FindString(v)
		if len(h)!= 0{
			dataFile.WriteString("head="+h[1:len(h)-1]+"\n")				
		}
	}
	return nil
}
