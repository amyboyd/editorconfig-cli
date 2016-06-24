package editorconfig

import (
	"regexp"
	"strconv"
	"strings"
)

// tab_width, charset, end_of_line, insert_final_newline and root do not have any affect on our line
// checkers (they apply to the full files, not lines).

var lineCheckers = map[string]LineChecker{
	"indent_style":             CheckIndentStyleRule,
	"indent_size":              CheckIndentSizeRule,
	"trim_trailing_whitespace": CheckTrimTrailingWhitespaceRule,
}

type LineChecker func(ruleValue string, line string) *LineCheckResult

// @todo - add fixers to each instance of LineCheckResult.
type LineCheckResult struct {
	isOk           bool
	messageIfNotOk string
}

var hasIndentationRegexp = regexp.MustCompile(`^[\t ]`)
var hasNoIndentationRegexp = regexp.MustCompile(`^([^\t ]|$)`)
var indentedWithMixedTabsAndSpacesRegexp = regexp.MustCompile(`^(\t+ +| +\t+)`)
var indentedWithTabsRegexp = regexp.MustCompile(`^\t+`)
var indentedWithTabsThenCommentLineRegexp = regexp.MustCompile(`^\t+ \*`)
var indentedWithSpacesRegexp = regexp.MustCompile(`^ +`)

func HasIndentation(s string) bool {
	return hasIndentationRegexp.MatchString(s)
}

func HasNoIndentation(s string) bool {
	return hasNoIndentationRegexp.MatchString(s)
}

func IsIndentedWithMixedTabsAndSpaces(s string) bool {
	return indentedWithMixedTabsAndSpacesRegexp.MatchString(s)
}

func IsIndentedWithTabs(s string) bool {
	return indentedWithTabsRegexp.MatchString(s)
}

// This allows comments like /**\n\t *\n\t */
func IsIndentedWithTabsThenCommentLine(s string) bool {
	return indentedWithTabsThenCommentLineRegexp.MatchString(s)
}

func IsIndentedWithSpaces(s string) bool {
	return indentedWithSpacesRegexp.MatchString(s)
}

func CheckIndentStyleRule(ruleValue string, line string) *LineCheckResult {
	if HasNoIndentation(line) {
		return &LineCheckResult{isOk: true}
	}

	if strings.ToLower(ruleValue) == "tab" {
		if IsIndentedWithSpaces(line) {
			return &LineCheckResult{isOk: false, messageIfNotOk: "starts with space instead of tab"}
		} else if IsIndentedWithMixedTabsAndSpaces(line) && !IsIndentedWithTabsThenCommentLine(line) {
			return &LineCheckResult{isOk: false, messageIfNotOk: "indented with mix of tabs and spaces instead of just tabs"}
		} else {
			return &LineCheckResult{isOk: true}
		}
	}

	if strings.ToLower(ruleValue) == "space" {
		if IsIndentedWithTabs(line) {
			return &LineCheckResult{isOk: false, messageIfNotOk: "starts with tab instead of space"}
		} else if IsIndentedWithMixedTabsAndSpaces(line) {
			return &LineCheckResult{isOk: false, messageIfNotOk: "indented with mix of tabs and spaces instead of just tabs"}
		} else {
			return &LineCheckResult{isOk: true}
		}
	}

	return &LineCheckResult{isOk: false, messageIfNotOk: "invalid value for indent_style: " + ruleValue}
}

func CheckIndentSizeRule(ruleValue string, line string) *LineCheckResult {
	if ruleValue == "tab" {
		return &LineCheckResult{isOk: true}
	}

	if HasNoIndentation(line) {
		return &LineCheckResult{isOk: true}
	}

	ruleValueInt, err := strconv.Atoi(ruleValue)
	if err != nil {
		return &LineCheckResult{isOk: false, messageIfNotOk: "value is not an integer: " + ruleValue}
	}

	if ruleValueInt < 1 {
		return &LineCheckResult{isOk: false, messageIfNotOk: "number of spaces must be 1 or more, is: " + ruleValue}
	}

	if strings.HasPrefix(line, "\t") {
		return &LineCheckResult{isOk: false, messageIfNotOk: "should be indented with spaces but is indented with tabs"}
	}

	// Indented with spaces. Ensure the number of spaces is divisible by the rule value, but also
	// allow an extra space followed by * to allow for comments like /**\n   *\n   */ (note the
	// extra space before the * on the 2nd and 3rd lines).
	trimmedLine := line
	for strings.HasPrefix(trimmedLine, strings.Repeat(" ", ruleValueInt)) {
		trimmedLine = (trimmedLine)[ruleValueInt:]
	}
	if strings.HasPrefix(trimmedLine, " *") {
		return &LineCheckResult{isOk: true}
	}
	if IsIndentedWithTabs(trimmedLine) {
		return &LineCheckResult{isOk: false, messageIfNotOk: "indented with mix of spaces and tabs instead of just spaces"}
	}
	if HasIndentation(trimmedLine) {
		leftSpaces := len(line) - len(strings.TrimLeft(line, " "))
		return &LineCheckResult{isOk: false, messageIfNotOk: "starts with " + strconv.Itoa(leftSpaces) + " spaces which does not divide by " + ruleValue}
	}

	return &LineCheckResult{isOk: true}
}

func CheckTrimTrailingWhitespaceRule(ruleValue string, line string) *LineCheckResult {
	if strings.ToLower(ruleValue) == "false" {
		return &LineCheckResult{isOk: true}
	}

	if strings.ToLower(ruleValue) != "true" {
		return &LineCheckResult{isOk: false, messageIfNotOk: "value must be true or false, but is: " + ruleValue}
	}

	trimmed := strings.TrimRight(line, " \t")
	if len(line) != len(trimmed) {
		return &LineCheckResult{isOk: false, messageIfNotOk: "line has trailing whitespace"}
	}

	return &LineCheckResult{isOk: true}
}
