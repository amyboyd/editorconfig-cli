package editorconfig

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"strconv"
)

func CheckCommand(c *cli.Context) error {
	files, err := FindSourceFiles(c.Args())
	if err != nil {
		return err
	}

	configs := FindConfigFiles(files)

	for _, f := range files {
		fh, err := os.Open(f)
		if err != nil {
			ExitBecauseOfInternalError("Could not read file: " + f)
		}

		// @todo - add "full file checkers" for end_of_line, insert_final_newline and charset.

		scanner := bufio.NewScanner(fh)
		rules := GetRulesToApplyToSourcePath(f, configs)

		lineNo := 0
		for scanner.Scan() {
			line := scanner.Text()
			lineNo++

			for ruleName, ruleValue := range rules {
				if lineCheckers[ruleName] == nil {
					continue
				}

				result := lineCheckers[ruleName](ruleValue, line)
				if !result.isOk {
					fmt.Println(f + ": line " + strconv.Itoa(lineNo) + ": " + ruleName + ": " + result.messageIfNotOk)
					// Don't show more than 1 error per line.
					break
				}
			}
		}
	}

	return nil
}
