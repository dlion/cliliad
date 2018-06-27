package command

import (
	emoji "gopkg.in/kyokomi/emoji.v1"
)

type Data struct{}

func (d Data) Run(result map[string]string) {
	emoji.Printf(":globe_with_meridians: %s ", result["data"])
}
