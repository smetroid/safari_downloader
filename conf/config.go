package conf

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/user"
)

//Config respresent credentials struct in our application
type Config struct {
	URL         string `json:"url"`
	Prefix      string `json:"prefix"`
	User        string `json:"username"`
	Pass        string `json:"password"`
	File        string `json:"errlog"`
	DataFile    string `json:"datafile"`
	DType       string `json:"document_type"`
	Destination string `json:"destination"`
	Logger      *log.Logger
}

//ReadConfig reads configuration file
func ReadConfig(url *string, username *string, pass *string, dest *string, dType *string) (Config, error) {

	conf := Config{}
	//----------open configuration file
	file, err := os.Open("conf/config.json")
	if err != nil {
		return conf, err
	}

	//---------decode
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		return conf, errors.New("Can't read configuration file")
	}

	//-------credentials
	conf.URL = *url
	conf.User = *username
	conf.Pass = *pass
	conf.DType = *dType

	//-------downloaded location
	if *dest == "" {
		usr, err := user.Current()
		if err != nil {
			return conf, err
		}
		conf.Destination = usr.HomeDir + "/Documents/safari"
	} else {
		conf.Destination = *dest
	}

	//----------- error log
	errlog, err := os.OpenFile(conf.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return conf, errors.New("Failed to open log file")
	}

	conf.Logger = log.New(errlog, "ERROR :", log.Ldate|log.Ltime|log.Lshortfile)
	return conf, nil
}
