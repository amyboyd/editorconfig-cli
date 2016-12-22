package editorconfig

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"strconv"
	"strings"
)

func FixCommand(c *cli.Context) error {
	files, err := FindSourceFiles(c.Args())
	if err != nil {
		return err
	}

	if len(files) == 0 {
		ExitBecauseOfInternalError("No files to check in " + strings.Join(c.Args(), ", "))
	}

	configs := FindConfigFiles(files)

	for _, f := range files {
		rules := GetRulesToApplyToSourcePath(f, configs)
		if len(rules) == 0 {
			continue
		}

		fileContent := MustGetFileAsString(f)
		hasChanged := false

		// Run full-file checkers and fixers.
		for ruleName, ruleValue := range rules {
			if fullFileChecker, ok := fullFileCheckers[ruleName]; ok {
				result := fullFileChecker(ruleValue, fileContent)
				if result.isOk {
					continue
				}

				if result.fixer != nil {
					fileContent = result.fixer(ruleValue, fileContent)
					hasChanged = true
					fmt.Println(f + ": " + ruleName + ": fixed")
				} else {
					fmt.Println(f + ": " + ruleName + ": cannot fix automatically")
				}
			}
		}

		// Run line checkers and fixers.
		lines := SplitIntoLines(fileContent)
		lineNo := 1
		for _, line := range lines {
			for ruleName, ruleValue := range rules {
				if lineChecker, ok := lineCheckers[ruleName]; ok {
					result := lineChecker(ruleValue, line)
					if !result.isOk {
						fmt.Println(f + ": line " + strconv.Itoa(lineNo) + ": " + ruleName + ": " + result.messageIfNotOk)
						// Don't show more than 1 error per line.
						break
					}
				}
			}
			lineNo++
		}

		if hasChanged {
			fileHandler, err := os.Open(f)
			if err != nil {
				fmt.Println("Could not write to " + f)
			}
			fileHandler.WriteString(fileContent)
			fmt.Println("Wrote to " + f)
		}
	}

	return nil
}
