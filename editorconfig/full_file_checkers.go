package editorconfig

import (
	"regexp"
	"strings"
)

var fullFileCheckers = map[string]FullFileChecker{
	"end_of_line":          CheckEndOfLineRule,
	"insert_final_newline": CheckInsertFinalNewLineRule,
	// @todo - add checker for charset.
	// "charset":              CheckCharsetRule,
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

var endsWithFinalNewLineRegexp = regexp.MustCompile(`(\n|\r|\r\n)$`)

func CheckInsertFinalNewLineRule(ruleValue string, fileContent string) *FullFileCheckResult {
	// Valid rules values are "true" or "false". The values are case insensitive.
	ruleValueLowercase := strings.ToLower(ruleValue)

	if ruleValueLowercase != "true" && ruleValueLowercase != "false" {
		return &FullFileCheckResult{isOk: false, messageIfNotOk: "insert_final_new_line value should be true or false, is: " + ruleValue}
	}

	if len(fileContent) == 0 {
		return &FullFileCheckResult{isOk: true}
	}

	if ruleValueLowercase == "true" {
		if endsWithFinalNewLineRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: true}
		} else {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should end with an empty line but it does not"}
		}
	}

	if ruleValueLowercase == "false" {
		if !endsWithFinalNewLineRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: true}
		} else {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should not end with an empty line but it does"}
		}
	}

	return &FullFileCheckResult{isOk: false, messageIfNotOk: "unexpected condition"}
}
