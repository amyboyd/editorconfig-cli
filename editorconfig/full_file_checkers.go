package editorconfig

import (
	"github.com/saintfish/chardet"
	"strings"
)

var fullFileCheckers = map[string]FullFileChecker{
	"end_of_line":          CheckEndOfLineRule,
	"insert_final_newline": CheckInsertFinalNewLineRule,
	"charset":              CheckCharsetRule,
}

type FullFileChecker func(ruleValue string, fileContent string) *FullFileCheckResult

type FullFileCheckResult struct {
	isOk           bool
	messageIfNotOk string
	fixer          FullFileFixer
}

func CheckEndOfLineRule(ruleValue string, fileContent string) *FullFileCheckResult {
	// Valid rules values are "lf", "cr", or "crlf". The values are case insensitive.
	ruleValueLowercase := strings.ToLower(ruleValue)

	if ruleValueLowercase != "lf" && ruleValueLowercase != "cr" && ruleValueLowercase != "crlf" {
		return &FullFileCheckResult{isOk: false, messageIfNotOk: "end_of_line value is not valid: " + ruleValue}
	}

	if ruleValueLowercase == "lf" {
		if crlfRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should use LF for new lines but contains CRLF", fixer: FixEndOfLineRule}
		}
		if crRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should use LF for new lines but contains CR", fixer: FixEndOfLineRule}
		}
	}

	if ruleValueLowercase == "cr" {
		if crlfRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should use CR for new lines but contains CRLF", fixer: FixEndOfLineRule}
		}
		if lfRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should use CR for new lines but contains LF", fixer: FixEndOfLineRule}
		}
	}

	if ruleValueLowercase == "crlf" {
		fileContent := crlfRegexp.ReplaceAllString(fileContent, "")
		if lfRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should use CRLF for new lines but contains LF", fixer: FixEndOfLineRule}
		}
		if crRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should use CRLF for new lines but contains CR", fixer: FixEndOfLineRule}
		}
	}

	return &FullFileCheckResult{isOk: true}
}

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
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should end with an empty line but it does not", fixer: FixInsertFinalNewLineRule}
		}
	}

	if ruleValueLowercase == "false" {
		if !endsWithFinalNewLineRegexp.MatchString(fileContent) {
			return &FullFileCheckResult{isOk: true}
		} else {
			return &FullFileCheckResult{isOk: false, messageIfNotOk: "should not end with an empty line but it does", fixer: FixInsertFinalNewLineRule}
		}
	}

	return &FullFileCheckResult{isOk: false, messageIfNotOk: "unexpected condition"}
}

func CheckCharsetRule(ruleValue string, fileContent string) *FullFileCheckResult {
	// Valid rules values are "latin1", "utf-8", "utf-8-bom", "utf-16be" or "utf-16le".
	if ruleValue != "latin1" && ruleValue != "utf-8" && ruleValue != "utf-8-bom" && ruleValue != "utf-16be" && ruleValue != "utf-16le" {
		return &FullFileCheckResult{isOk: false, messageIfNotOk: "charset value is invalid: " + ruleValue}
	}

	detector := chardet.NewTextDetector()

	bestGuess, err := detector.DetectBest([]byte(fileContent))
	if err != nil {
		ExitBecauseOfInternalError(err.Error())
	}

	actual := strings.ToLower(bestGuess.Charset)

	if ruleValue == actual {
		return &FullFileCheckResult{isOk: true}
	}

	if ruleValue == "utf-8" && strings.HasPrefix(actual, "iso-8859-") {
		// iso-8859-* is a subset of utf-8.
		return &FullFileCheckResult{isOk: true}
	}

	return &FullFileCheckResult{isOk: false, messageIfNotOk: "expected " + ruleValue + " but is: " + actual}
}
