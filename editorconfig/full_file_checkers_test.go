package editorconfig

import (
	"testing"
)

func ExpectFilePass(file string, ruleValue string, fileChecker FullFileChecker, t *testing.T) {
	result := fileChecker(ruleValue, file)
	if !result.isOk {
		t.Error("Expected file to pass, but it failed: \"" + file + "\" for rule value \"" + ruleValue + "\", had error message: " + result.messageIfNotOk)
	}
}

func ExpectFileFail(file string, ruleValue string, fileChecker FullFileChecker, t *testing.T, expectedError string) {
	result := fileChecker(ruleValue, file)
	if result.isOk {
		t.Error("Expected file to fail, but it passed: \"" + file + "\" for rule value \"" + ruleValue + "\"")
		return
	}

	if !result.isOk && result.messageIfNotOk != expectedError {
		t.Error("File \"" + file + "\" failed with error message \"" + result.messageIfNotOk + "\" but had expected \"" + expectedError + "\"")
		return
	}
}

func TestCheckEndOfLineRule(t *testing.T) {
	f := CheckEndOfLineRule

	ExpectFilePass("", "lf", f, t)
	ExpectFilePass("", "cr", f, t)
	ExpectFilePass("", "crlf", f, t)

	ExpectFilePass("Aardvark\nBunny\n", "lf", f, t)
	ExpectFilePass("Aardvark\rBunny\r", "cr", f, t)
	ExpectFilePass("Aardvark\r\nBunny\r\n", "crlf", f, t)

	ExpectFileFail("Aardvark\nBunny\n", "cr", f, t, "should use CR for new lines but contains LF")
	ExpectFileFail("Aardvark\rBunny\r", "crlf", f, t, "should use CRLF for new lines but contains CR")
	ExpectFileFail("Aardvark\r\nBunny\r\n", "lf", f, t, "should use LF for new lines but contains CRLF")
}

func TestCheckInsertFinalNewLineRule(t *testing.T) {
	f := CheckInsertFinalNewLineRule

	// Empty files should never error in this rule.
	ExpectFilePass("", "true", f, t)
	ExpectFilePass("", "false", f, t)

	ExpectFilePass("\n", "true", f, t)
	ExpectFilePass("\r", "true", f, t)
	ExpectFilePass("a\n", "true", f, t)
	ExpectFilePass("a\r", "true", f, t)
	ExpectFilePass("a\r\n", "true", f, t)
	ExpectFilePass("a\r\n", "true", f, t)
	ExpectFilePass("a\nb\n", "true", f, t)
	ExpectFileFail("a", "true", f, t, "should end with an empty line but it does not")
	ExpectFileFail("a\nb", "true", f, t, "should end with an empty line but it does not")
	ExpectFileFail("a\rb", "true", f, t, "should end with an empty line but it does not")

	ExpectFilePass("a", "false", f, t)
	ExpectFilePass("a\nb", "false", f, t)
	ExpectFileFail("\n", "false", f, t, "should not end with an empty line but it does")
	ExpectFileFail("\r", "false", f, t, "should not end with an empty line but it does")
	ExpectFileFail("\r\n", "false", f, t, "should not end with an empty line but it does")
	ExpectFileFail("a\r\n", "false", f, t, "should not end with an empty line but it does")

	ExpectFileFail("anything", "rabbits", f, t, "insert_final_new_line value should be true or false, is: rabbits")
}
