package jobs

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/smetroid/safari_downloader/conf"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

const extension = ".mp4"

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
	location := config.Destination
	var file string

	//----------read lines
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		line := scanner.Text()
		//----------file
		head := strings.HasPrefix(line, "h=")
		if head {
			count(c)
			file = strings.TrimLeft(line, "h=")
			info(" : " + file)
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
			_, err := os.Stat(location + "/" + fmt.Sprintf("%02d", c) + "-" + file + extension)
			if err != nil {
				url := config.Prefix + strings.TrimRight(strings.TrimLeft(line, "l=\""), "\"")
				//----------download
				err = exec.Command("youtube-dl", "-o", location+"/"+fmt.Sprintf("%02d", c)+"-"+file+extension, "-u", config.User, "-p", config.Pass, url).Run()
				if err != nil {
					return err
				}
			}
			c++
		}
	}
	//--------error check for reading file content
	if err := scanner.Err(); err != nil {
		return errors.New("read result file failed")
	}
	return nil
}
