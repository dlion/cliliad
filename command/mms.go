package command

import (
	"fmt"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

type Mms struct{}

func (m *Mms) Help() string {
	return "Returns the number of the MMS sent"
}

func (m *Mms) Run(args []string) int {
	result, err := startFunction()
	if err != nil {
		fmt.Printf("Error to extract the mms value: %v\n", err)
		return -1
	}
	emoji.Printf(":outbox_tray: %s ", result["mms"])
	return 0
}

func (m *Mms) Synopsis() string {
	return m.Help()
}
