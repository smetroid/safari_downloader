package jobs

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"safari_downloader/conf"
	"strings"

	"github.com/fatih/color"
)

//DownloadFiles handling downloading of documents
func DownloadFiles(config *conf.Config) error {

	//----------display msg funcs
	info := color.New(color.Bold, color.FgHiMagenta).PrintlnFunc()
	msg := color.New(color.Bold, color.FgHiGreen).PrintlnFunc()
	count := color.New(color.Bold, color.FgHiWhite).PrintFunc()

	//-----------open result file
	f, err := os.Open(config.DataFile)
	if err != nil {
		return err
	}
	defer f.Close()

	//-----------create destination folder if not exists
	if _, err := os.Stat(config.Destination); os.IsNotExist(err) {
		err = os.Mkdir(config.Destination, 0775)
		if err != nil {
			return err
		}
	}

	var c int64
	var file string

	//----------read lines
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		location := config.Destination
		//----------file
		head := strings.HasPrefix(line, "h=")
		if head {
			count(c)
			file = strings.TrimLeft(line, "h=")
			info(" : " + file)
			c++
		}
		//---------folder
		folder := strings.HasPrefix(line, "f=")
		if folder {
			c = 0
			f := strings.TrimLeft(line, "f=")
			msg(f)
			//----------create folder
			location = config.Destination + "/" + strings.Replace(f, "/", "-", 5)
			if _, err := os.Stat(location); os.IsNotExist(err) {
				err = os.Mkdir(location, 0775)
				if err != nil {
					return err
				}
			}
		}
		//--------link
		link := strings.HasPrefix(line, "l=")
		if link {
			url := config.Prefix + strings.TrimRight(strings.TrimLeft(line, "l=\""), "\"")
			err = videoDLWorker(location, file, url)
			if err != nil {
				return err
			}
		}
	}
	//--------error check for reading file content
	if err := scanner.Err(); err != nil {
		return errors.New("read result file failed")
	}
	return nil
}

func videoDLWorker(dest string, file string, target string) error {
	resp, err := http.Get(target)
	if err != nil {
		return err
	}
	fmt.Println("video info :", resp.ContentLength)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("non 200 status code received")
	}
	out, err := os.Create(dest + "/" + file + ".mp4")
	if err != nil {
		return err
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
