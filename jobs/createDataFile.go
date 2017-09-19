package jobs

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"fmt"
	"regexp"
	"safari_downloader/conf"
	"strings"
	"strconv"
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
	for i, v := range result {
		urlreg, err := regexp.Compile(`href=\".*html\"`)
		if err != nil {
			return err
		}
		u := urlreg.FindString(v)
		dataFile.WriteString(u+strconv.Itoa(i)+"\n")
		
		headreg , err := regexp.Compile(`>.*:.*<`)
		if err != nil{
			return err
		}
		h := headreg.FindString(v)
		dataFile.WriteString(h+strconv.Itoa(i)+"\n")		
		
		fmt.Println(h,i)
		
	}
	return nil
}
