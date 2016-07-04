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

/**
 * This must be called before FixEndOfLineRule so the \n added will be converted to whatever the
 * 'end_of_line' rule dictates.
 */
func FixInsertFinalNewLineRule(ruleValue string, fileContent string) string {
	ruleValueLowercase := strings.ToLower(ruleValue)

	if ruleValueLowercase == "true" && !endsWithFinalNewLineRegexp.MatchString(fileContent) {
		return fileContent + "\n"
	}

	if ruleValueLowercase == "false" {
		for endsWithFinalNewLineRegexp.MatchString(fileContent) {
			fileContent = endsWithFinalNewLineRegexp.ReplaceAllString(fileContent, "")
		}
		return fileContent
	}

	return fileContent
}
