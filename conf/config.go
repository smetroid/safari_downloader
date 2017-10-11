package conf

import (
	"encoding/json"
	"log"
	"os"
)

//Config respresent credentials struct in our application
type Config struct {
	URL         string `json:"url"`
	Prefix      string `json:"prefix"`
	User        string `json:"username"`
	Pass        string `json:"password"`
	File        string `json:"errlog"`
	DataFile    string `json:"datafile"`
	Destination string `json:"destination"`
	Logger      *log.Logger
}

//ReadConfig reads configuration file
func ReadConfig() Config {

	//----------open configuration file
	file, err := os.Open("conf/config.json")
	if err != nil {
		log.Println(err.Error())
		log.Fatal("error occured : reading configuration file")
	}

	//---------decode
	conf := Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal("error occured : decoding section of configuration file")
	}

	//-----------open error log file with proper permission
	errlog, err := os.OpenFile(conf.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file : ", err.Error())
	}
	conf.Logger = log.New(errlog, "ERROR :", log.Ldate|log.Ltime|log.Lshortfile)
	return conf
}
