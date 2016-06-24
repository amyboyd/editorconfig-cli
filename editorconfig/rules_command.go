package editorconfig

import (
	"fmt"
	"github.com/codegangsta/cli"
)

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
