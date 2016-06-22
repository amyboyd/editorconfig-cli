package editorconfig

import (
	"testing"
)

func TestGetRulesToApplyToSourcePath(t *testing.T) {
	result := GetRulesToApplyToSourcePath(
		"tests/a/b/file.go",
		[]ConfigFile{
			CreateConfigFileStruct("tests/.editorconfig"),
		},
	)

	if result["end_of_line"] != "lf" {
		t.Error("The end_of_line rule should come from the * file pattern")
	}
	if result["indent_style"] != "tabs" {
		t.Error("The indent_style rule should come from the **.go file pattern, overriding the *'s indent_style")
	}
}

func TestGetRulesToApplyToSourcePathWhenNoRulesShouldApply(t *testing.T) {
	result := GetRulesToApplyToSourcePath(
		"some-file-not-affected-by-rules",
		[]ConfigFile{
			CreateConfigFileStruct("tests/.editorconfig"),
		},
	)

	if len(result) != 0 {
		t.Error("No rules should be applied for the file")
	}
}
