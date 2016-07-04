package editorconfig

import (
	"testing"
)

func TestFixEndOfLineRule(t *testing.T) {
	input := "\nline\nline 2\rline 3  \n  \r\n\r"

	toLfResult := FixEndOfLineRule("lF", input)
	if toLfResult != "\nline\nline 2\nline 3  \n  \n\n" {
		t.Error("Converting to LF did not work, got: " + GetErrorWithLineBreaksVisible(toLfResult))
	}

	toCrResult := FixEndOfLineRule("Cr", input)
	if toCrResult != "\rline\rline 2\rline 3  \r  \r\r" {
		t.Error("Converting to CR did not work, got: " + GetErrorWithLineBreaksVisible(toCrResult))
	}

	toCrlfResult := FixEndOfLineRule("CrlF", input)
	if toCrlfResult != "\r\nline\r\nline 2\r\nline 3  \r\n  \r\n\r\n" {
		t.Error("Converting to CRLR did not work, got: " + GetErrorWithLineBreaksVisible(toCrlfResult))
	}
}

func TestFixInsertFinalNewLineRule(t *testing.T) {
	input1 := "a\nb\nc\n"
	result1 := FixInsertFinalNewLineRule("true", input1)
	if result1 != input1 {
		t.Error("String was changed despite already having a line at the end")
	}

	input2 := "a\rb\rc\r\r"
	result2 := FixInsertFinalNewLineRule("true", input2)
	if result2 != input2 {
		t.Error("String was changed despite already having a line at the end")
	}

	input3 := "a\r\nb\r\nc\r\n"
	result3 := FixInsertFinalNewLineRule("true", input3)
	if result3 != input3 {
		t.Error("String was changed despite already having a line at the end")
	}

	input4 := "a\nb"
	result4 := FixInsertFinalNewLineRule("true", input4)
	if result4 != "a\nb\n" {
		t.Error("Line was not added at the end")
	}

	input5 := "a\r\nb\r\nc\r\n\n\n\r"
	result5 := FixInsertFinalNewLineRule("false", input5)
	if result5 != "a\r\nb\r\nc" {
		t.Error("Trailing lines were not removed")
	}
}
