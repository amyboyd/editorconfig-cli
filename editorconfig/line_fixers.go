package editorconfig

import (
	"strconv"
	"strings"
)

type LineFixer func(ruleValue string, line string) string

func FixTabIndentationToSpaces(ruleValueNumberOfSpaces string, line string) string {
	numberOfSpaces, _ := strconv.Atoi(ruleValueNumberOfSpaces)

	line = indentedWithTabsRegexp.ReplaceAllStringFunc(line, func(tabs string) string {
		return strings.Repeat(" ", len(tabs)*numberOfSpaces)
	})

	return line
}

func FixMixedIndentationToSpaces(ruleValueNumberOfSpaces string, line string) string {
	numberOfSpaces, _ := strconv.Atoi(ruleValueNumberOfSpaces)

	for indentedWithMixedTabsAndSpacesRegexp.MatchString(line) {
		line = indentedWithMixedTabsAndSpacesRegexp.ReplaceAllStringFunc(line, func(tabsAndSpaces string) string {
			tabs := strings.Replace(tabsAndSpaces, " ", "", -1)
			spaces := strings.Replace(tabsAndSpaces, "\t", "", -1)

			return strings.Repeat(" ", (len(tabs)*numberOfSpaces)+len(spaces))
		})
	}

	return line
}

func FixUndividableIndentationToNearestSpacesAmount(ruleValueNumberOfSpaces string, line string) string {
	numberOfSpaces, _ := strconv.Atoi(ruleValueNumberOfSpaces)
	if numberOfSpaces < 1 {
		ExitBecauseOfInternalError("Number of spaces must be integer greater than 0, is: " + ruleValueNumberOfSpaces)
	}

	if GetNumberOfLeftSpaces(line) == 0 {
		return line
	}

	for true {
		leftSpaces := GetNumberOfLeftSpaces(line)
		if leftSpaces%numberOfSpaces != 0 {
			line = " " + line
		} else {
			break
		}
	}

	return line
}

func FixTrimTrailingWhitespaceRule(ruleValue string, line string) string {
	if strings.ToLower(ruleValue) != "true" {
		return line
	}

	return endsWithTabsAndSpacesRegexp.ReplaceAllString(line, "")
}
