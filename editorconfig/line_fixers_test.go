package editorconfig

import (
	"testing"
)

func TestFixTabIndentationToSpaces(t *testing.T) {
	var result string

	result = FixTabIndentationToSpaces("4", "\t\thello world")
	if result != "        hello world" {
		t.Error("Unexpected result: " + result)
	}

	result = FixTabIndentationToSpaces("3", "\thello world")
	if result != "   hello world" {
		t.Error("Unexpected result: " + result)
	}

	result = FixTabIndentationToSpaces("2", "\t\t\thello world")
	if result != "      hello world" {
		t.Error("Unexpected result: " + result)
	}
}

func TestFixMixedIndentationToSpaces(t *testing.T) {
	var result string

	result = FixMixedIndentationToSpaces("2", "\t  \t  hello worl d")
	if result != "        hello worl d" {
		t.Error("Unexpected result: " + result)
	}

	result = FixMixedIndentationToSpaces("3", " \thello world  !")
	if result != "    hello world  !" {
		t.Error("Unexpected result: " + result)
	}

	result = FixMixedIndentationToSpaces("2", "  \t hello world !")
	if result != "     hello world !" {
		t.Error("Unexpected result: " + result)
	}
}

func TestFixUndividableIndentationToNearestSpacesAmount(t *testing.T) {
	var result string

	result = FixUndividableIndentationToNearestSpacesAmount("2", "hello")
	if result != "hello" {
		t.Error("String changed but it was already fine. Changed to: " + result)
	}

	result = FixUndividableIndentationToNearestSpacesAmount("2", "  hello")
	if result != "  hello" {
		t.Error("String changed but it was already fine. Changed to: " + result)
	}

	result = FixUndividableIndentationToNearestSpacesAmount("1", "  hello")
	if result != "  hello" {
		t.Error("String changed but it was already fine. Changed to: " + result)
	}

	result = FixUndividableIndentationToNearestSpacesAmount("3", "  hello")
	if result != "   hello" {
		t.Error("Unexpected result: " + result)
	}

	result = FixUndividableIndentationToNearestSpacesAmount("5", "  hello")
	if result != "     hello" {
		t.Error("Unexpected result: " + result)
	}
}

func TestFixTrimTrailingWhitespaceRule(t *testing.T) {
	if FixTrimTrailingWhitespaceRule("true", "") != "" {
		t.Error()
	}

	if FixTrimTrailingWhitespaceRule("true", " a b c") != " a b c" {
		t.Error()
	}

	if FixTrimTrailingWhitespaceRule("true", "abc    \t\t   \t \t \t   ") != "abc" {
		t.Error()
	}

	if FixTrimTrailingWhitespaceRule("false", "abc    \t\t   \t \t \t   ") != "abc    \t\t   \t \t \t   " {
		t.Error()
	}
}
