package command

import (
	"fmt"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

type Calls struct{}

func (c *Calls) Help() string {
	return "Returns the number of time spent for calls"
}

func (c *Calls) Run(args []string) int {
	result, err := startFunction()
	if err != nil {
		fmt.Printf("Error to extract the calls value: %v\n", err)
		return -1
	}
	emoji.Printf(":telephone_receiver: %s ", result["calls"])
	return 0
}

func (c *Calls) Synopsis() string {
	return c.Help()
}
