package command

import (
	emoji "gopkg.in/kyokomi/emoji.v1"
)

type Sms struct{}

func (s Sms) Run(result map[string]string) {
	emoji.Printf(":envelope: %s ", result["sms"])
}
