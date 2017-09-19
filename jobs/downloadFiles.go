package jobs

import(
	"safari_downloader/conf"
	"os"
	"bufio"
	"net/http"
) 

func DownloadFiles(config *conf.Config)error{
	//------------>Create destination folder
	err := os.Mkdir(config.Destination, 0700)
	if err!=nil{
		return err
	}

	file,err := os.Open(config.DataFile)
	if err!=nil{
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url :=scanner.Text()
		var head string
		if scanner.Scan(){
			head = scanner.Text()		
		}
		if len(head)!=0{
			_, err := os.Stat(config.Destination+"/"+head[5:]); 
			if err!=nil{
				err:=os.Mkdir(config.Destination+"/"+head[5:], 0700)
				if err!=nil{
					return err
				}
			}
		}
		if len(url)!=0{
			response, err := http.Get(config.Prefix+url)
			if err != nil {
				return err
			}
			defer response.Body.Close()
			// _, err := io.Copy(out, response.Body)
			// if err != nil {
			// 	return err
			// }
		}
    }
	return nil
}