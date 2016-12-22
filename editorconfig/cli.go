package editorconfig

import (
	"github.com/codegangsta/cli"
)

func CreateCliApp() *cli.App {
	app := cli.NewApp()

	app.Name = "editorconfig-cli"

	app.Usage = "Validate and fix files based on the rules in your .editorconfig file"

	app.Commands = []cli.Command{
		{
			Name:      "ls",
			Usage:     "List files that will be matched by the arguments you give",
			Action:    LsCommand,
			ArgsUsage: "[PATH1] [PATH2...]",
		},
		{
			Name:      "rules",
			Usage:     "List rules that match a given file",
			Action:    RulesCommand,
			ArgsUsage: "[PATH1] [PATH2...]",
		},
		{
			Name:      "check",
			Usage:     "Validate files",
			Action:    CheckCommand,
			ArgsUsage: "[PATH1] [PATH2...]",
		},
		{
			Name:      "fix",
			Usage:     "Fix invalid files",
			Action:    FixCommand,
			ArgsUsage: "[PATH1] [PATH2...]",
		},
	}

	return app
}
