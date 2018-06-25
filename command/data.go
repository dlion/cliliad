package command

import (
	"fmt"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

type Data struct{}

func (d *Data) Help() string {
	return "Returns the number of the datas used"
}

func (d *Data) Run(args []string) int {
	result, err := startFunction()
	if err != nil {
		fmt.Printf("Error to extract the data value: %v\n", err)
		return -1
	}
	emoji.Printf(":globe_with_meridians: %s ", result["data"])
	return 0
}

func (d *Data) Synopsis() string {
	return d.Help()
}
