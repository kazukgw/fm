package main

import (
	"os"
	"reflect"

	"github.com/codegangsta/cli"
	"github.com/kazukgw/coa"
)

func main() {
	app := cli.NewApp()
	app.Name = "fm"
	app.Usage = "fm [command]"
	app.Commands = []cli.Command{
		{
			Name:   "build",
			Usage:  "fm build",
			Action: build(buildIndexes{}),
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "fm list [key]",
			Action:  build(showList{}),
		},
	}
	app.Run(os.Args)
}

func build(zeroAG interface{}) func(*cli.Context) {
	actionType := reflect.TypeOf(zeroAG)
	if _, ok := reflect.New(actionType).Interface().(coa.ActionGroup); !ok {
		panic(actionType.String() + " dose not implement coa.ActionGroup interface")
	}

	return func(c *cli.Context) {
		ag := reflect.New(actionType).Interface().(coa.ActionGroup)
		coa.Exec(ag, &Ctx{c, ag})
	}
}

type Ctx struct {
	*cli.Context
	ag coa.ActionGroup
}

func (c *Ctx) ActionGroup() coa.ActionGroup {
	return c.ag
}
