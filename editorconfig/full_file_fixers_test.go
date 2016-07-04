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
