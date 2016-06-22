package editorconfig

import (
	"testing"
)

func TestGetParentDir(t *testing.T) {
	if GetParentDir("a/b/c.txt") != "a/b" {
		t.Error()
	}
	if GetParentDir("/") != "/" {
		t.Error()
	}
	if GetParentDir("/a/b") != "/a" {
		t.Error()
	}
}
