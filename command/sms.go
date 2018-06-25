package command

import (
	"fmt"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

type Sms struct{}

func (s *Sms) Help() string {
	return "Returns the number of the SMS sent"
}

func (s *Sms) Run(args []string) int {
	result, err := startFunction()
	if err != nil {
		fmt.Printf("Error to extract the sms value: %v\n", err)
		return -1
	}
	emoji.Printf(":envelope: %s ", result["sms"])
	return 0
}

func (s *Sms) Synopsis() string {
	return s.Help()
}
