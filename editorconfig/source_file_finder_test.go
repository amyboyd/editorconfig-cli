package editorconfig

import (
	"strconv"
	"testing"
)

func TestGetSourceFileExtensions(t *testing.T) {
	result := GetSourceFileExtensions()

	ExpectExtension := func(ext string) {
		if !ContainsString(result, ext) {
			t.Error("Result does not contain extension '" + ext + "'")
		}
	}

	ExpectExtension("go")
	ExpectExtension("java")
	ExpectExtension("php")
}

func TestFindSourceFiles(t *testing.T) {
	result, _ := FindSourceFiles([]string{"tests/a/b/c"})

	if len(result) != 6 {
		t.Error("Result should have 6 files, but has " + strconv.Itoa(len(result)))
	}

	ExpectPath := func(path string) {
		if !ContainsString(result, path) {
			t.Error("Result does not contain path '" + path + "'")
		}
	}

	ExpectPath("tests/a/b/c/.editorconfig")
	ExpectPath("tests/a/b/c/d/file.go")
	ExpectPath("tests/a/b/c/d/file.java")
	ExpectPath("tests/a/b/c/d/file.php")
	ExpectPath("tests/a/b/c/d/keep-trailing-spaces.txt")
	ExpectPath("tests/a/b/c/file.java")
}
