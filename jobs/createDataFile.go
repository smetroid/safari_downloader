package jobs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"safari_downloader/conf"
	"strings"
)

func CreateDataFile(config *conf.Config) error {
	res, err := http.Get(config.Url)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fullData := string(html)
	if !strings.Contains(fullData, "href") {
		return errors.New("There is no documents exits in this link")
	}
	index := strings.Index(fullData, "Table of Contents")
	subData := fullData[index:]

	re, err := regexp.Compile(`<a href=\"/(\w+|\d+|\s)/(\w+|\d+|\s)/(\w+|\d+|\s)/\d+`)
	if err != nil {
		return err
	}

	fmt.Println("regular expression:", re)
	result := re.FindAllStringSubmatch(subData, -1)
	for i, v := range result {
		fmt.Println("url : ", v, i)
	}
	return nil
}
