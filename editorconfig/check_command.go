package editorconfig

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"strconv"
	"strings"
)

func CheckCommand(c *cli.Context) error {
	files, err := FindSourceFiles(c.Args())
	if err != nil {
		return err
	}

	if len(files) == 0 {
		ExitBecauseOfInternalError("No files to check in " + strings.Join(c.Args(), ", "))
	}

	configs := FindConfigFiles(files)

	hasError := false

	for _, f := range files {
		rules := GetRulesToApplyToSourcePath(f, configs)
		if len(rules) == 0 {
			continue
		}

		fileContent := MustGetFileAsString(f)

		// Run full-file checkers.
		for ruleName, ruleValue := range rules {
			if fullFileChecker, ok := fullFileCheckers[ruleName]; ok {
				result := fullFileChecker(ruleValue, fileContent)
				if !result.isOk {
					hasError = true
					fmt.Println(f + ": " + ruleName + ": " + result.messageIfNotOk)
				}
			}
		}

		// Run line checkers.
		lines := SplitIntoLines(fileContent)
		lineNo := 1
		for _, line := range lines {
			for ruleName, ruleValue := range rules {
				if lineChecker, ok := lineCheckers[ruleName]; ok {
					result := lineChecker(ruleValue, line)
					if !result.isOk {
						fmt.Println(f + ": line " + strconv.Itoa(lineNo) + ": " + ruleName + ": " + result.messageIfNotOk)
						hasError = true
						// Don't show more than 1 error per line.
						break
					}
				}
			}
			lineNo++
		}
	}

	if hasError {
		os.Exit(1)
	}

	return nil
}
