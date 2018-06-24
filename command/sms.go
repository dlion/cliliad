package command

import "fmt"

type Sms struct{}

func (h *Sms) Help() string {
	return "Returns the number of the SMS sent"
}

func (h *Sms) Run(args []string) int {
	result, err := Init()
	if err != nil {
	}
	fmt.Printf("✉️   %s\n", result["sms"])
	return 0
}

func (h *Sms) Synopsis() string {
	return h.Help()
}
