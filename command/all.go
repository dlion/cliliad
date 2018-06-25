package command

import "github.com/mitchellh/cli"

type All struct{}

func (a *All) Help() string {
	return "Returns all stats"
}

func (a *All) Run(args []string) int {
	var all = []cli.Command{&Sms{}, &Calls{}, &Data{}, &Mms{}}
	for _, command := range all {
		command.Run([]string{})
	}
	return 0
}

func (a *All) Synopsis() string {
	return a.Help()
}
