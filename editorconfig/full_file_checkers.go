package editorconfig

import (
	"regexp"
	"strings"
)

var fullFileCheckers = map[string]FullFileChecker{
	"end_of_line": CheckEndOfLineRule,
}

type FullFileChecker func(ruleValue string, fileContent string) *FullFileCheckResult

// @todo - add fixers to each instance of FullFileCheckResult.
type FullFileCheckResult struct {
	isOk           bool
	messageIfNotOk string
}

var lfRegexp = regexp.MustCompile(`\n`)
var crRegexp = regexp.MustCompile(`\r`)
var crlfRegexp = regexp.MustCompile(`\r\n`)

func CheckEndOfLineRule(ruleValue string, fileContent string) *FullFileCheckResult {
	// Valid rules values are "lf", "cr", or "crlf". The values are case insensitive.
	ruleValueLowercase := strings.ToLower(ruleValue)

	if ruleValueLowercase != "lf" && ruleValueLowercase != "cr" && ruleValueLowercase != "crlf" {
		return &FullFileCheckResult{isOk: false, messageIfNotOk: "end_of_line value is not valid: " + ruleValue}
	}

	if ruleValueLowercase == "lf" {
		if crlfRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should use LF for new lines but contains CRLF"}
		}
		if crRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should use LF for new lines but contains CR"}
		}
	}

	if ruleValueLowercase == "cr" {
		if crlfRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should use CR for new lines but contains CRLF"}
		}
		if lfRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should use CR for new lines but contains LF"}
		}
	}

	if ruleValueLowercase == "crlf" {
		fileContent := crlfRegexp.ReplaceAllString(fileContent, "")
		if lfRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should use CRLF for new lines but contains LF"}
		}
		if crRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should use CRLF for new lines but contains CR"}
		}
	}

	return &FullFileCheckResult{isOk: true}
}
