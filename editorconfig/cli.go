package editorconfig

import (
	"fmt"
	"github.com/codegangsta/cli"
	"strconv"
	"strings"
)

func LsCommand(c *cli.Context) error {
	files, err := FindSourceFiles(c.Args())
	if err != nil {
		return err
	}
	nbFiles := strconv.Itoa(len(files))
	fmt.Println(nbFiles + " source files are matched by the paths you gave " + "(" + strings.Join(c.Args(), ", ") + ")")
	for _, file := range files {
		fmt.Println(file)
	}
	fmt.Println()

	configs := FindConfigFiles(files)
	nbConfig := strconv.Itoa(len(configs))
	fmt.Println(nbConfig + " .editorconfig files were found that apply to your files")
	for _, c := range configs {
		if c.IsRoot() {
			fmt.Println("A root .editorconfig is in " + c.Path)
		} else {
			fmt.Println("A non-root .editorconfig is in " + c.Path)
		}
	}

	fmt.Println()

	return nil
}

func RulesCommand(c *cli.Context) error {
	path := c.Args()[0]
	configs := FindConfigFiles([]string{path})
	rules := GetRulesToApplyToSourcePath(path, configs)

	fmt.Printf("%d rules match the path %s\n", len(rules), path)
	for key, value := range rules {
		fmt.Println(key + " = " + value)
	}

	return nil
}

func CreateCliApp() *cli.App {
	app := cli.NewApp()

	app.Name = "editorconfig-cli"

	app.Usage = "Validate and fix files based on the rules in your .editorconfig file"

	app.Commands = []cli.Command{
		{
			Name:   "ls",
			Usage:  "List files that will be matched by the arguments you give",
			Action: LsCommand,
		},
		{
			Name:   "rules",
			Usage:  "List rules that match a given file",
			Action: RulesCommand,
		},
	}

	return app
}
