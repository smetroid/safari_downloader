package jobs

import (
	"errors"
	"github.com/smetroid/safari_downloader/conf"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//CreateDataFile create formated files with necessary data from url
func CreateDataFile(config *conf.Config) error {
	//---------get url
	res, err := http.Get(config.URL)
	if err != nil {
		return errors.New("url error : unsupported url")
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
		return errors.New("Table of Contents not exists")
	} else if !strings.Contains(fullData, "href") {
		return errors.New("There is no video documents exits in this link")
	}

	//--------- extract all documents links form table of contents
	subData := fullData[index:]
	re, err := regexp.Compile(`<a.*>`)
	if err != nil {
		return err
	}
	//--------- remove file if exist
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
		//---------extract heading from the line
		headreg, err := regexp.Compile(`>.*<`)
		if err != nil {
			return err
		}
		//--------checking for empty heading
		head := headreg.FindString(v)
		if len(head) != 0 {
			//----------folder name
			matchSession, _ := regexp.MatchString(`^>(Lesson|Chapter|SECTION)\s*[0-9]*\s*:`, head)
			if matchSession {
				dataFile.WriteString("\n\nf=" + head[1:len(head)-1] + "\n")
			}
			//---------file name
			matchHead, _ := regexp.MatchString("<span>", head)
			if !matchHead {
				dataFile.WriteString("h=" + head[1:len(head)-1] + "\n")
			}
		}
		//----------extract url from the line
		urlreg, err := regexp.Compile(`href=\".*html\"`)
		if err != nil {
			return err
		}
		url := urlreg.FindString(v)
		//-----------remove href
		finalurl := strings.TrimLeft(strings.TrimLeft(url, "href="), " ")
		if len(finalurl) != 0 {
			dataFile.WriteString("l=" + finalurl + "\n")
		}

	}
	return nil
}
