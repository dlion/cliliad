package command

import (
	emoji "gopkg.in/kyokomi/emoji.v1"
)

type Mms struct{}

func (m Mms) Run(result map[string]string) {
	emoji.Printf(":outbox_tray: %s ", result["mms"])
}
