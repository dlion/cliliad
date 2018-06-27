package command

import (
	emoji "gopkg.in/kyokomi/emoji.v1"
)

type Calls struct{}

func (c Calls) Run(result map[string]string) {
	emoji.Printf(":telephone_receiver: %s ", result["calls"])
}
