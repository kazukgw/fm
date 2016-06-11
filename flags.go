package main

import (
	"github.com/codegangsta/cli"
)

var flagFm = cli.StringSliceFlag{
	Name:  "frontmatter, fm",
	Value: &cli.StringSlice{},
	Usage: "-fm key:value [-fm key:value -fm...]",
}

var flagFmJson = cli.StringSliceFlag{
	Name:  "fm-json",
	Value: &cli.StringSlice{},
	Usage: `-fm-json '{key:["value", "value"]}'`,
}

var flagFilter = cli.StringFlag{
	Name:  "filter, f",
	Value: "",
	Usage: "-f ",
}

var flagIndexesPath = cli.StringFlag{
	Name:  "indexes-path",
	Value: "",
	Usage: "indexes path",
}

var flagSrc = cli.StringFlag{
	Name:  "src, s",
	Value: "",
	Usage: "source",
}
