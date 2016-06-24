package editorconfig

import (
	"testing"
)

func ExpectPass(line string, ruleValue string, lineChecker LineChecker, t *testing.T) {
	result := lineChecker(ruleValue, line)
	if !result.isOk {
		t.Error("Expected line to pass, but it failed: \"" + line + "\" for rule value \"" + ruleValue + "\", had error message: " + result.messageIfNotOk)
	}
}

func ExpectFail(line string, ruleValue string, lineChecker LineChecker, t *testing.T, expectedError string) {
	result := lineChecker(ruleValue, line)
	if result.isOk {
		t.Error("Expected line to fail, but it passed: \"" + line + "\" for rule value \"" + ruleValue + "\"")
		return
	}

	if !result.isOk && result.messageIfNotOk != expectedError {
		t.Error("Line \"" + line + "\" failed with error message \"" + result.messageIfNotOk + "\" but had expected \"" + expectedError + "\"")
		return
	}
}

func TestCheckIndentStyleRule(t *testing.T) {
	f := CheckIndentStyleRule

	ExpectPass("", "space", f, t)
	ExpectPass("line", "space", f, t)
	ExpectPass(" line", "space", f, t)
	ExpectPass(" ", "space", f, t)
	ExpectPass("  line", "space", f, t)
	ExpectFail(" \tline", "space", f, t, "indented with mix of tabs and spaces instead of just tabs")
	ExpectFail("\tline", "space", f, t, "starts with tab instead of space")
	ExpectFail("\t line", "space", f, t, "starts with tab instead of space")

	ExpectPass("", "tab", f, t)
	ExpectPass("line", "tab", f, t)
	ExpectPass("\n", "tab", f, t)
	ExpectPass("\tline", "tab", f, t)
	ExpectPass("\t\tline", "tab", f, t)
	ExpectFail(" \tline", "tab", f, t, "starts with space instead of tab")
	ExpectFail("\t line", "tab", f, t, "indented with mix of tabs and spaces instead of just tabs")
	ExpectFail("\t ", "tab", f, t, "indented with mix of tabs and spaces instead of just tabs")

	// Allow comments like /**\n\t *\n\t */
	ExpectPass("\t *line", "tab", f, t)

	ExpectFail("  line", "dinosaurs", f, t, "invalid value for indent_style: dinosaurs")
}

func TestCheckIndentSizeRule(t *testing.T) {
	f := CheckIndentSizeRule

	// 'indent_size=tab' can never fail this rule.
	ExpectPass("line", "tab", f, t)
	ExpectPass(" line", "tab", f, t)
	ExpectPass("  line", "tab", f, t)

	ExpectPass("", "2", f, t)
	ExpectPass("line", "2", f, t)
	ExpectPass("  line", "2", f, t)
	ExpectPass("    line", "2", f, t)
	ExpectPass("      line", "2", f, t)
	ExpectFail(" line", "2", f, t, "starts with 1 spaces which does not divide by 2")
	ExpectFail("   line", "2", f, t, "starts with 3 spaces which does not divide by 2")
	ExpectFail("     line", "2", f, t, "starts with 5 spaces which does not divide by 2")
	ExpectFail("\tline", "2", f, t, "should be indented with spaces but is indented with tabs")
	ExpectFail("\t\tline", "2", f, t, "should be indented with spaces but is indented with tabs")
	ExpectFail("  \tline", "2", f, t, "indented with mix of spaces and tabs instead of just spaces")

	// Allow comments like /**\n   *\n   */ (note the extra space before the * on the 2nd and 3rd lines)
	ExpectPass("   * line", "2", f, t)
	ExpectFail("   ^ line", "2", f, t, "starts with 3 spaces which does not divide by 2")

	ExpectPass("", "3", f, t)
	ExpectPass("line", "3", f, t)
	ExpectPass("      line", "3", f, t)
	ExpectFail("     line", "3", f, t, "starts with 5 spaces which does not divide by 3")

	ExpectFail(" anything", "0", f, t, "number of spaces must be 1 or more, is: 0")
	ExpectFail(" anything", "-1", f, t, "number of spaces must be 1 or more, is: -1")
	ExpectFail(" anything", "asd", f, t, "value is not an integer: asd")
}

func TestCheckTrimTrailingWhitespaceRule(t *testing.T) {
	f := CheckTrimTrailingWhitespaceRule

	// 'trim_trailing_whitespace=false' can never fail this rule.
	ExpectPass("line", "false", f, t)
	ExpectPass("line ", "false", f, t)
	ExpectPass("line\t", "false", f, t)
	ExpectPass("line \t \t \t", "false", f, t)

	ExpectPass("line", "true", f, t)
	ExpectPass("line line", "true", f, t)
	ExpectPass("line line", "true", f, t)

	ExpectFail("line line ", "true", f, t, "line has trailing whitespace")
	ExpectFail("line line\t", "true", f, t, "line has trailing whitespace")
	ExpectFail("line line   ", "true", f, t, "line has trailing whitespace")
	ExpectFail("line line   \t", "true", f, t, "line has trailing whitespace")
}
