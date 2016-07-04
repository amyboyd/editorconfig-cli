package editorconfig

import (
	"regexp"
	"strings"
)

type FullFileFixer func(ruleValue string, fileContent string) string

func FixEndOfLineRule(ruleValue string, fileContent string) string {
	ruleValueLowercase := strings.ToLower(ruleValue)

	if ruleValueLowercase == "lf" {
		fileContent = crlfRegexp.ReplaceAllString(fileContent, "\n")
		fileContent = crRegexp.ReplaceAllString(fileContent, "\n")
		return fileContent
	}

	if ruleValueLowercase == "cr" {
		fileContent = crlfRegexp.ReplaceAllString(fileContent, "\r")
		fileContent = lfRegexp.ReplaceAllString(fileContent, "\r")
		return fileContent
	}

	if ruleValueLowercase == "crlf" {
		fileContent = regexp.MustCompile("(\r\n|\r|\n)").ReplaceAllString(fileContent, "\r\n")
		return fileContent
	}

	return fileContent
}
