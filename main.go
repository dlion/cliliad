package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/dlion/cliliad/command"
	"github.com/dlion/cliliad/goiliad"
)

const logo = `
          oooo   o8o  oooo   o8o                  .o8  
          '888   '"'  '888   '"'                 "888  
 .ooooo.   888  oooo   888  oooo   .oooo.    .oooo888  
d88' '"Y8  888  '888   888  '888  'P  )88b  d88' '888  
888        888   888   888   888   .oP"888  888   888  
888   .o8  888   888   888   888  d8(  888  888   888  
'Y8bod8P' o888o o888o o888o o888o 'Y888""8o 'Y8bod88P" 
                                                       
                                                    `

func main() {
	color.Red(logo)

	sms := flag.Bool("sms", false, "Returns the number of the sms sent")
	mms := flag.Bool("mms", false, "Returns the number of the mms sent")
	calls := flag.Bool("calls", false, "Returns the time of the calls done")
	data := flag.Bool("data", false, "Returns the number of the data sent")

	flag.Parse()

	viper.SetConfigName(".cliliad")
	home, err := homedir.Dir()
	if err != nil {
		color.Red("Home dir not found")
		os.Exit(-1)
	}
	viper.AddConfigPath(home)

	if err = viper.ReadInConfig(); err != nil {
		color.Red("Fatal error reading the configuration file: %v \n", err)
		os.Exit(-1)
	}

	resultMap, err := startFunction(viper.GetString("user"), viper.GetString("password"))
	if err != nil {
		color.Cyan(err.Error())
		os.Exit(1)

	}

	nFlag := flag.NFlag()

	if *sms || nFlag == 0 {
		command.Sms{}.Run(resultMap)
	}

	if *mms || nFlag == 0 {
		command.Mms{}.Run(resultMap)
	}

	if *calls || nFlag == 0 {
		command.Calls{}.Run(resultMap)
	}

	if *data || nFlag == 0 {
		command.Data{}.Run(resultMap)
	}
}

func startFunction(u, p string) (map[string]string, error) {
	cookie, err := goiliad.GetInitialCookie()
	if err != nil {
		return nil, err
	}

	resultBody, err := login(u, p, cookie)
	if err != nil {
		return nil, err
	}

	resultMap, err := goiliad.PageScraper(resultBody)
	if err != nil {
		return nil, err
	}

	return resultMap, nil
}

func login(u, p string, c *http.Cookie) (string, error) {
	data := goiliad.CreateCredentials(u, p)

	req, err := goiliad.CreateRequest(data, c)
	if err != nil {
		return "Error on CreateRequest", err
	}

	rsp, err := goiliad.PerformRequest(req)
	if err != nil {
		return "Error on PerformRequest", err
	}

	return goiliad.ReadResponse(rsp)
}
