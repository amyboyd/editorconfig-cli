package editorconfig

import (
	"strconv"
	"testing"
)

func TestGetConfigFilePathInDirectoryForNonRootPath(t *testing.T) {
	if GetConfigFilePathInDirectory("/home/x") != "/home/x/.editorconfig" {
		t.Error("GetConfigFilePathInDirectory failed for path /home/x")
	}
}

func TestGetConfigFilePathInDirectoryForRootPath(t *testing.T) {
	if GetConfigFilePathInDirectory("/") != "/.editorconfig" {
		t.Error("GetConfigFilePathInDirectory failed for path /")
	}
}

func TestFindConfigFilesWhenDirectoryHasRootConfig(t *testing.T) {
	files := FindConfigFiles([]string{"tests/a/b/has-root-config/file.java"})

	if len(files) != 1 {
		t.Error("Should have found 1 .editorconfig file but found " + strconv.Itoa(len(files)))
	}
}

func TestFindConfigFilesFindsRootConfigInParentDirectory(t *testing.T) {
	files := FindConfigFiles([]string{"tests/a/b/c/d/file.java"})

	if len(files) != 2 {
		t.Error("Should have found 2 .editorconfig files but found " + strconv.Itoa(len(files)))
	}
}
