# Safari Online Downloader

## Requirements

   1. [golang](https://golang.org/)
   2. [latest version of youtube-dl](https://github.com/rg3/youtube-dl) [prefer to use version:`2017.10.12`]

## Running Locally
    ```
    $ go get github.com/Kutty1995/safari_downloader
    $ cd $GOPATH/src/github.com/kutty1995/safari_downloader
    $ go run main.go
    ```
    

## Configure
   ###### maternity input arguments. 
   ```
       -l        url           url point to the safaribooksonline tutorial
       -u        username      safaribooksonline username
       -p        password      safaribooksonline  password
   ```
   ###### optional input arguments. 
   ```
       -d        directory     where you need to store downloaded file
   ```
   ##### default downloaded path : _$HOME/Documents/safari_
