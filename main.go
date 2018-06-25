package main

import (
	"fmt"
	"os"

	"github.com/dlion/cliliad/command"
	"github.com/mitchellh/cli"
)

func main() {
	app := cli.NewCLI("Cliliad", "1.0.0")
	app.Args = os.Args[1:]
	app.Commands = map[string]cli.CommandFactory{
		"sms": func() (cli.Command, error) {
			return &command.Sms{}, nil
		},
		"mms": func() (cli.Command, error) {
			return &command.Mms{}, nil
		},
		"calls": func() (cli.Command, error) {
			return &command.Calls{}, nil
		},
		"data": func() (cli.Command, error) {
			return &command.Data{}, nil
		},
		"all": func() (cli.Command, error) {
			return &command.All{}, nil
		},
	}

	status, err := app.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(status)
}
